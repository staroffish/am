// Command legacyimport 把旧版本（MongoDB + Redis）的数据离线搬到新的 SQLite。
//
// 它故意做成一个独立的 Go 子模块（tools/legacyimport/go.mod），
// 这样主模块的依赖里不会重新出现 mongo-driver / go-redis。
// 只依赖数据库可达，不需要旧服务进程在跑。
//
// 用法：
//
//	cd tools/legacyimport
//	go run . -db /本机磁盘/am.db \
//	         -mongo-pass 'xxx' -redis-pass 'xxx'
//
// 或者用 make migrate 编出 bin/migrate 再运行。
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/staroffish/am/internal/config"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/store"
)

var (
	mongoURI  = flag.String("mongo", "mongodb://mongo.holygrail.com:27017", "MongoDB URI")
	mongoDB   = flag.String("mongo-db", "am", "MongoDB database")
	mongoUser = flag.String("mongo-user", "am", "MongoDB username")
	mongoPass = flag.String("mongo-pass", os.Getenv("MONGO_PASS"), "MongoDB password（也可用环境变量 MONGO_PASS）")

	redisAddr = flag.String("redis", "redis.holygrail.com:6379", "Redis 地址")
	redisPass = flag.String("redis-pass", os.Getenv("REDIS_PASS"), "Redis password（也可用环境变量 REDIS_PASS）")
	redisDB   = flag.Int("redis-db", 2, "Redis db")

	dbPath = flag.String("db", "data/am.db", "目标 SQLite 文件路径（必须在本地磁盘上）")

	skipImages  = flag.Bool("skip-images", false, "不迁移封面图")
	skipMagnets = flag.Bool("skip-magnets", false, "不迁移磁力缓存")
	dryRun      = flag.Bool("dry-run", false, "只统计，不写库")
)

// 旧 Mongo 文档结构（字段名与 bson tag 保持和重构前一致）。
type mongoAnime struct {
	ObjectID    primitive.ObjectID `bson:"_id"`
	AnimeNameCn string             `bson:"animenamecn"`
	AnimeNameJp string             `bson:"animenamejp"`
	Cast        string             `bson:"cast"`
	SerialsDuri string             `bson:"serialsduri"`
	Type        string             `bson:"type"`
	ImageBin    []byte             `bson:"image"`
	StorDir     string             `bson:"stordir"`
	Status      string             `bson:"status"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updatetime"`
}

type mongoTask struct {
	ID            int32     `bson:"id"`
	Regexp        string    `bson:"regexp"`
	LatestChapter int32     `bson:"latest_chapter"`
	AnimeID       string    `bson:"anime_id"`
	CreatedAt     time.Time `bson:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at"`
}

const magnetKeyPrefix = "anime:link:"

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags)

	ctx := context.Background()

	var db *store.DB
	if !*dryRun {
		var err error
		db, err = store.NewDB(config.SQLiteConfig{Path: *dbPath})
		if err != nil {
			log.Fatalf("open sqlite %s: %v", *dbPath, err)
		}
		defer db.Close()
		log.Printf("target sqlite: %s", *dbPath)
	} else {
		log.Printf("DRY RUN: 不会写入任何数据")
	}

	mongoCli, err := connectMongo(ctx)
	if err != nil {
		log.Fatalf("connect mongo: %v", err)
	}
	defer func() {
		c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = mongoCli.Disconnect(c)
	}()

	animeTotal, imageTotal, err := importAnime(ctx, mongoCli, db)
	if err != nil {
		log.Fatalf("import anime: %v", err)
	}
	log.Printf("anime: %d (with image %d)", animeTotal, imageTotal)

	taskTotal, err := importTasks(ctx, mongoCli, db)
	if err != nil {
		log.Fatalf("import tasks: %v", err)
	}
	log.Printf("download_tasks: %d", taskTotal)

	if *skipMagnets {
		log.Printf("magnets: skipped")
		log.Printf("done")
		return
	}

	magnetTotal, dateTotal, err := importMagnets(ctx, db)
	if err != nil {
		log.Fatalf("import magnets: %v", err)
	}
	log.Printf("magnets: %d (dates %d)", magnetTotal, dateTotal)
	log.Printf("done")
}

func connectMongo(ctx context.Context) (*mongo.Client, error) {
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(*mongoURI)
	if *mongoUser != "" {
		opts.SetAuth(options.Credential{
			AuthSource: *mongoDB,
			Username:   *mongoUser,
			Password:   *mongoPass,
		})
	}
	cli, err := mongo.Connect(cctx, opts)
	if err != nil {
		return nil, err
	}
	if err := cli.Ping(cctx, nil); err != nil {
		return nil, fmt.Errorf("ping（用户名/密码是否正确？）: %w", err)
	}
	return cli, nil
}

