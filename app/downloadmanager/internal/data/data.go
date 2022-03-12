package data

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/staroffish/am/app/downloadmanager/internal/conf"
	"github.com/staroffish/am/common/config"
	"github.com/staroffish/am/common/util"
	etcd "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDownloadManagerRepo, NewDownloaderClient, NewDownloadTask, util.NewReidsClient, config.NewRedisConfig, conf.NewDownloadManagerServerConfig)

// Data .
type Data struct {
	redisCli *redis.Client
	etcdCli  *etcd.Client
}

// NewData .
func NewData(c *conf.DownloadManagerServerConfig, etcdCli *etcd.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Username:     c.Redis.User,
		Password:     c.Redis.Password,
		ReadTimeout:  time.Duration(c.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.Redis.WriteTimeout) * time.Second,
		DB:           int(c.Redis.Db),
	})
	_, err := redisClient.Get(context.Background(), "").Result()
	if err != nil && err != redis.Nil {
		log.NewHelper(logger).Errorf("redisClient.Get error %v", err)
		panic(fmt.Sprintf("redisClient.Get error %v", err))
	}
	return &Data{
		redisCli: redisClient,
		etcdCli:  etcdCli,
	}, cleanup, nil
}
