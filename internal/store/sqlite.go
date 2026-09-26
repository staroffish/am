package store

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	gormlogger "gorm.io/gorm/logger"

	"github.com/staroffish/am/internal/config"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/util"
)

// DB 是唯一的持久化入口（SQLite + GORM）。
// 设计前提：单实例部署，SQLite 文件放在本机磁盘上。
type DB struct {
	gorm *gorm.DB
}

func NewDB(cfg config.SQLiteConfig) (*DB, error) {
	path := cfg.Path
	if path == "" {
		path = "data/am.db"
	}
	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, fmt.Errorf("create db dir %s: %w", dir, err)
		}
	}

	dsn := fmt.Sprintf(
		"file:%s?_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=foreign_keys(1)",
		path,
	)
	gdb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("open sqlite %s: %w", path, err)
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, fmt.Errorf("get sql.DB: %w", err)
	}
	// SQLite 同一时刻只能有一个写者。单连接可以彻底避免 "database is locked"，
	// 以这个应用的访问量完全够用。
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(0)

	if err := gdb.AutoMigrate(&model.Anime{}, &model.DownloadTask{}, &model.Magnet{}); err != nil {
		return nil, fmt.Errorf("auto migrate: %w", err)
	}

	// download_tasks.anime_id 需要唯一，但空字符串不参与约束（对应原来 Mongo 的 sparse unique index）。
	if err := gdb.Exec(
		`CREATE UNIQUE INDEX IF NOT EXISTS idx_download_tasks_anime_id ON download_tasks(anime_id) WHERE anime_id <> ''`,
	).Error; err != nil {
		return nil, fmt.Errorf("create anime_id index: %w", err)
	}

	return &DB{gorm: gdb}, nil
}

func (d *DB) Close() error {
	sqlDB, err := d.gorm.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func newHexID() string {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%024x", time.Now().UnixNano())
	}
	return hex.EncodeToString(b)
}

// ---------------- anime ----------------

func (d *DB) GetAnime(ctx context.Context, id string) (*model.Anime, error) {
	var ani model.Anime
	if err := d.gorm.WithContext(ctx).First(&ani, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &ani, nil
}

func animeFilter(q *gorm.DB, keyword string) *gorm.DB {
	if keyword == "" {
		return q
	}
	like := "%" + keyword + "%"
	return q.Where(
		`animenamecn LIKE ? OR animenamejp LIKE ? OR "cast" LIKE ? OR serialsduri LIKE ? OR "type" LIKE ? OR status LIKE ? OR stordir LIKE ?`,
		like, like, like, like, like, like, like,
	)
}

func (d *DB) ListAnime(ctx context.Context, keyword string, skip, limit int, asc bool) ([]model.Anime, int64, error) {
	var total int64
	if err := animeFilter(d.gorm.WithContext(ctx).Model(&model.Anime{}), keyword).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit <= 0 {
		limit = 50
	}
	order := "updatetime DESC"
	if asc {
		order = "updatetime ASC"
	}

	var list []model.Anime
	err := animeFilter(d.gorm.WithContext(ctx).Model(&model.Anime{}), keyword).
		Omit("ImageBin").
		Order(order).
		Offset(skip).
		Limit(limit).
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	if list == nil {
		list = []model.Anime{}
	}
	return list, total, nil
}

func (d *DB) SaveAnime(ctx context.Context, ani *model.Anime) error {
	now := time.Now()
	if ani.ID == "" {
		ani.ID = newHexID()
	}
	if ani.CreatedAt.IsZero() {
		ani.CreatedAt = now
	}
	ani.UpdatedAt = now
	return d.gorm.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(ani).Error
}

// ImportAnime 原样写入（保留 created_at / updatetime），只给一次性数据迁移用。
func (d *DB) ImportAnime(ctx context.Context, ani *model.Anime) error {
	if ani.ID == "" {
		ani.ID = newHexID()
	}
	return d.gorm.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(ani).Error
}

func (d *DB) DeleteAnime(ctx context.Context, id string) error {
	return d.gorm.WithContext(ctx).Where("id = ?", id).Delete(&model.Anime{}).Error
}

func (d *DB) MarkAnimeDone(ctx context.Context, id string) error {
	return d.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var ani model.Anime
		if err := tx.Select("id", "serialsduri").First(&ani, "id = ?", id).Error; err != nil {
			return err
		}

		y, s := util.GetNowSeason()
		serialsDuri := strings.TrimRight(ani.SerialsDuri, "~")

		return tx.Model(&model.Anime{}).Where("id = ?", id).Updates(map[string]interface{}{
			"status":      "已完结",
			"updatetime":  time.Now(),
			"serialsduri": serialsDuri + "~" + fmt.Sprintf("%04d/%02d", y, s),
		}).Error
	})
}

