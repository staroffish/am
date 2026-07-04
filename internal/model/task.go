package model

import "time"

type DownloadTask struct {
	ID            int32     `bson:"id"             json:"id"`
	Regexp        string    `bson:"regexp"         json:"regexp"`
	LatestChapter int32     `bson:"latest_chapter" json:"latest_chapter"`
	AnimeID       string    `bson:"anime_id"       json:"anime_id"`
	CreatedAt     time.Time `bson:"created_at"     json:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at"     json:"updated_at"`
}

type MatchedTask struct {
	TaskID       int32  `json:"task_id"`
	TaskName     string `json:"task_name"`
	ChapterStart int32  `json:"chapter_start"`
	ChapterEnd   int32  `json:"chapter_end"`
	MagnetLink   string `json:"magnet_link"`
	StorePath    string `json:"store_path"`
	AnimeID      string `json:"anime_id"`
}
