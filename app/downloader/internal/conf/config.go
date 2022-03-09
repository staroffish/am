package conf

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/staroffish/am/common/config"
	etcd "go.etcd.io/etcd/client/v3"
)

type DownloaderServerConfig struct {
	Http  *config.HttpServerConfig
	Grpc  *config.GrpcServerConfig
	Redis *config.ReidsConfig
}

func NewDownloaderServerConfig(http *config.HttpServerConfig, grpc *config.GrpcServerConfig, redis *config.ReidsConfig) *DownloaderServerConfig {
	return &DownloaderServerConfig{
		Http:  http,
		Grpc:  grpc,
		Redis: redis,
	}
}

type DownloaderConfig struct {
	UserName      string
	Password      string
	componentName config.ComponentName
	log           *log.Helper
	etcdCli       *etcd.Client
	mutex         *sync.RWMutex
}

func NewDownloaderConfig(etcdCli *etcd.Client, componentName config.ComponentName, logger log.Logger) *DownloaderConfig {
	return &DownloaderConfig{
		log:           log.NewHelper(logger),
		etcdCli:       etcdCli,
		componentName: componentName,
		mutex:         &sync.RWMutex{},
	}
}

func (c *DownloaderConfig) Get() (string, string) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.UserName, c.Password
}

func (c *DownloaderConfig) Watch() {
	prefix := fmt.Sprintf("/downloader/%s", c.componentName)
	watchChan := c.etcdCli.Watch(context.Background(), prefix, etcd.WithPrefix())

	for range watchChan {
		resp, err := c.etcdCli.Get(context.Background(), prefix, etcd.WithPrefix())
		if err != nil {
			c.log.Errorf("DownloaderConfig.Watch: etcdCli.Get %s error:%v", prefix, err)
		}
		c.mutex.Lock()
		for _, kv := range resp.Kvs {
			switch string(kv.Key) {
			case fmt.Sprintf("%s%s", prefix, "user"):
				c.UserName = string(kv.Value)
			case fmt.Sprintf("%s%s", prefix, "password"):
				c.Password = string(kv.Value)
			}
		}
		c.mutex.Unlock()
	}
}
