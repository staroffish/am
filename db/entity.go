package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// AdTask - 自动下载任务的结构体
type AdTask struct {
	// Task ID
	Id string `bson:"_id"`
	// 动画网页的搜索地址
	Url string `json:"url"`
	// 检索动漫关键字
	UrlParam string `json:"urlparam"`
	// 抓取动画单集网页的正则
	SchExp string `json:"schexp"`
	// 抓取集数
	SchChapt int `json:"schchapt"`
	// 抓取磁链正则
	MagExp string `json:"magexp"`
	// 更新时间
	UpdateTime time.Time `bson:"updatetime"`
	// 动漫ID
	AnimeID bson.ObjectId `json:"animeid" bson:"animeid"`
	// 任务名称
	Name string `json:"-"`
	// 存储路径
	StorePath string `json:"-"`
}

// Anime - 动画结构体
type Anime struct {
	// 动漫ID(HEX)
	AnimeID string
	// 动漫ID(binary)
	ID bson.ObjectId `bson:"_id"`
	// 动漫中文名称
	AnimeNameCn string `bson:"animenamecn"`
	// 动漫日文名称
	AnimeNameJp string `bson:"animenamejp"`
	// 主要声优
	Cast string `bson:"cast"`
	// 连载期间
	SerialsDuri string `bson:"serialsduri"`
	// 类型
	Type string `bson:"type"`
	// 图片
	ImageBin []byte `bson:"image"`
	// 存储路径
	StorDir string `bson:"stordir"`
	// 播放路径
	PlayDir string `bson:"playdir"`
	// 连载状态
	Status string `bson:"status"`
	// 更新时间
	UpdateTime time.Time `bson:"updatetime"`
}
