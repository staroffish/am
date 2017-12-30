package rd

import (
	"fmt"
	"global"
)

// RdTask - 远程下载任务的结构体
type RdTask struct {
	// 任务标识
	Id string
	// 任务链接
	Link string
	// 文件/目录名
	Name string
	// 路径
	SavePath string
	// 进度
	Progress int
	// 状态
	State string
	// 文件大小
	Size int
	// 任务类型
	TaskType string
}

var downloaderList map[string]Downloader

// Downloader - 远程下载任务的接口
type Downloader interface {
	AddTask(t *RdTask) error
	PauseTask(t *RdTask) error
	ResumeTask(t *RdTask) error
	DeleteTask(t *RdTask) error
	GetAllTask() ([]RdTask, error)
}

// SetDownloader - 添加下载器
func SetDownloader(typ string, dnlr Downloader) {
	defer global.TraceLog("rd.SetDownloader")()
	downloaderList[typ] = dnlr
}

// AddTask - 添加任务
func AddTask(t *RdTask, typ string) error {
	defer global.TraceLog("rd.AddTask")()
	dnldr, ok := downloaderList[typ]
	if !ok {
		return fmt.Errorf("rd.AddTask:Unsupport Type:%s", typ)
	}

	return dnldr.AddTask(t)
}

// PauseTask - 暂停任务
func PauseTask(t *RdTask, typ string) error {
	defer global.TraceLog("rd.PauseTask")()
	dnldr, ok := downloaderList[typ]
	if !ok {
		return fmt.Errorf("rd.AddTask:Unsupport Type:%s", typ)
	}

	return dnldr.PauseTask(t)
}

// DeleteTask - 删除任务
func DeleteTask(t *RdTask, typ string) error {
	defer global.TraceLog("rd.DeleteTask")()
	dnldr, ok := downloaderList[typ]
	if !ok {
		return fmt.Errorf("rd.AddTask:Unsupport Type:%s", typ)
	}

	return dnldr.DeleteTask(t)
}

// GetAllTask - 取得所有任务
func GetAllTask() ([]RdTask, error) {
	defer global.TraceLog("rd.GetAllTask")()
	var rtnList = make([]RdTask, 0, 0)

	for typ, downloader := range downloaderList {
		taskList, err := downloader.GetAllTask()
		if err != nil {
			global.Log.Errorf("am:rd.GetAllTask:%v:type:%s", err, typ)
			continue
		}
		rtnList = append(rtnList, taskList...)
	}

	return rtnList, nil
}

// ResumeTask - 重启任务
func ResumeTask(t *RdTask, typ string) error {
	defer global.TraceLog("rd.ResumeTask")()
	dnldr, ok := downloaderList[typ]
	if !ok {
		return fmt.Errorf("rd.AddTask:Unsupport Type:%s", typ)
	}

	return dnldr.ResumeTask(t)
}
