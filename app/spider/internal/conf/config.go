package conf

import (
	"sync"

	"github.com/staroffish/am/common/config"
)

type SpiderConfig struct {
	spiderParameter *SpiderParameter
	rwLock          *sync.RWMutex
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

func NewSpiderConfig(spiderConfig *SpiderParameter) *SpiderConfig {
	return &SpiderConfig{
		spiderParameter: spiderConfig,
		rwLock:          &sync.RWMutex{},
	}
}

func (s *SpiderConfig) GetConfig() *SpiderParameter {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.spiderParameter
}

func (s *SpiderConfig) SetConfig(config *SpiderParameter) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.copySpiderConfig(config)
}

func (s *SpiderConfig) copySpiderConfig(config *SpiderParameter) {
	s.spiderParameter.Method = config.Method
	s.spiderParameter.Proxy = config.Proxy
	s.spiderParameter.Url = config.Url
	s.spiderParameter.Interval = config.Interval
	s.spiderParameter.AnimeMagnetTimeout = config.AnimeMagnetTimeout
}
