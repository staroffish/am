package ad

import (
	"context"
	"log"
	"time"

	etcd "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	apiCommon "github.com/staroffish/am/api/common"
	spiderv1 "github.com/staroffish/am/api/spider/v1"
	"github.com/staroffish/am/global"

	"github.com/staroffish/am/rd"

	clientv3 "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
)

const DayDuration = 24 * time.Hour

// Ad - 自动下载
type Ad struct {
	manualAction chan struct{}
	config       *global.Config
	spiderClient spiderv1.SpiderClient
}

// New 创建 自动下载
func New(cfg *global.Config) (*Ad, error) {
	defer global.TraceLog("ad.New")()
	ad := Ad{config: cfg}
	if err := rd.InitDownloader(); err != nil {
		return nil, err
	}

	ad.manualAction = make(chan struct{}, 2)

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{global.Cfg.EtcdEndpoints},
		DialTimeout: time.Duration(global.Cfg.EtcdDialTimeout) * time.Second,
		DialOptions: []ggrpc.DialOption{ggrpc.WithBlock()},
	})
	if err != nil {
		global.Log.Errorf("ad:Run:connectRedis error:%v", err)
		log.Fatalf("ad:Run:connectRedis error:%v", err)
	}

	r := etcd.New(client)

	connSpider, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///spider"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		global.Log.Errorf("ad:Run:grpc.DialInsecure error:%v", err)
		log.Fatalf("ad:Run:grpc.DialInsecure error:%v", err)
	}

	ad.spiderClient = spiderv1.NewSpiderClient(connSpider)

	return &ad, nil
}

// 手动查看是否有可以下载的动画
func (ad *Ad) ManualCheckDownload() {
	go func() { ad.manualAction <- struct{}{} }()
}

func (ad *Ad) Run() {
	defer global.TraceLog("Ad.Run")()
	crawlerTicker := time.Tick(time.Duration(ad.config.AdInter) * time.Second)

	go func() {
		for {
			select {
			case <-ad.manualAction:
				global.Log.Infof("Start manual action.")
			case <-crawlerTicker:
				global.Log.Infof("Start autodownload.")
			}

			spiderCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			_, err := ad.spiderClient.Crawl(spiderCtx, &apiCommon.Empty{})
			cancel()
			if err != nil {
				global.Log.Infof("ad.Run:spiderClient.Crawl error:%v", err)
				continue
			}
		}
	}()
	// 在最后出发一次网页抓取
	ad.ManualCheckDownload()
}
