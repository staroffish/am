package downloader

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/staroffish/am/app/downloader/internal/biz"
	"github.com/staroffish/am/common/config"
)

const (
	QBITTORENT = 0
)

type DownloaderBaseRepo struct {
	log           *log.Helper
	componentName config.ComponentName
	redisCli      *redis.Client
}

func NewDownloaderBaseRepo(componentName config.ComponentName, redisCli *redis.Client, logger log.Logger) *DownloaderBaseRepo {
	return &DownloaderBaseRepo{
		componentName: componentName,
		redisCli:      redisCli,
		log:           log.NewHelper(logger),
	}
}

func (d *DownloaderBaseRepo) Add(ctx context.Context, link, storePath string) error {
	return nil
}

func (d *DownloaderBaseRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func (d *DownloaderBaseRepo) List(ctx context.Context) ([]*biz.DownloaderTaskInfo, error) {
	return nil, nil
}

func (d *DownloaderBaseRepo) Pause(ctx context.Context, id string) error {
	return nil
}

func (d *DownloaderBaseRepo) Resume(ctx context.Context, id string) error {
	return nil
}

func (d *DownloaderBaseRepo) SaveCookies(ctx context.Context, key string, cookies []*http.Cookie) error {
	for _, cookie := range cookies {
		_, err := d.redisCli.HSet(ctx, key, cookie.Name, cookie.Value).Result()
		if err != nil {
			d.log.Errorf("DownloaderBaseRepo.SaveCookie:redisCli.HSet %s %s %s error:%v", key, cookie.Name, cookie.Value, err)
			return err
		}
	}
	if err := d.redisCli.Expire(ctx, key, 5*time.Minute).Err(); err != nil {
		d.log.Errorf("DownloaderBaseRepo.SaveCookie:redisCli.Expire %s error:%v", key, err)
		return err
	}
	return nil
}

func (d *DownloaderBaseRepo) LoadCookies(ctx context.Context, key string) ([]*http.Cookie, error) {
	value, err := d.redisCli.HGetAll(ctx, key).Result()
	if err != nil && err != redis.Nil {
		d.log.Errorf("DownloaderBaseRepo.SaveCookie:redisCli.Get %s error:%v", key, err)
		return nil, err
	}
	cookies := []*http.Cookie{}
	for k, v := range value {
		cookies = append(cookies, &http.Cookie{
			Name:  k,
			Value: v,
		})
	}
	return cookies, nil
}

func (d *DownloaderBaseRepo) ClearCookies(ctx context.Context, key string) error {
	_, err := d.redisCli.Del(ctx, key).Result()
	if err != nil {
		d.log.Errorf("DownloaderBaseRepo.SaveCookie:redisCli.Get %s error:%v", key, err)
		return err
	}
	return nil
}

func (d *DownloaderBaseRepo) GetComponentName() string {
	return string(d.componentName)
}
