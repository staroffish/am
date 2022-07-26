package conf

import (
	"github.com/staroffish/am/common/config"
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
