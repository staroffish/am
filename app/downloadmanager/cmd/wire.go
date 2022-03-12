//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/staroffish/am/app/downloadmanager/internal/biz"
	"github.com/staroffish/am/app/downloadmanager/internal/data"
	"github.com/staroffish/am/app/downloadmanager/internal/server"
	"github.com/staroffish/am/app/downloadmanager/internal/service"
	etcd "go.etcd.io/etcd/client/v3"

	"github.com/staroffish/am/common/config"
)

// initApp init kratos application.
func initApp(config.ComponentName, *etcd.Client, log.Logger, registry.Registrar, data.TaskEtcdPrefix, registry.Discovery) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
