package data

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/staroffish/am/api/common"
	pb "github.com/staroffish/am/api/downloadmanager/v1"
	"github.com/staroffish/am/app/spider/internal/biz"
	"github.com/staroffish/am/app/spider/internal/conf"
	dto "github.com/staroffish/am/common/dto/spider"
)

type amSpiderRepo struct {
	spiderConfig      *conf.SpiderConfig
	log               *log.Helper
	redisClient       *redis.Client
	downloadMangerCli pb.DownloadmanagerClient
}

func NewDownloadManagerClient(r registry.Discovery) pb.DownloadmanagerClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///download-manager"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return pb.NewDownloadmanagerClient(conn)
}

// NewGreeterRepo .
func NewAmSpiderRepo(c *conf.SpiderConfig, database *Data, confData *conf.SpiderServerConfig, downloadMangerCli pb.DownloadmanagerClient, logger log.Logger) biz.AnimeSpiderRepo {
	return &amSpiderRepo{
		spiderConfig:      c,
		log:               log.NewHelper(logger),
		downloadMangerCli: downloadMangerCli,
		redisClient:       database.Client,
	}
}

func (a *amSpiderRepo) GetWebContent(ctx context.Context) (webContent string, err error) {
	client := &http.Client{}
	conf := a.spiderConfig.GetConfig()
	a.log.Infof("webSpiderRepo.GetWebContent get config %v", conf)

	req, err := http.NewRequest(conf.Method, conf.GetUrl(), nil)
	if err != nil {
		return "", fmt.Errorf("GetWebContent:http.NewRequest error:%v", err)
	}

	// 由于go自己的 user-agent 可能会被屏蔽 所以 这里改成firefox的user-agent
	req.Header.Add("User-Agent", conf.GetUserAgent())

	// 如果存在使用代理
	proxy := conf.GetProxy()
	if proxy != "" {
		tr := &http.Transport{}
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			return "", fmt.Errorf("GetWebContent:url.Parse proxy error:%v", err)
		}
		tr.Proxy = http.ProxyURL(proxyUrl)
		client.Transport = tr
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("GetWebContent:http.DefaultClient.Do error:%v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GetWebContent:resp.StatusCode != OK status:%v", resp.StatusCode)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return "", fmt.Errorf("GetWebContent:io.Copy error:%v", err)
	}

	return buf.String(), nil
}

func (a *amSpiderRepo) SaveAnimeMagnets(ctx context.Context, animeMagnets []*dto.AnimeMagnet) error {
	today := time.Now().Format("2006-01-02")
	conf := a.spiderConfig.GetConfig()
	key := fmt.Sprintf("anime:link:%s", today)

	var retError error
	a.log.Infof("key=%s, timeout=%d", key, conf.AnimeMagnetTimeout)
	for _, animeMagnet := range animeMagnets {
		exists, err := a.redisClient.HExists(ctx, key, animeMagnet.Name).Result()
		if err != nil {
			a.log.Errorf("amSpiderRepo.redisClient.HExists %s %s error:%v", key, animeMagnet.Name, err)
			retError = err
			continue
		}

		if exists {
			continue
		}

		_, err = a.redisClient.HSet(ctx, key, animeMagnet.Name, animeMagnet.MagnetLink).Result()
		if err != nil {
			a.log.Errorf("amSpiderRepo.redisClient.HSet %s %s %s error:%v", key, animeMagnet.Name, animeMagnet.MagnetLink, err)
			retError = err
		}
	}

	ttl, err := a.redisClient.TTL(ctx, key).Result()
	if err != nil {
		a.log.Errorf("a.redisClient.TTL %s error:%v", key, err)
		retError = err
	}

	if ttl == -1 {
		expire := time.Duration(conf.AnimeMagnetTimeout*24*60*60) * time.Second
		_, err := a.redisClient.Expire(ctx, key, expire).Result()
		if err != nil {
			a.log.Errorf("a.redisClient.Expire %s %d error:%v", key, expire, err)
			retError = err
		}
	}

	go func() {
		timeOutctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		reply, err := a.downloadMangerCli.ScanTaskAndDownload(timeOutctx, &common.Empty{})
		if err != nil {
			a.log.Errorf("downloadMangerCli.ScanTaskAndDownload error: %v", err)
			return
		}
		a.log.Infof("downloadMangerCli.ScanTaskAndDownload:matched task count:%d", len(reply.CreatedTasks))
	}()

	return retError // 只是为了让上层感知到发生了错误
}
