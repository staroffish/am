package db

import "gopkg.in/mgo.v2/bson"
import "time"

// AdTask - 自动下载任务的结构体
type AdTask struct {
	// 动画网页的搜索地址
	Url string `json:"url"`
	// 抓取动画单集网页的正则
	SchExp string `json:"schexp"`
	// 抓取磁链正则
	MagExp string `json:"magexp"`
	// 动漫ID
	AnimeID bson.ObjectId `json:"animeid" bson:"_id"`
}

// Anime - 动画结构体
type Anime struct {
	// 动漫ID
	AnimeID bson.ObjectId `json:"animeid" bson:"_id"`
	// 动漫中文名称
	AnimeNameCn string `json:"animenamecn"`
	// 动漫日文名称
	AnimeNameJp string `json:"animenamejp"`
	// 集数
	ChapterCnt int `json:"chaptercnt"`
	// 连载期间
	SerialsDuri string `json:"serialsduri"`
	// 类型
	Type string `json:"type"`
	// 图片路径
	ImagePath string `json:"imagepath"`
	// 存储路径
	StorDir string `json:"stordir"`
	// 播放路径
	PlayDir string `json:"playdir"`
	// 更新时间
	UpdateTime time.Time `json:"updatetime"`
}
