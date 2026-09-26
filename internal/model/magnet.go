package model

import "time"

// Magnet 是爬虫抓到的磁力，按抓取日期归档，替代原来的 Redis hash(anime:link:<date>)。
// 对外的 DTO 仍然是 spider.go 里的 AnimeMagnet。
type Magnet struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement" json:"-"`
	Date       string    `gorm:"column:date;uniqueIndex:idx_magnet_date_name;index" json:"date"`
	Name       string    `gorm:"column:name;uniqueIndex:idx_magnet_date_name" json:"name"`
	MagnetLink string    `gorm:"column:magnet_link" json:"magnet_link"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
}

func (Magnet) TableName() string { return "magnets" }
