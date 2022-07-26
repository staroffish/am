package data

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/staroffish/am/app/downloader/internal/biz"
	"github.com/staroffish/am/app/downloader/internal/conf"
	"github.com/staroffish/am/app/downloader/internal/data/downloader"
	"github.com/staroffish/am/app/downloader/internal/data/downloader/qbittorrent"
	"github.com/staroffish/am/common/config"
	"github.com/staroffish/am/common/util"
	etcd "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var DataProviderSet = wire.NewSet(NewData, NewDownloadRepo, util.NewReidsClient, config.NewRedisConfig, conf.NewDownloaderServerConfig)

// Data .
type Data struct {
	redisCli *redis.Client
	etcdCli  *etcd.Client
}

type downloaderRepoWithSyncConfig interface {
	biz.DownloaderRepo
	SyncConfig() error
}

// NewData .
func NewData(c *conf.DownloaderServerConfig, redisClient *redis.Client, etcdCli *etcd.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		redisCli: redisClient,
		etcdCli:  etcdCli,
	}, cleanup, nil
}

func NewDownloadRepo(database *Data, componentName config.ComponentName, logger log.Logger) (biz.DownloaderRepo, error) {
	key := fmt.Sprintf("/downloader/%s/type", componentName)
	resp, err := database.etcdCli.Get(context.Background(), key)
	if err != nil {
		return nil, fmt.Errorf("NewDownloadRepo: database.etcdCli.Get %s error:%v", key, err)
	}
	if resp.Count < 1 {
		return nil, fmt.Errorf("NewDownloadRepo: database.etcdCli.Get %s empty", key)
	}

	downloaderType, err := strconv.ParseInt(string(resp.Kvs[0].Value), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("NewDownloadRepo: downloader type %s dose not numbric", resp.Kvs[0].Value)
	}

	etcdWatcher := util.NewEtcdWatcher(database.etcdCli, util.TaskEtcdPrefix(fmt.Sprintf("/downloader/%s", componentName)), logger)
	var repo downloaderRepoWithSyncConfig

	switch downloaderType {
	case downloader.QBITTORENT:
		repo = qbittorrent.NewQbittorrentDownloaderRepo(downloader.NewDownloaderBaseRepo(componentName, database.redisCli, logger), database.etcdCli, logger)
	default:
		err = fmt.Errorf("NewDownloadRepo: incorrect downloader type %d", downloaderType)
	}

	go func() {
		etcdWatcher.Watch(func(_ []*etcd.Event) error {
			repo.SyncConfig()
			return nil
		})
	}()
	return repo, err
}