// ---------------- download tasks ----------------

func (d *DB) ListTasks(ctx context.Context) ([]model.DownloadTask, error) {
	var tasks []model.DownloadTask
	if err := d.gorm.WithContext(ctx).Order("updated_at DESC").Find(&tasks).Error; err != nil {
		return nil, err
	}
	if tasks == nil {
		tasks = []model.DownloadTask{}
	}
	return tasks, nil
}

func (d *DB) GetTask(ctx context.Context, id int32) (*model.DownloadTask, error) {
	var task model.DownloadTask
	if err := d.gorm.WithContext(ctx).First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (d *DB) GetTaskByAnimeID(ctx context.Context, animeID string) (*model.DownloadTask, error) {
	var task model.DownloadTask
	if err := d.gorm.WithContext(ctx).First(&task, "anime_id = ?", animeID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// CreateTask 插入一条规则；task.ID 为 0 时由 SQLite 自增分配，非 0 时按给定值写入（迁移用）。
func (d *DB) CreateTask(ctx context.Context, task *model.DownloadTask) error {
	return d.gorm.WithContext(ctx).Create(task).Error
}

// UpdateTask 只更新用户可编辑的字段，避免把 created_at 覆盖成零值。
func (d *DB) UpdateTask(ctx context.Context, task *model.DownloadTask) error {
	return d.gorm.WithContext(ctx).Model(&model.DownloadTask{}).Where("id = ?", task.ID).Updates(map[string]interface{}{
		"regexp":         task.Regexp,
		"latest_chapter": task.LatestChapter,
		"anime_id":       task.AnimeID,
	}).Error
}

func (d *DB) UpdateTaskChapter(ctx context.Context, id, latestChapter int32) error {
	return d.gorm.WithContext(ctx).Model(&model.DownloadTask{}).Where("id = ?", id).Update("latest_chapter", latestChapter).Error
}

func (d *DB) DeleteTask(ctx context.Context, id int32) error {
	return d.gorm.WithContext(ctx).Where("id = ?", id).Delete(&model.DownloadTask{}).Error
}

// ---------------- magnets ----------------

// SaveMagnets 按日期批量写入磁力，已存在的 (date, name) 直接跳过。
func (d *DB) SaveMagnets(ctx context.Context, date string, magnets []*model.AnimeMagnet) (int, error) {
	rows := make([]model.Magnet, 0, len(magnets))
	now := time.Now()
	for _, m := range magnets {
		if m == nil || m.Name == "" {
			continue
		}
		rows = append(rows, model.Magnet{
			Date:       date,
			Name:       m.Name,
			MagnetLink: m.MagnetLink,
			CreatedAt:  now,
		})
	}
	if len(rows) == 0 {
		return 0, nil
	}
	res := d.gorm.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(rows, 200)
	return int(res.RowsAffected), res.Error
}

func (d *DB) ListMagnetsByDate(ctx context.Context, date string) ([]model.Magnet, error) {
	var rows []model.Magnet
	err := d.gorm.WithContext(ctx).Where("date = ?", date).Order("name ASC").Find(&rows).Error
	return rows, err
}

func (d *DB) ListMagnetDates(ctx context.Context) ([]string, error) {
	dates := []string{}
	err := d.gorm.WithContext(ctx).Model(&model.Magnet{}).
		Distinct().Order("date DESC").Pluck("date", &dates).Error
	return dates, err
}

func (d *DB) ListMagnetsSince(ctx context.Context, sinceDate string) ([]model.Magnet, error) {
	var rows []model.Magnet
	err := d.gorm.WithContext(ctx).Where("date >= ?", sinceDate).Find(&rows).Error
	return rows, err
}

func (d *DB) DeleteMagnetsBefore(ctx context.Context, beforeDate string) (int64, error) {
	res := d.gorm.WithContext(ctx).Where("date < ?", beforeDate).Delete(&model.Magnet{})
	return res.RowsAffected, res.Error
}