func importAnime(ctx context.Context, cli *mongo.Client, db *store.DB) (int, int, error) {
	cursor, err := cli.Database(*mongoDB).Collection("anime").Find(ctx, bson.M{})
	if err != nil {
		return 0, 0, err
	}
	defer cursor.Close(ctx)

	total, withImage := 0, 0
	for cursor.Next(ctx) {
		var doc mongoAnime
		if err := cursor.Decode(&doc); err != nil {
			return total, withImage, fmt.Errorf("decode: %w", err)
		}

		ani := &model.Anime{
			ID:          doc.ObjectID.Hex(),
			AnimeNameCn: doc.AnimeNameCn,
			AnimeNameJp: doc.AnimeNameJp,
			Cast:        doc.Cast,
			SerialsDuri: doc.SerialsDuri,
			Type:        doc.Type,
			StorDir:     doc.StorDir,
			Status:      doc.Status,
			CreatedAt:   doc.CreatedAt,
			UpdatedAt:   doc.UpdatedAt,
		}
		if !*skipImages {
			ani.ImageBin = doc.ImageBin
		}
		if len(ani.ImageBin) > 0 {
			withImage++
		}
		total++

		if *dryRun {
			continue
		}
		if err := db.ImportAnime(ctx, ani); err != nil {
			return total, withImage, fmt.Errorf("import %s: %w", ani.ID, err)
		}
	}
	return total, withImage, cursor.Err()
}

func importTasks(ctx context.Context, cli *mongo.Client, db *store.DB) (int, error) {
	opts := options.Find().SetSort(bson.D{{Key: "id", Value: 1}})
	cursor, err := cli.Database(*mongoDB).Collection("download_tasks").Find(ctx, bson.M{}, opts)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	total := 0
	for cursor.Next(ctx) {
		var doc mongoTask
		if err := cursor.Decode(&doc); err != nil {
			return total, fmt.Errorf("decode: %w", err)
		}
		total++

		if *dryRun {
			continue
		}
		task := &model.DownloadTask{
			ID:            doc.ID,
			Regexp:        doc.Regexp,
			LatestChapter: doc.LatestChapter,
			AnimeID:       doc.AnimeID,
			CreatedAt:     doc.CreatedAt,
			UpdatedAt:     doc.UpdatedAt,
		}
		if err := db.ImportTask(ctx, task); err != nil {
			return total, fmt.Errorf("import task %d: %w", doc.ID, err)
		}
	}
	return total, cursor.Err()
}

func importMagnets(ctx context.Context, db *store.DB) (int, int, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     *redisAddr,
		Password: *redisPass,
		DB:       *redisDB,
	})
	defer rdb.Close()

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := rdb.Ping(pingCtx).Err(); err != nil {
		return 0, 0, fmt.Errorf("connect redis（密码是否正确？）: %w", err)
	}

	var (
		keys   []string
		cursor uint64
	)
	for {
		batch, next, err := rdb.Scan(ctx, cursor, magnetKeyPrefix+"*", 500).Result()
		if err != nil {
			return 0, 0, err
		}
		keys = append(keys, batch...)
		cursor = next
		if cursor == 0 {
			break
		}
	}
	log.Printf("redis: %d magnet keys", len(keys))

	magnetTotal, dateTotal := 0, 0
	for _, key := range keys {
		date := strings.TrimPrefix(key, magnetKeyPrefix)
		if len(date) != 10 { // YYYY-MM-DD
			continue
		}

		vals, err := rdb.HGetAll(ctx, key).Result()
		if err != nil {
			return magnetTotal, dateTotal, fmt.Errorf("hgetall %s: %w", key, err)
		}
		if len(vals) == 0 {
			continue
		}

		dateTotal++
		magnetTotal += len(vals)
		if *dryRun {
			continue
		}

		rows := make([]*model.AnimeMagnet, 0, len(vals))
		for name, link := range vals {
			rows = append(rows, &model.AnimeMagnet{Name: name, MagnetLink: link})
		}
		if _, err := db.SaveMagnets(ctx, date, rows); err != nil {
			return magnetTotal, dateTotal, fmt.Errorf("save %s: %w", key, err)
		}
	}
	return magnetTotal, dateTotal, nil
}
