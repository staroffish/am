package data

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	dto "github.com/staroffish/am/common/dto/downloadmanager"
	etcd "go.etcd.io/etcd/client/v3"
)

type DownloadTask struct {
	prefix                string
	watchPrefix           string
	db                    *Data
	downloadTaskInfos     []*dto.DownloadTaskInfo
	downloadTaskByName    map[string]*dto.DownloadTaskInfo
	downloadTaskById      map[int32]*dto.DownloadTaskInfo
	downloadTaskByAnimeId map[string]*dto.DownloadTaskInfo
	rwLock                *sync.RWMutex
	log                   log.Helper
}

type TaskEtcdPrefix string

func NewDownloadTask(db *Data, prefix TaskEtcdPrefix, logger log.Logger) *DownloadTask {
	downloadTask := &DownloadTask{
		prefix:      string(prefix) + "/",
		watchPrefix: string(prefix),
		db:          db,
		rwLock:      &sync.RWMutex{},
		log:         *log.NewHelper(logger),
	}

	watchChan := downloadTask.db.etcdCli.Watch(context.Background(), downloadTask.watchPrefix)
	go func() {
		for range watchChan {
			downloadTask.log.Infof("%s changed", downloadTask.watchPrefix)
			if err := downloadTask.SyncDownloadTask(); err != nil {
				downloadTask.log.Errorf("watch downloadTask.UpdateDownloadTask error:%v", err)
			}
		}
	}()
	if err := downloadTask.SyncDownloadTask(); err != nil {
		panic(fmt.Sprintf("init downloadTask.UpdateDownloadTask error:%v", err))
	}
	return downloadTask
}

func (s *DownloadTask) GetDownloadTaskInfos() []*dto.DownloadTaskInfo {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	var downloadTaskInfos = make([]*dto.DownloadTaskInfo, 0, len(s.downloadTaskInfos))

	for _, downloadTask := range s.downloadTaskInfos {
		returnTaskInfo := *downloadTask
		downloadTaskInfos = append(downloadTaskInfos, &returnTaskInfo)
	}
	return downloadTaskInfos
}

func (s *DownloadTask) SyncDownloadTask() error {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	resp, err := s.db.etcdCli.Get(context.Background(), s.prefix, etcd.WithPrefix(), etcd.WithSort(etcd.SortByCreateRevision, etcd.SortAscend))
	if err != nil {
		s.log.Errorf("client.Get %s --prefix error:%v", err)
		return err
	}

	downloadInfos := []*dto.DownloadTaskInfo{}
	var downloadTaskByName = make(map[string]*dto.DownloadTaskInfo)
	var downloadTaskById = make(map[int32]*dto.DownloadTaskInfo)
	var downloadTaskByAnimeId = make(map[string]*dto.DownloadTaskInfo)
	for _, item := range resp.Kvs {
		downloadInfo := &dto.DownloadTaskInfo{}
		err = json.Unmarshal(item.Value, downloadInfo)
		if err != nil {
			s.log.Errorf("json.Unmarshal %s error:%v", item.Key, err)
			return err
		}
		downloadInfo.Version = item.Version
		downloadInfo.CreateRevision = item.CreateRevision
		downloadInfo.ModRevision = item.ModRevision
		downloadInfos = append(downloadInfos, downloadInfo)
		downloadTaskByName[downloadInfo.Name] = downloadInfo
		downloadTaskById[downloadInfo.Id] = downloadInfo
		downloadTaskByAnimeId[downloadInfo.AnimeId] = downloadInfo
	}

	s.downloadTaskInfos = downloadInfos
	s.downloadTaskByName = downloadTaskByName
	s.downloadTaskById = downloadTaskById
	s.downloadTaskByAnimeId = downloadTaskByAnimeId
	return nil
}

func (s *DownloadTask) isTaskCanAdd(taskInfo *dto.DownloadTaskInfo) error {
	if _, ok := s.downloadTaskById[taskInfo.Id]; ok {
		return fmt.Errorf("DownloadTask.isTaskCanAdd error: download task id %d already exists", taskInfo.Id)
	}
	if _, ok := s.downloadTaskByName[taskInfo.Name]; ok {
		return fmt.Errorf("DownloadTask.isTaskCanAdd error: download task Name %s already exists", taskInfo.Name)
	}
	if _, ok := s.downloadTaskByAnimeId[taskInfo.AnimeId]; ok {
		return fmt.Errorf("DownloadTask.isTaskCanAdd error: download task anime_id %s already exists", taskInfo.AnimeId)
	}
	return nil
}

