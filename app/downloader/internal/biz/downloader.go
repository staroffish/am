package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type DownloaderRepo interface {
	Add(ctx context.Context, link, storePath string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*DownloaderTaskInfo, error)
	Pause(ctx context.Context, id string) error
	Resume(ctx context.Context, id string) error
}

type DownloaderTaskInfo struct {
	Name        string
	Size        int64
	StorePath   string
	Progress    float32
	Status      string
	CreatedTime string
	Id          string
}

type Downloader struct {
	repo DownloaderRepo
	log  *log.Helper
}

func NewDownloader(repo DownloaderRepo, logger log.Logger) *Downloader {
	return &Downloader{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (d *Downloader) List(ctx context.Context) ([]*DownloaderTaskInfo, error) {
	taskInfos, err := d.repo.List(ctx)
	if err != nil {
		d.log.Errorf("Downloader.List error:%v", err)
		return nil, err
	}
	return taskInfos, nil
}

func (d *Downloader) Add(ctx context.Context, link, storePath string) error {
	if err := d.repo.Add(ctx, link, storePath); err != nil {
		d.log.Errorf("Downloader.Add error:%v", err)
		return err
	}
	return nil
}

func (d *Downloader) Delete(ctx context.Context, id string) error {
	if err := d.repo.Delete(ctx, id); err != nil {
		d.log.Errorf("Downloader.Delete error:%v", err)
		return err
	}
	return nil
}

func (d *Downloader) Pause(ctx context.Context, id string) error {
	if err := d.repo.Pause(ctx, id); err != nil {
		d.log.Errorf("Downloader.Pause error:%v", err)
		return err
	}
	return nil
}

func (d *Downloader) Resume(ctx context.Context, id string) error {
	if err := d.repo.Resume(ctx, id); err != nil {
		d.log.Errorf("Downloader.Resume error:%v", err)
		return err
	}
	return nil
}
