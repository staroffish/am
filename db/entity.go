package db

import "gopkg.in/mgo.v2/bson"
import "time"

// AdTask - 自动下载任务的结构体
type AdTask struct {
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
	// 动漫ID
	AnimeID bson.ObjectId `json:"animeid" bson:"animeid"`
}

const (
	// 下载中
	Downloading = "iota"
	// 暂停
	Stoped
	// 下载完成
	Completed
)

// RdTask - 远程下载任务的结构体
type RdTask struct {
	// 文件/目录名
	Name string `json:"name"`
	// 路径
	SavePath string `json:"save_path"`
	// 进度
	Progress string `json:"progress"`
	// 状态
	State string `json:"state"`
	// 文件大小
	Size int `json:"total_wanted"`
}

const (
	// 连载中
	OnAir = iota
	// 已完结
	Finished
)

// Anime - 动画结构体
type Anime struct {
	// 动漫ID
	AnimeID bson.ObjectId `json:"animeid" bson:"_id"`
	// 动漫中文名称
	AnimeNameCn string `json:"animenamecn"`
	// 动漫日文名称
	AnimeNameJp string `json:"animenamejp"`
	// 集数
	Chapter int `json:"chapter"`
	// 总集数
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
	// 连载状态
	Status int `json:"status"`
	// 更新时间
	UpdateTime time.Time `json:"updatetime"`
}