func (s *DownloadTask) AddTaskInfo(ctx context.Context, taskInfo *dto.DownloadTaskInfo) error {
	s.rwLock.RLock()
	if err := s.isTaskCanAdd(taskInfo); err != nil {
		s.rwLock.RUnlock()
		return err
	}
	var id int32 = 1
	if len(s.downloadTaskInfos) > 0 {
		id = s.downloadTaskInfos[len(s.downloadTaskInfos)-1].Id + 1
	}
	s.rwLock.RUnlock()
	for {
		updateTime := time.Now()
		updateUnixTime := updateTime.Unix()

		taskInfo.UpdateTime = updateTime.Format("2006-01-02 15:04:05")
		taskInfo.Id = id

		value, err := json.Marshal(taskInfo)
		if err != nil {
			s.log.Errorf("DownloadTask.AddTaskInfo:json.Marshal error:%v", err)
			return err
		}

		key := fmt.Sprintf("%s%d", s.prefix, id)
		txn := s.db.etcdCli.Txn(ctx).
			If(
				etcd.Compare(etcd.Version(key), "=", 0),
			).
			Then(
				etcd.OpPut(key, string(value)),
				etcd.OpPut(s.watchPrefix, fmt.Sprintf("%d", updateUnixTime)),
			)
		resp, err := txn.Commit()
		if err != nil {
			s.log.Errorf("DownloadTask.AddTaskInfo:txn.Commit error:%v", err)
			return err
		}
		if resp.Succeeded {
			break
		}
		id++
	}

	return nil
}

func (s *DownloadTask) DeleteTask(ctx context.Context, id int32) error {
	key := fmt.Sprintf("%s%d", s.prefix, id)
	updateTime := time.Now().Unix()
	txn := s.db.etcdCli.Txn(ctx).
		If(
			etcd.Compare(etcd.Version(key), ">", 0),
		).
		Then(
			etcd.OpDelete(key),
			etcd.OpPut(s.watchPrefix, fmt.Sprintf("%d", updateTime)),
		)
	resp, err := txn.Commit()
	if err != nil {
		s.log.Errorf("DownloadTask.DeleteTask:txn.Commit error:%v", err)
		return err
	}
	if !resp.Succeeded {
		return fmt.Errorf("download task %d already deleted", id)
	}
	return nil
}

func (s *DownloadTask) UpdateLatestChapter(ctx context.Context, id, latestChapter int32) error {
	taskInfo := s.GetDownloadTaskInfoById(id)
	if taskInfo == nil {
		err := fmt.Errorf("DownloadTask.UpdateLatestChapter:task id:%d not found", id)
		s.log.Error(err)
		return err
	}
	taskInfo.LatestChapter = latestChapter
	err := s.UpdateTaskInfo(ctx, taskInfo)
	if err != nil {
		s.log.Errorf("DownloadTask.UpdateLatestChapter:UpdateTaskInfo error:%v", err)
		return err
	}

	return nil
}

func (s *DownloadTask) UpdateTaskInfo(ctx context.Context, taskInfo *dto.DownloadTaskInfo) error {
	taskInfoCached := s.GetDownloadTaskInfoById(taskInfo.Id)
	if taskInfoCached == nil {
		err := fmt.Errorf("DownloadTask.UpdateTaskInfo:task id:%d not found", taskInfo.Id)
		s.log.Error(err)
		return err
	}

	updateTime := time.Now()
	taskInfo.UpdateTime = updateTime.Format("2006-01-02 15:04:05")
	valueStr, err := json.Marshal(taskInfo)
	if err != nil {
		s.log.Errorf("DownloadTask.UpdateLatestChapter:json.Marshal error:%v", err)
		return err
	}
	key := fmt.Sprintf("%s%d", s.prefix, taskInfo.Id)

	s.log.Infof("key=%s, value=%s, ModRevision=%d", key, valueStr, taskInfoCached.ModRevision)

	txn := s.db.etcdCli.Txn(ctx).
		If(
			etcd.Compare(etcd.ModRevision(key), "=", taskInfoCached.ModRevision),
		).
		Then(
			etcd.OpPut(key, string(valueStr)),
			etcd.OpPut(s.watchPrefix, fmt.Sprintf("%d", updateTime.Unix())),
		)
	_, err = txn.Commit()
	if err != nil {
		s.log.Errorf("DownloadTask.UpdateTaskInfo:txn.Commit error:%v", err)
		return err
	}
	return nil
}

func (s *DownloadTask) GetDownloadTaskInfoById(id int32) *dto.DownloadTaskInfo {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	taskInfo, ok := s.downloadTaskById[id]
	if !ok {
		return nil
	}
	retTaskInfo := *taskInfo

	return &retTaskInfo
}

func (s *DownloadTask) GetDownloadTaskInfoByAnimeId(id string) *dto.DownloadTaskInfo {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	taskInfo, ok := s.downloadTaskByAnimeId[id]
	if !ok {
		return nil
	}
	retTaskInfo := *taskInfo

	return &retTaskInfo
}
