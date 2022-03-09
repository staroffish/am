package service

import (
	"context"

	commonPb "github.com/staroffish/am/api/common"
	pb "github.com/staroffish/am/api/downloader/v1"
	"github.com/staroffish/am/app/downloader/internal/biz"
)

type DownloaderService struct {
	*biz.Downloader
	pb.UnimplementedDownloaderServer
}

func NewDownloaderService(downloader *biz.Downloader) *DownloaderService {
	return &DownloaderService{
		Downloader: downloader,
	}
}

func (s *DownloaderService) Add(ctx context.Context, req *pb.AddRequest) (*commonPb.Empty, error) {
	return &commonPb.Empty{}, s.Downloader.Add(ctx, req.Link, req.StorePath)
}
func (s *DownloaderService) Delete(ctx context.Context, req *pb.DeleteRequest) (*commonPb.Empty, error) {
	return &commonPb.Empty{}, s.Downloader.Delete(ctx, req.Id)
}
func (s *DownloaderService) List(ctx context.Context, req *commonPb.Empty) (*pb.ListResponse, error) {
	taskInfos, err := s.Downloader.List(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListResponse{
		TaskInfos: []*pb.ListResponseTaskInfo{},
	}

	for _, taskInfo := range taskInfos {
		retTaskInfo := &pb.ListResponseTaskInfo{
			Id:          taskInfo.Id,
			Name:        taskInfo.Name,
			Size:        taskInfo.Size,
			Progress:    taskInfo.Progress,
			StorePath:   taskInfo.StorePath,
			Status:      taskInfo.Status,
			CreatedTime: taskInfo.CreatedTime,
		}
		resp.TaskInfos = append(resp.TaskInfos, retTaskInfo)
	}

	return resp, nil
}
func (s *DownloaderService) Pause(ctx context.Context, req *pb.PauseRequest) (*commonPb.Empty, error) {
	return &commonPb.Empty{}, s.Downloader.Pause(ctx, req.Id)
}
func (s *DownloaderService) Resume(ctx context.Context, req *pb.ResumeRequest) (*commonPb.Empty, error) {
	return &commonPb.Empty{}, s.Downloader.Resume(ctx, req.Id)
}
