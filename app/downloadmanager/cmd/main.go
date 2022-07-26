package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	etcd "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/staroffish/am/app/downloadmanager/internal/conf"
	commonConfig "github.com/staroffish/am/common/config"
	"github.com/staroffish/am/common/util"
	clientv3 "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	// component name
	componentName string

	id, _ = os.Hostname()

	componentType = "download-manager"
)

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&componentName, "name", "", "component name, eg: -name download-manager")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(componentName),
		kratos.Name(componentType),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(r),
	)
}

func main() {
	flag.Parse()
	if componentName == "" {
		flag.Usage()
		os.Exit(-1)
	}

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.Caller(4),
		"service.id", id,
		"service.name", componentName,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	log := log.NewHelper(logger)

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	log.Infof("Bootstrap loaded config:%v", bc.Data)

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   bc.Data.Etcd.EndPoints,
		DialTimeout: time.Duration(bc.Data.Etcd.DialTimeout) * time.Second,
		DialOptions: []ggrpc.DialOption{ggrpc.WithBlock()},
		Username:    bc.Data.Etcd.User,
		Password:    bc.Data.Etcd.Password,
	})
	if err != nil {
		log.Errorf("New etcd client error:%v", err)
		os.Exit(-1)
	}

	prefixString := util.TaskEtcdPrefix(fmt.Sprintf("/%s/%s/tasks", componentType, componentName))

	registry := etcd.New(client)

	app, cleanup, err := initApp(commonConfig.ComponentName(componentName), commonConfig.Version(Version), client, registry, prefixString, registry)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
