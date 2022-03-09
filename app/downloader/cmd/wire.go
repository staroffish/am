//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/staroffish/am/app/downloader/internal/biz"

	"github.com/staroffish/am/app/downloader/internal/data"
	"github.com/staroffish/am/app/downloader/internal/server"
	"github.com/staroffish/am/app/downloader/internal/service"
	etcd "go.etcd.io/etcd/client/v3"

	"github.com/staroffish/am/common/config"
)

// initApp init kratos application.
func initApp(config.ComponentName, *etcd.Client, log.Logger, registry.Registrar) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.DataProviderSet, biz.BizProviderSet, service.ServiceProviderSet, newApp))
}
