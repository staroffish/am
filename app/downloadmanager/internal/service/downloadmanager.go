package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
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

func (s *DownloadmanagerService) ScanTaskAndDownload(ctx context.Context, req *pb.Empty) (*pb.ScanTaskAndDownloadReply, error) {
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

func (s *DownloadmanagerService) ScanTask(ctx context.Context, req *pb.Empty) (*pb.ScanTaskReply, error) {
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

func (s *DownloadmanagerService) AddTask(ctx context.Context, req *pb.AddTaskRequest) (*pb.Empty, error) {
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
	return &pb.Empty{}, nil
}

func (s *DownloadmanagerService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.Empty, error) {
	if err := s.downloadManager.DeleteTask(ctx, req.Id); err != nil {
		s.log.Errorf("DownloadmanagerService.ScanTaskAndDownload:biz.DownloadManager.DeleteTask error: %v", err)
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *DownloadmanagerService) ListTask(ctx context.Context, req *pb.ListTaskRequest) (*pb.ListTaskReply, error) {
	downloadTaskInofs := s.downloadManager.ListTask(ctx)

	reply := &pb.ListTaskReply{
		Tasks: []*pb.ListTaskReply_DownloadTaskInfo{},
	}

	for _, downloadTaskInfo := range downloadTaskInofs {
		replyTaskInfo := pb.ListTaskReply_DownloadTaskInfo{
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

func (s *DownloadmanagerService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.Empty, error) {
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
	return &pb.Empty{}, nil
}

func downloadTasksToReplayTaskInfos(downloadTasks []*biz.DownloadTask) []*pb.DownloadTask {
	replyTasks := []*pb.DownloadTask{}
	for _, downloadTask := range downloadTasks {
		replyTask := &pb.DownloadTask{
			Name:         downloadTask.TaskName,
			ChapterStart: downloadTask.ChapterStart,
			ChapterEnd:   downloadTask.ChapterEnd,
			MagnetLink:   downloadTask.MagnetLink,
		}
		replyTasks = append(replyTasks, replyTask)
	}

	return replyTasks
}
