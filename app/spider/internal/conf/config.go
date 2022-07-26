package conf

import (
	"context"
	"fmt"
	"path"
	"strconv"
	"sync"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/staroffish/am/common/config"
	"github.com/staroffish/am/common/util"
	etcd "go.etcd.io/etcd/client/v3"
)

type SpiderConfig struct {
	spiderParameter *SpiderParameter
	rwLock          *sync.RWMutex
	componentName   string
	componentType   string
	log             log.Helper
	etcdCli         *etcd.Client
}

type SpiderServerConfig struct {
	Http  *config.HttpServerConfig
	Grpc  *config.GrpcServerConfig
	Redis *config.ReidsConfig
}

func NewSpiderServerConfig(http *config.HttpServerConfig, grpc *config.GrpcServerConfig, redis *config.ReidsConfig) *SpiderServerConfig {
	return &SpiderServerConfig{
		Http:  http,
		Grpc:  grpc,
		Redis: redis,
	}
}

func NewSpiderConfig(componentName config.ComponentName, componentType config.ComponentType, etcdCli *etcd.Client, logger log.Logger) (*SpiderConfig, error) {
	spiderConfig := &SpiderConfig{
		componentName:   string(componentName),
		componentType:   string(componentType),
		spiderParameter: &SpiderParameter{},
		rwLock:          &sync.RWMutex{},
		log:             *log.NewHelper(logger),
		etcdCli:         etcdCli,
	}

	if err := spiderConfig.syncSpiderConfig(); err != nil {
		return nil, fmt.Errorf("NewSpiderConfig error %v", err)
	}

	etcdWather := util.NewEtcdWatcher(etcdCli, util.TaskEtcdPrefix(fmt.Sprintf("/%s/%s", componentType, componentName)), logger)
	go etcdWather.Watch(func(_ []*etcd.Event) error {
		if err := spiderConfig.syncSpiderConfig(); err != nil {
			return fmt.Errorf("etcdWather.Watch: %v", err)
		}
		return nil
	})

	return spiderConfig, nil
}

func (s *SpiderConfig) GetConfig() *SpiderParameter {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.spiderParameter
}

func (s *SpiderConfig) syncSpiderConfig() error {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	prefix := fmt.Sprintf("/%s/%s", s.componentType, s.componentName)
	resp, err := s.etcdCli.Get(context.Background(), prefix, etcd.WithPrefix())
	if err != nil {
		return fmt.Errorf("SpiderConfig.syncSpiderConfig:client.Get %s --prefix error:%v", prefix, err)
	}

	for _, item := range resp.Kvs {
		keyBase := path.Base(string(item.Key))
		switch keyBase {
		case "anime_magnet_timeout":
			value, err := strconv.ParseInt(string(item.Value), 10, 64)
			if err != nil {
				s.log.Errorf("SpiderConfig.syncSpiderConfig:anime_magnet_timeout: parse int error %v", err)
				continue
			}
			s.spiderParameter.AnimeMagnetTimeout = value
		case "interval":
			value, err := strconv.ParseInt(string(item.Value), 10, 64)
			if err != nil {
				s.log.Errorf("SpiderConfig.syncSpiderConfig:interval: parse int error %v", err)
				continue
			}
			s.spiderParameter.Interval = value
		case "method":
			s.spiderParameter.Method = string(item.Value)
		case "proxy":
			s.spiderParameter.Proxy = string(item.Value)
		case "url":
			s.spiderParameter.Url = string(item.Value)
		case "user_agent":
			s.spiderParameter.UserAgent = string(item.Value)
		case "type":
			// 类型不能更改 只有为空的时候才赋值
			if s.spiderParameter.Type == "" {
				s.spiderParameter.Type = string(item.Value)
			}
		}
	}
	s.log.Infof("spider config changed: %v", s.spiderParameter)
	return nil
}
