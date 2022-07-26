//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/staroffish/am/app/spider/internal/biz"
	"github.com/staroffish/am/app/spider/internal/data"
	"github.com/staroffish/am/app/spider/internal/server"
	"github.com/staroffish/am/app/spider/internal/service"
	"github.com/staroffish/am/common"
	"github.com/staroffish/am/common/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// initApp init kratos application.
func initApp(config.ComponentName, config.ComponentType, config.Version, *clientv3.Client, registry.Registrar, registry.Discovery) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, common.ProviderSet, newApp))
}
