package http

import (
	"fmt"
	"global"
	"rd"
)

// HttpDownloader - http下载库
type HttpDownloader struct {
	pool  redisPool
	tasks map[string]*httpTask
	db    int
}

// 任务状态
const (
	Downloading = "Downloading"
	Paused      = "Paused"
	Finished    = "Finished"
	Error       = "Error"

	CanTruncate    = "OK"
	CanNotTruncate = "NO"
)

// HttpType/HttpsType - 下载类型字符串
var HttpType = "http"
var HttpsType = "https"

type httpTask struct {
	Id           string   // 任务ID(保存文件的全路径)
	Path         string   // 存储路径
	FileName     string   // 文件名
	Url          string   // 下载链接
	CanTruncate  string   // 是否可被分割(是否支持断点续传)
	TaskType     string   // 任务类型(http/https)
	State        string   // 任务状态
	TotalSize    float64  // 总大小
	FinishedSize float64  // 完成大小
	CreateTime   string   // 创建时间
	Controler    chan int // 控制用信道
}

func init() {
	rd.SetDownloader(HttpType, &HttpDownloader{db: 0})
	rd.SetDownloader(HttpsType, &HttpDownloader{db: 1})
}

// NewDownloader - 初始化下载器
func (d *HttpDownloader) InitDownloader(cfg *global.Config) error {
	defer global.TraceLog("http.HttpDownloader.NewDownloader")()

	if d.pool.Pool == nil {
		d.pool.initRedis(cfg.RedisAddr, cfg.RedisPasswd, d.db, 3)
	}
	d.tasks = make(map[string]*httpTask)
	// err := d.pool.getAllTask(d.tasks)
	// if err != nil {
	// 	return fmt.Errorf("HttpDownloader.InitDownloader:%v", err)
	// }
	return nil
}

// AddTask - 添加任务
func (d *HttpDownloader) AddTask(t *rd.RdTask) error {
	defer global.TraceLog("http.HttpDownloader.AddTask")()

	return nil
}

// PauseTask - 暂停任务
func (d *HttpDownloader) PauseTask(t *rd.RdTask) error {
	defer global.TraceLog("http.HttpDownloader.PauseTask")()

	return nil
}

// DeleteTask - 删除任务(只删除任务不删除文件)
func (d *HttpDownloader) DeleteTask(t *rd.RdTask) error {
	defer global.TraceLog("http.HttpDownloader.DeleteTask")()

	return nil
}

// ResumeTask - 重新开始任务
func (d *HttpDownloader) ResumeTask(t *rd.RdTask) error {
	defer global.TraceLog("http.HttpDownloader.ResumeTask")()

	return nil
}

// GetAllTask - 获取所有任务
func (d *HttpDownloader) GetAllTask() ([]rd.RdTask, error) {
	defer global.TraceLog("http.HttpDownloader.GetAllTask")()
	var rdTaskList []rd.RdTask

	err := d.pool.getAllTask(d.tasks)
	if err != nil {
		return nil, fmt.Errorf("http.GetAllTask:%v", err)
	}

	for id, task := range d.tasks {
		var rdt rd.RdTask
		rdt.Ids = id
		rdt.Link = task.Url
		rdt.Name = task.FileName
		rdt.SavePath = task.Path
		rdt.Progress = int(task.FinishedSize / task.TotalSize * 100)
		rdt.Size = int64(task.TotalSize)
		rdt.TaskType = task.TaskType
		rdt.State = task.State
		rdt.CreateTime = task.CreateTime
		rdTaskList = append(rdTaskList, rdt)
	}

	return rdTaskList, nil
}
