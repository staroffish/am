package data

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	downloaderV1 "github.com/staroffish/am/api/downloader/v1"
	"github.com/staroffish/am/app/downloadmanager/internal/biz"
	dtoSpider "github.com/staroffish/am/common/dto/spider"
)

type downloadManagerRepo struct {
	log *log.Helper
	db  *Data
	*DownloadTask
	downloader downloaderV1.DownloaderClient
}

func NewDownloaderClient(r registry.Discovery) downloaderV1.DownloaderClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///downloader"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return downloaderV1.NewDownloaderClient(conn)
}

// NewGreeterRepo .
func NewDownloadManagerRepo(database *Data, downloadTask *DownloadTask, downloader downloaderV1.DownloaderClient, logger log.Logger) biz.DownloadManagerRepo {
	return &downloadManagerRepo{
		log:          log.NewHelper(logger),
		db:           database,
		downloader:   downloader,
		DownloadTask: downloadTask,
	}
}

func (d *downloadManagerRepo) CreateDownloadTask(ctx context.Context, id int32, latestChapter int32, MagnetLink string) error {
	downloadTask := d.GetDownloadTaskInfoById(id)
	if downloadTask == nil {
		err := fmt.Errorf("downloadManagerRepo.CreateDownloadTask:GetTaskInfoById downloadTask %d not found", id)
		d.log.Error(err)
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 调用下载器的接口，创建下载任务
	_, err := d.downloader.Add(timeoutCtx, &downloaderV1.AddRequest{
		Link:      MagnetLink,
		StorePath: downloadTask.StorePath,
	})
	if err != nil {
		d.log.Errorf("downloadManagerRepo.CreateDownloadTask:downloader.Add error: %v", err)
		return err
	}

	// 更新最新集数
	if err := d.UpdateLatestChapter(timeoutCtx, id, latestChapter); err != nil {
		d.log.Errorf("downloadManagerRepo.CreateDownloadTask:UpdateLatestChapter error: %v", err)
		return err
	}
	return nil
}

func (d *downloadManagerRepo) GetAnimeMagnets(ctx context.Context) ([]*dtoSpider.AnimeMagnet, error) {
	animeMagnets := []*dtoSpider.AnimeMagnet{}

	date := time.Now()
	for {
		dateStr := date.Format("2006-01-02")
		date = date.Add(time.Hour * time.Duration(-24))
		key := fmt.Sprintf("anime:link:%s", dateStr)
		d.log.Infof("downloadManagerRepo.GetAnimeMagnets HGetAll %s", key)
		animeMagnetMap, err := d.db.redisCli.HGetAll(ctx, key).Result()
		if err != nil {
			d.log.Errorf("downloadManagerRepo.GetAnimeMagnets.HGetAll %s error:%v", key, err)
			break
		}

		if len(animeMagnetMap) == 0 {
			break
		}

		for k, v := range animeMagnetMap {
			animeMagnet := &dtoSpider.AnimeMagnet{
				Name:       k,
				MagnetLink: v,
			}
			animeMagnets = append(animeMagnets, animeMagnet)
		}
	}
	return animeMagnets, nil
}
