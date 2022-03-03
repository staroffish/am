package data

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/staroffish/am/app/downloadmanager/internal/biz"
	"github.com/staroffish/am/app/downloadmanager/internal/conf"
	dtoSpider "github.com/staroffish/am/common/dto/spider"
)

type downloadManagerRepo struct {
	log *log.Helper
	db  *Data
	*DownloadTask
}

// NewGreeterRepo .
func NewDownloadManagerRepo(database *Data, confData *conf.DownloadManagerServerConfig, downloadTask *DownloadTask, logger log.Logger) biz.DownloadManagerRepo {
	return &downloadManagerRepo{
		log:          log.NewHelper(logger),
		db:           database,
		DownloadTask: downloadTask,
	}
}

func (d *downloadManagerRepo) CreateDownloadTask(ctx context.Context, id int32, latestChapter int32, MagnetLink string) error {
	// 调用下载器的接口，创建下载任务

	// 更新最新集数
	if err := d.UpdateLatestChapter(ctx, id, latestChapter); err != nil {
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
