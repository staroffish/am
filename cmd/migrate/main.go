package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/staroffish/am/internal/config"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/store"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	etcdEndpoint = flag.String("etcd", "192.168.3.5:2379", "etcd endpoint")
	etcdPrefix   = flag.String("prefix", "/download-manager/manager-archer/tasks/", "etcd key prefix")
	mongoURI     = flag.String("mongo", "mongodb://localhost:27017", "MongoDB URI")
	mongoDB      = flag.String("db", "am", "MongoDB database")
	mongoUser    = flag.String("user", "am", "MongoDB username")
	mongoPass    = flag.String("pass", "amPasswd", "MongoDB password")
	dryRun       = flag.Bool("dry-run", false, "print tasks without writing to MongoDB")
)

type etcdTask struct {
	ID            int32  `json:"id"`
	Name          string `json:"name"`
	Regexp        string `json:"regexp"`
	LatestChapter int32  `json:"latest_chapter"`
	StorePath     string `json:"store_path"`
	UpdateTime    string `json:"updateTime"`
	AnimeID       string `json:"anime_id"`
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Connecting to etcd:", *etcdEndpoint)
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{*etcdEndpoint},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("etcd connect: %v", err)
	}
	defer etcdCli.Close()

	ctx := context.Background()
	resp, err := etcdCli.Get(ctx, *etcdPrefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		log.Fatalf("etcd get: %v", err)
	}

	log.Printf("Found %d tasks in etcd", len(resp.Kvs))

	type migrateItem struct {
		task model.DownloadTask
		name string
		path string
	}
	var items []migrateItem

	for _, kv := range resp.Kvs {
		var t etcdTask
		if err := json.Unmarshal(kv.Value, &t); err != nil {
			log.Printf("skip key %s: parse error: %v", string(kv.Key), err)
			continue
		}

		updateTime, _ := time.Parse("2006-01-02 15:04:05", t.UpdateTime)

		task := model.DownloadTask{
			ID:            t.ID,
			Regexp:        t.Regexp,
			LatestChapter: t.LatestChapter,
			AnimeID:       t.AnimeID,
			CreatedAt:     updateTime,
			UpdatedAt:     updateTime,
		}

		items = append(items, migrateItem{task: task, name: t.Name, path: t.StorePath})
		fmt.Printf("  [%d] %s (ch:%d) %s\n", task.ID, t.Name, task.LatestChapter, t.StorePath)
	}

	if *dryRun {
		log.Println("DRY RUN - no data written to MongoDB")
		return
	}

	if len(items) == 0 {
		log.Println("No tasks to migrate")
		return
	}

	log.Println("Connecting to MongoDB:", *mongoURI)
	mongoCfg := config.MongoDBConfig{
		URI:      *mongoURI,
		Database: *mongoDB,
		Username: *mongoUser,
		Password: *mongoPass,
	}
	mongoCli, err := store.NewMongoClient(mongoCfg)
	if err != nil {
		log.Fatalf("mongo connect: %v", err)
	}
	defer mongoCli.Close()

	ctx2 := context.Background()
	written := 0
	skipped := 0
	animeCreated := 0

	for _, item := range items {
		err := mongoCli.CreateTask(ctx2, &item.task)
		if err != nil {
			log.Printf("write task %d: %v", item.task.ID, err)
			skipped++
			continue
		}
		written++

		if item.task.AnimeID != "" {
			_, err := mongoCli.GetAnime(ctx2, item.task.AnimeID)
			if err != nil {
				ani := &model.Anime{
					AnimeNameJp: item.name,
					Status:      "连载中",
					StorDir:     item.path,
					CreatedAt:   item.task.CreatedAt,
					UpdatedAt:   item.task.UpdatedAt,
				}
				ani.ID = item.task.AnimeID
				if err := mongoCli.SaveAnime(ctx2, ani); err != nil {
					log.Printf("create anime for task %d: %v", item.task.ID, err)
				} else {
					animeCreated++
				}
			}
		}
	}

	log.Printf("Migration complete: %d tasks written, %d skipped, %d anime created", written, skipped, animeCreated)
	if skipped > 0 {
		os.Exit(1)
	}
}
