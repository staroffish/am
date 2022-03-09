package service

import (
	"github.com/google/wire"
	"github.com/staroffish/am/common/config"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewDownloadmanagerService, config.NewHTTPServerConfig, config.NewGrpcServerConfig)
