package conf

import "github.com/staroffish/am/common/config"

type DownloadManagerServerConfig struct {
	Http  *config.HttpServerConfig
	Grpc  *config.GrpcServerConfig
	Redis *config.ReidsConfig
}

func NewDownloadManagerServerConfig(http *config.HttpServerConfig, grpc *config.GrpcServerConfig, redis *config.ReidsConfig) *DownloadManagerServerConfig {
	return &DownloadManagerServerConfig{
		Http:  http,
		Grpc:  grpc,
		Redis: redis,
	}
}
