package biz

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	dtoDownloadmanager "github.com/staroffish/am/common/dto/downloadmanager"
	dtoSpider "github.com/staroffish/am/common/dto/spider"
)

type DownloadManagerRepo interface {
	CreateDownloadTask(ctx context.Context, id int32, latestChapter int32, MagnetLink string) error
	GetAnimeMagnets(ctx context.Context) ([]*dtoSpider.AnimeMagnet, error)
	AddTaskInfo(context.Context, *dtoDownloadmanager.DownloadTaskInfo) error
	DeleteTask(ctx context.Context, id int32) error
	GetDownloadTaskInfos() []*dtoDownloadmanager.DownloadTaskInfo
	UpdateLatestChapter(ctx context.Context, id, latestChapter int32) error
	UpdateTaskInfo(ctx context.Context, taskInfo *dtoDownloadmanager.DownloadTaskInfo) error
	GetDownloadTaskInfoById(id int32) *dtoDownloadmanager.DownloadTaskInfo
	GetDownloadTaskInfoByAnimeId(id string) *dtoDownloadmanager.DownloadTaskInfo
}

type DownloadTask struct {
	TaskId       int32
	TaskName     string
	ChapterStart int32 // 存在 一个连接包含多集的情况，使用ChapterStart和ChapterEnd来表达
	ChapterEnd   int32
	MagnetLink   string
	StorePath    string
}

type DownloadManager struct {
	repo DownloadManagerRepo
	log  *log.Helper
}

func NewDownloadManager(repo DownloadManagerRepo, logger log.Logger) *DownloadManager {
	return &DownloadManager{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (a *DownloadManager) Scan(ctx context.Context) ([]*DownloadTask, error) {

	animeMagnets, err := a.repo.GetAnimeMagnets(ctx)
	if err != nil {
		a.log.Errorf("DownloadManager.repo.GetAnimeMagnets errror: %v", err)
		return nil, err
	}

	downloadTasks := []*DownloadTask{}
	taskInfos := a.repo.GetDownloadTaskInfos()

	for _, taskInfo := range taskInfos {
		// 最新集数+1 作为匹配条件
		latestChapter := taskInfo.LatestChapter + 1
		regexpStr := fmt.Sprintf(taskInfo.Regexp, latestChapter)
		a.log.Infof("task info: id:%d, name:%s, regexpStr:%s, latestChapter:%d, store_path:%s", taskInfo.Id, taskInfo.Name, regexpStr, latestChapter, taskInfo.StorePath)
		reg, err := regexp.Compile(regexpStr)
		if err != nil {
			a.log.Errorf("DownloadManager.Scan:regexp.Compile error:%v", err)
			continue
		}
		for _, animeMagnet := range animeMagnets {
			// 匹配动漫条目
			foundList := reg.FindStringSubmatch(animeMagnet.Name)
			if foundList == nil || len(foundList) < 2 {
				continue
			}

			downloadTask := &DownloadTask{
				TaskId:       taskInfo.Id,
				TaskName:     animeMagnet.Name,
				MagnetLink:   animeMagnet.MagnetLink,
				ChapterStart: latestChapter,
				ChapterEnd:   latestChapter,
			}

			if len(foundList) > 2 {
				chapterEnd, err := strconv.ParseInt(foundList[2], 10, 64)
				if err == nil {
					downloadTask.ChapterEnd = int32(chapterEnd)
				}
			}

			latestChapter = downloadTask.ChapterEnd + 1 // 如果是一个任务多集的情况去左后一集的集数作为最新集的集数
			a.log.Infof("matched: name:%s, latestChapter:%d, MangetLink:%s, matchedList:%v", downloadTask.TaskName, downloadTask.ChapterEnd, downloadTask.MagnetLink, foundList)

			downloadTasks = append(downloadTasks, downloadTask)

			// 更新最后一集的集数
			regexpStr = fmt.Sprintf(taskInfo.Regexp, latestChapter)
			a.log.Infof("update latestChapter task info: id:%d, name:%s, regexpStr:%s, latestChapter:%d, store_path:%s", taskInfo.Id, taskInfo.Name, regexpStr, latestChapter, taskInfo.StorePath)
			reg, err = regexp.Compile(regexpStr)
			if err != nil {
				a.log.Errorf("update latestChapter DownloadManager.Scan:regexp.Compile error:%v", err)
				continue
			}
		}
	}

	return downloadTasks, nil
}

func (a *DownloadManager) ScanTaskAndDownload(ctx context.Context) ([]*DownloadTask, error) {
	downloadTasks, err := a.Scan(ctx)
	if err != nil {
		a.log.Errorf("DownloadManager.ScanTaskAndDownload.Scan error:%v", err)
		return nil, err
	}

	createdTask := []*DownloadTask{}

	for _, downloadTask := range downloadTasks {
		if err := a.repo.CreateDownloadTask(ctx, downloadTask.TaskId, downloadTask.ChapterEnd, downloadTask.MagnetLink); err != nil {
			a.log.Errorf("DownloadManager.ScanTaskAndDownload.CreateDownloadTask error:%v", err)
			continue
		}
		createdTask = append(createdTask, downloadTask)
	}

	return createdTask, nil
}

func (a *DownloadManager) AddTask(ctx context.Context, taskInfo *dtoDownloadmanager.DownloadTaskInfo) error {
	if err := a.repo.AddTaskInfo(ctx, taskInfo); err != nil {
		a.log.Errorf("DownloadManager.SaveTask.SaveTaskInfo error:%v", err)
		return err
	}
	return nil
}

func (a *DownloadManager) DeleteTask(ctx context.Context, id int32) error {
	if err := a.repo.DeleteTask(ctx, id); err != nil {
		a.log.Errorf("DownloadManager.DeleteTask.DeleteTask error:%v", err)
		return err
	}
	return nil
}

func (a *DownloadManager) ListTask(ctx context.Context) []*dtoDownloadmanager.DownloadTaskInfo {
	return a.repo.GetDownloadTaskInfos()
}

func (a *DownloadManager) UpdateTaskInfo(ctx context.Context, taskInfo *dtoDownloadmanager.DownloadTaskInfo) error {
	if err := a.repo.UpdateTaskInfo(ctx, taskInfo); err != nil {
		a.log.Errorf("DownloadManager.SaveTask.SaveTaskInfo error:%v", err)
		return err
	}
	return nil
}

func (a *DownloadManager) GetDownloadTaskInfoById(ctx context.Context, id int32) (*dtoDownloadmanager.DownloadTaskInfo, error) {
	taskInfo := a.repo.GetDownloadTaskInfoById(id)
	if taskInfo == nil {
		err := fmt.Errorf("DownloadManager.GetDownloadTaskInfoById error: %d task not found", id)
		a.log.Error(err)
		return nil, err
	}
	return taskInfo, nil
}

func (a *DownloadManager) GetDownloadTaskInfoByAnimeId(ctx context.Context, id string) (*dtoDownloadmanager.DownloadTaskInfo, error) {
	taskInfo := a.repo.GetDownloadTaskInfoByAnimeId(id)
	if taskInfo == nil {
		err := fmt.Errorf("DownloadManager.GetDownloadTaskInfoByAnimeId error: %s task not found", id)
		a.log.Error(err)
		return nil, err
	}
	return taskInfo, nil
}
