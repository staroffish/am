package service

import (
	"context"
	"log"

	"github.com/staroffish/am/internal/downloader"
	"github.com/staroffish/am/internal/model"
)

type DownloaderService struct {
	qb   *downloader.QBittorrentClient
	log  *log.Logger
}

func NewDownloaderService(qb *downloader.QBittorrentClient, logger *log.Logger) *DownloaderService {
	return &DownloaderService{qb: qb, log: logger}
}

func (s *DownloaderService) Add(ctx context.Context, link, storePath string) error {
	return s.qb.Add(ctx, link, storePath)
}

func (s *DownloaderService) Delete(ctx context.Context, hash string) error {
	return s.qb.Delete(ctx, hash)
}

func (s *DownloaderService) Pause(ctx context.Context, hash string) error {
	return s.qb.Pause(ctx, hash)
}

func (s *DownloaderService) Resume(ctx context.Context, hash string) error {
	return s.qb.Resume(ctx, hash)
}

func (s *DownloaderService) List(ctx context.Context) ([]model.TorrentInfo, error) {
	return s.qb.List(ctx)
}
