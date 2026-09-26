package model

import "time"

type DownloadTask struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Regexp        string    `gorm:"column:regexp" json:"regexp"`
	LatestChapter int32     `gorm:"column:latest_chapter" json:"latest_chapter"`
	AnimeID       string    `gorm:"column:anime_id" json:"anime_id"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (DownloadTask) TableName() string { return "download_tasks" }

type MatchedTask struct {
	TaskID       int32  `json:"task_id"`
	TaskName     string `json:"task_name"`
	ChapterStart int32  `json:"chapter_start"`
	ChapterEnd   int32  `json:"chapter_end"`
	MagnetLink   string `json:"magnet_link"`
	StorePath    string `json:"store_path"`
	AnimeID      string `json:"anime_id"`
}
