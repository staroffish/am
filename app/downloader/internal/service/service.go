package service

import (
	"github.com/google/wire"
	"github.com/staroffish/am/common/config"
)

// ProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(NewDownloaderService, config.NewHTTPServerConfig, config.NewGrpcServerConfig)
