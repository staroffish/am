package model

import "time"

type Anime struct {
	ID          string    `gorm:"column:id;primaryKey;type:text" json:"id"`
	AnimeNameCn string    `gorm:"column:animenamecn" json:"animenamecn"`
	AnimeNameJp string    `gorm:"column:animenamejp" json:"animenamejp"`
	Cast        string    `gorm:"column:cast" json:"cast"`
	SerialsDuri string    `gorm:"column:serialsduri" json:"serialsduri"`
	Type        string    `gorm:"column:type" json:"type"`
	ImageBin    []byte    `gorm:"column:image" json:"-"`
	StorDir     string    `gorm:"column:stordir" json:"stordir"`
	Status      string    `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updatetime" json:"updated_at"`
}

func (Anime) TableName() string { return "anime" }

type AnimeFile struct {
	Name     string `json:"name"`
	FullPath string `json:"full_path"`
}
