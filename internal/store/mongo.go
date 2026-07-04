package store

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/staroffish/am/internal/config"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	cli *mongo.Client
	db  *mongo.Database
}

func NewMongoClient(cfg config.MongoDBConfig) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(cfg.URI)
	if cfg.Username != "" {
		cred := options.Credential{
			AuthSource: cfg.Database,
			Username:   cfg.Username,
			Password:   cfg.Password,
		}
		opts.SetAuth(cred)
	}

	cli, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}

	if err := cli.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping: %w", err)
	}

	db := cli.Database(cfg.Database)

	if _, err := db.Collection("download_tasks").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}},
		Options: options.Index().SetUnique(true).SetSparse(true),
	}); err != nil {
		return nil, fmt.Errorf("create download_tasks id index: %w", err)
	}

	if _, err := db.Collection("download_tasks").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "anime_id", Value: 1}},
		Options: options.Index().SetUnique(true).SetSparse(true),
	}); err != nil {
		return nil, fmt.Errorf("create download_tasks anime_id index: %w", err)
	}

	return &MongoClient{cli: cli, db: db}, nil
}

func (m *MongoClient) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.cli.Disconnect(ctx)
}

func (m *MongoClient) animeCol() *mongo.Collection { return m.db.Collection("anime") }
func (m *MongoClient) taskCol() *mongo.Collection  { return m.db.Collection("download_tasks") }

func fillAnimeID(a *model.Anime) {
	a.ID = a.ObjectID.Hex()
}

func (m *MongoClient) GetAnime(ctx context.Context, id string) (*model.Anime, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	var ani model.Anime
	err = m.animeCol().FindOne(ctx, bson.M{"_id": oid}).Decode(&ani)
	if err != nil {
		return nil, err
	}
	fillAnimeID(&ani)
	return &ani, nil
}

func (m *MongoClient) ListAnime(ctx context.Context, keyword string, skip, limit int, asc bool) ([]model.Anime, int64, error) {
	filter := bson.M{}
	if keyword != "" {
		filter = bson.M{"$or": []bson.M{
			{"animenamecn": bson.M{"$regex": keyword, "$options": "i"}},
			{"animenamejp": bson.M{"$regex": keyword, "$options": "i"}},
			{"cast": bson.M{"$regex": keyword, "$options": "i"}},
			{"serialsduri": bson.M{"$regex": keyword, "$options": "i"}},
			{"type": bson.M{"$regex": keyword, "$options": "i"}},
			{"status": bson.M{"$regex": keyword, "$options": "i"}},
			{"stordir": bson.M{"$regex": keyword, "$options": "i"}},
		}}
	}

	total, _ := m.animeCol().CountDocuments(ctx, filter)

	if limit <= 0 {
		limit = 50
	}
	sortDir := -1
	if asc {
		sortDir = 1
	}
	opts := options.Find().SetSort(bson.D{{Key: "updatetime", Value: sortDir}}).SetSkip(int64(skip)).SetLimit(int64(limit))
	cursor, err := m.animeCol().Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var animes []model.Anime
	if err := cursor.All(ctx, &animes); err != nil {
		return nil, 0, err
	}
	for i := range animes {
		fillAnimeID(&animes[i])
	}
	if animes == nil {
		animes = []model.Anime{}
	}
	return animes, total, nil
}

func (m *MongoClient) SaveAnime(ctx context.Context, ani *model.Anime) error {
	now := time.Now()

	if ani.ID == "" {
		ani.ObjectID = primitive.NewObjectID()
		fillAnimeID(ani)
		ani.CreatedAt = now
	} else {
		oid, err := primitive.ObjectIDFromHex(ani.ID)
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}
		ani.ObjectID = oid
	}

	ani.UpdatedAt = now
	_, err := m.animeCol().UpdateOne(ctx,
		bson.M{"_id": ani.ObjectID},
		bson.M{"$set": ani},
		options.Update().SetUpsert(true),
	)
	return err
}

func (m *MongoClient) DeleteAnime(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}
	_, err = m.animeCol().DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (m *MongoClient) MarkAnimeDone(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	ani, err := m.GetAnime(ctx, id)
	if err != nil {
		return err
	}

	now := time.Now()
	y, s := util.GetNowSeason()
	serialsDuri := strings.TrimRight(ani.SerialsDuri, "~")
	_, err = m.animeCol().UpdateOne(ctx,
		bson.M{"_id": oid},
		bson.M{"$set": bson.M{
			"status":      "已完结",
			"updatetime":  now,
			"serialsduri": serialsDuri + "~" + fmt.Sprintf("%04d/%02d", y, s),
		}},
	)
	return err
}

func (m *MongoClient) ListTasks(ctx context.Context) ([]model.DownloadTask, error) {
	opts := options.Find().SetSort(bson.D{{Key: "updated_at", Value: -1}})
	cursor, err := m.taskCol().Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []model.DownloadTask
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	if tasks == nil {
		tasks = []model.DownloadTask{}
	}
	return tasks, nil
}

func (m *MongoClient) GetTask(ctx context.Context, id int32) (*model.DownloadTask, error) {
	var task model.DownloadTask
	err := m.taskCol().FindOne(ctx, bson.M{"id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (m *MongoClient) GetTaskByAnimeID(ctx context.Context, animeID string) (*model.DownloadTask, error) {
	var task model.DownloadTask
	err := m.taskCol().FindOne(ctx, bson.M{"anime_id": animeID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (m *MongoClient) CreateTask(ctx context.Context, task *model.DownloadTask) error {
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	var maxTask model.DownloadTask
	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	if err := m.taskCol().FindOne(ctx, bson.M{}, opts).Decode(&maxTask); err == nil {
		task.ID = maxTask.ID + 1
	} else {
		task.ID = 1
	}

	_, err := m.taskCol().InsertOne(ctx, task)
	return err
}

func (m *MongoClient) UpdateTask(ctx context.Context, task *model.DownloadTask) error {
	task.UpdatedAt = time.Now()
	_, err := m.taskCol().ReplaceOne(ctx, bson.M{"id": task.ID}, task)
	return err
}

func (m *MongoClient) UpdateTaskChapter(ctx context.Context, id, latestChapter int32) error {
	_, err := m.taskCol().UpdateOne(ctx,
		bson.M{"id": id},
		bson.M{"$set": bson.M{
			"latest_chapter": latestChapter,
			"updated_at":     time.Now(),
		}},
	)
	return err
}

func (m *MongoClient) DeleteTask(ctx context.Context, id int32) error {
	_, err := m.taskCol().DeleteOne(ctx, bson.M{"id": id})
	return err
}
