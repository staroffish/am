package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	etcdCfg "github.com/go-kratos/kratos/contrib/config/etcd/v2"
	etcd "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/staroffish/am/app/am-spider/internal/conf"
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
)

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&componentName, "name", "", "component name, eg: -name miobt-spider")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(componentName),
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
	log.Infof("Bootstrap loaded server config:%v", bc.Server)

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

	etcdSource, err := etcdCfg.New(client, etcdCfg.WithPath(fmt.Sprintf("/spider/%s", componentName)), etcdCfg.WithPrefix(true))
	if err != nil {
		log.Errorf("New etcd source error:%v", err)
		os.Exit(-1)
	}

	etcdConfig := config.New(
		config.WithSource(
			etcdSource,
		),
	)
	defer etcdConfig.Close()

	if err := etcdConfig.Load(); err != nil {
		log.Errorf("load etcd config error:%v", err)
		os.Exit(-1)
	}

	spiderParam, err := initSpiderParameters(etcdConfig, componentName)
	if err != nil {
		log.Errorf("load spider config error:%v", err)
		os.Exit(-1)
	}

	log.Infof("Bootstrap loaded spider config:%v", spiderParam)

	spiderConfig := conf.NewSpiderConfig(spiderParam)

	err = etcdConfig.Watch(fmt.Sprintf("/spider/%s", componentName), func(key string, v config.Value) {
		spiderParam, err = initSpiderParameters(etcdConfig, componentName)
		if err != nil {
			log.Errorf("load spider config error:%v", err)
		}
		log.Infof("loaded spider config:%v", spiderParam)
		spiderConfig.SetConfig(spiderParam)
	})
	if err != nil {
		log.Errorf("etcdConfig.Watch error:%v", err)
		os.Exit(-1)
	}

	app, cleanup, err := initApp(bc.Server, bc.Data, spiderConfig, logger, etcd.New(client))
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func initSpiderParameters(etcdConfig config.Config, conponentName string) (*conf.SpiderParameter, error) {
	prefix := fmt.Sprintf("/spider/%s", conponentName)
	var err error

	spiderConf := &conf.SpiderParameter{}
	spiderConf.Url, err = etcdConfig.Value(fmt.Sprintf("%s/%s", prefix, "url")).String()
	if err != nil {
		return nil, fmt.Errorf("get %s error:%v", fmt.Sprintf("%s/%s", prefix, "url"), err)
	}
	spiderConf.Proxy, _ = etcdConfig.Value(fmt.Sprintf("%s/%s", prefix, "proxy")).String()

	spiderConf.Method, err = etcdConfig.Value(fmt.Sprintf("%s/%s", prefix, "method")).String()
	if err != nil {
		return nil, fmt.Errorf("get %s error:%v", fmt.Sprintf("%s/%s", prefix, "method"), err)
	}
	spiderConf.Type, err = etcdConfig.Value(fmt.Sprintf("%s/%s", prefix, "type")).String()
	if err != nil {
		return nil, fmt.Errorf("get %s error:%v", fmt.Sprintf("%s/%s", prefix, "type"), err)
	}
	spiderConf.Interval, err = etcdConfig.Value(fmt.Sprintf("%s/%s", prefix, "interval")).String()
	if err != nil {
		return nil, fmt.Errorf("get %s error:%v", fmt.Sprintf("%s/%s", prefix, "interval"), err)
	}
	spiderConf.UserAgent, err = etcdConfig.Value(fmt.Sprintf("%s/%s", prefix, "user-agent")).String()
	if err != nil {
		return nil, fmt.Errorf("get %s error:%v", fmt.Sprintf("%s/%s", prefix, "user-agent"), err)
	}

	return spiderConf, err
}
