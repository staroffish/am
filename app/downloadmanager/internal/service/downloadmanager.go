package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	commonPb "github.com/staroffish/am/api/common"
	pb "github.com/staroffish/am/api/downloadmanager/v1"
	"github.com/staroffish/am/app/downloadmanager/internal/biz"
	dtoDownloadManager "github.com/staroffish/am/common/dto/downloadmanager"
)

type DownloadmanagerService struct {
	downloadManager *biz.DownloadManager
	log             *log.Helper
	pb.UnimplementedDownloadmanagerServer
}

func NewDownloadmanagerService(downloadManager *biz.DownloadManager, logger log.Logger) *DownloadmanagerService {
	return &DownloadmanagerService{
		log:             log.NewHelper(logger),
		downloadManager: downloadManager,
	}
}

func (s *DownloadmanagerService) ScanTaskAndDownload(ctx context.Context, req *commonPb.Empty) (*pb.ScanTaskAndDownloadReply, error) {
	deadline, _ := ctx.Deadline()
	s.log.Infof("start ScanTaskAndDownload now=%v ctx.deadline=%v", time.Now(), deadline)
	downloadTasks, err := s.downloadManager.ScanTaskAndDownload(ctx)
	if err != nil {
		s.log.Errorf("DownloadmanagerService.ScanTaskAndDownload:biz.DownloadManager.ScanTaskAndDownload error: %v", err)
		return nil, err
	}

	replyTasks := downloadTasksToReplayTaskInfos(downloadTasks)

	reply := &pb.ScanTaskAndDownloadReply{
		CreatedTasks: replyTasks,
	}
	return reply, nil
}

func (s *DownloadmanagerService) ScanTask(ctx context.Context, req *commonPb.Empty) (*pb.ScanTaskReply, error) {
	deadline, _ := ctx.Deadline()
	s.log.Infof("start ScanTask now=%v ctx.deadline=%v", time.Now(), deadline)
	downloadTasks, err := s.downloadManager.Scan(ctx)
	if err != nil {
		s.log.Errorf("DownloadmanagerService.ScanTaskAndDownload:biz.DownloadManager.ScanTask error: %v", err)
		return nil, err
	}

	replyTasks := downloadTasksToReplayTaskInfos(downloadTasks)

	reply := &pb.ScanTaskReply{
		MatchedTasks: replyTasks,
	}
	return reply, nil
}

func (s *DownloadmanagerService) AddTask(ctx context.Context, req *pb.AddTaskRequest) (*commonPb.Empty, error) {
	deadline, _ := ctx.Deadline()
	s.log.Infof("start AddTask now=%v ctx.deadline=%v", time.Now(), deadline)
	downloadTaskInfo := &dtoDownloadManager.DownloadTaskInfo{
		Name:          req.Name,
		Regexp:        req.Regexp,
		AnimeId:       req.AnimeId,
		StorePath:     req.StorePath,
		LatestChapter: req.LatestChapter,
	}
	if err := s.downloadManager.AddTask(ctx, downloadTaskInfo); err != nil {
		s.log.Errorf("DownloadmanagerService.ScanTaskAndDownload:biz.DownloadManager.SaveTask error: %v", err)
		return nil, err
	}
	return &commonPb.Empty{}, nil
}

func (s *DownloadmanagerService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*commonPb.Empty, error) {
	deadline, _ := ctx.Deadline()
	s.log.Infof("start DeleteTask now=%v ctx.deadline=%v", time.Now(), deadline)
	if err := s.downloadManager.DeleteTask(ctx, req.Id); err != nil {
		s.log.Errorf("DownloadmanagerService.ScanTaskAndDownload:biz.DownloadManager.DeleteTask error: %v", err)
		return nil, err
	}
	return &commonPb.Empty{}, nil
}

func (s *DownloadmanagerService) ListTask(ctx context.Context, req *commonPb.Empty) (*pb.ListTaskReply, error) {
	deadline, _ := ctx.Deadline()
	s.log.Infof("start ListTask now=%v ctx.deadline=%v", time.Now(), deadline)
	downloadTaskInofs := s.downloadManager.ListTask(ctx)

	reply := &pb.ListTaskReply{
		Tasks: []*pb.DownloadTaskInfo{},
	}

	for _, downloadTaskInfo := range downloadTaskInofs {
		replyTaskInfo := pb.DownloadTaskInfo{
			Id:            downloadTaskInfo.Id,
			Name:          downloadTaskInfo.Name,
			Regexp:        downloadTaskInfo.Regexp,
			LatestChapter: downloadTaskInfo.LatestChapter,
			StorePath:     downloadTaskInfo.StorePath,
			UpdateTime:    downloadTaskInfo.UpdateTime,
			AnimeId:       downloadTaskInfo.AnimeId,
		}

		reply.Tasks = append(reply.Tasks, &replyTaskInfo)
	}
	return reply, nil
}

func (s *DownloadmanagerService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*commonPb.Empty, error) {
	deadline, _ := ctx.Deadline()
	s.log.Infof("start UpdateTask now=%v ctx.deadline=%v", time.Now(), deadline)
	err := s.downloadManager.UpdateTaskInfo(ctx, &dtoDownloadManager.DownloadTaskInfo{
		Id:            req.Id,
		Name:          req.Name,
		Regexp:        req.Regexp,
		AnimeId:       req.AnimeId,
		StorePath:     req.StorePath,
		LatestChapter: req.LatestChapter,
	})
	if err != nil {
		s.log.Errorf("DownloadmanagerService.UpdateTask:biz.DownloadManager.UpdateTaskInfo error: %v", err)
		return nil, err
	}
	return &commonPb.Empty{}, nil
}

func (s *DownloadmanagerService) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskReply, error) {
	deadline, _ := ctx.Deadline()
	s.log.Infof("start GetTask now=%v ctx.deadline=%v", time.Now(), deadline)
	var taskInfo *dtoDownloadManager.DownloadTaskInfo
	var err error
	if req.AnimeId != "" {
		taskInfo, err = s.downloadManager.GetDownloadTaskInfoByAnimeId(ctx, req.AnimeId)
		if err != nil {
			s.log.Errorf("DownloadmanagerService.GetTask error: %v", err)
			return nil, err
		}
	} else {
		taskInfo, err = s.downloadManager.GetDownloadTaskInfoById(ctx, req.Id)
		if err != nil {
			s.log.Errorf("DownloadmanagerService.GetTask error: %v", err)
			return nil, err
		}
	}

	return &pb.GetTaskReply{Task: &pb.DownloadTaskInfo{
		Id:            taskInfo.Id,
		Name:          taskInfo.Name,
		Regexp:        taskInfo.Regexp,
		LatestChapter: taskInfo.LatestChapter,
		StorePath:     taskInfo.StorePath,
		UpdateTime:    taskInfo.UpdateTime,
		AnimeId:       taskInfo.AnimeId,
	}}, nil
}

func downloadTasksToReplayTaskInfos(downloadTasks []*biz.DownloadTask) []*pb.DownloadTask {
	replyTasks := []*pb.DownloadTask{}
	for _, downloadTask := range downloadTasks {
		replyTask := &pb.DownloadTask{
			Name:         downloadTask.TaskName,
			ChapterStart: downloadTask.ChapterStart,
			ChapterEnd:   downloadTask.ChapterEnd,
			MagnetLink:   downloadTask.MagnetLink,
			AnimeId:      downloadTask.AnimeId,
		}
		replyTasks = append(replyTasks, replyTask)
	}

	return replyTasks
}
