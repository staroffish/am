package qbittorrent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/staroffish/am/app/downloader/internal/biz"
	"github.com/staroffish/am/app/downloader/internal/data/downloader"
	"github.com/staroffish/am/common/util"
	etcd "go.etcd.io/etcd/client/v3"
)

type qbittorrentConfig struct {
	userName string
	password string
	url      string
}

type qbittorrentInfo struct {
	Name        string  `json:"name"`
	Size        int64   `json:"size"`
	Progress    float32 `json:"progress"`
	State       string  `json:"state"`
	StorePath   string  `json:"save_path"`
	CreatedTime int64   `json:"added_on"`
	Hash        string  `json:"hash"`
}

type QbittorrentDownloaderRepo struct {
	log *log.Helper
	*qbittorrentConfig
	mutex   *sync.RWMutex
	etcdCli *etcd.Client
	*downloader.DownloaderBaseRepo
}

const BOUNDARY = "---------------------------6688794727912"

func NewQbittorrentDownloaderRepo(baseRepo *downloader.DownloaderBaseRepo, etcdCli *etcd.Client, logger log.Logger) *QbittorrentDownloaderRepo {
	repo := &QbittorrentDownloaderRepo{
		log:                log.NewHelper(logger),
		DownloaderBaseRepo: baseRepo,
		etcdCli:            etcdCli,
		mutex:              &sync.RWMutex{},
		qbittorrentConfig:  &qbittorrentConfig{},
	}
	if err := repo.syncConfig(); err != nil {
		panic(fmt.Sprintf("NewQbittorrentDownloaderRepo:repo.syncConfig error: %v", err))
	}
	go func() {
		repo.watchConfig()
	}()
	return repo
}

func (d *QbittorrentDownloaderRepo) Add(ctx context.Context, link, storePath string) error {
	qbitConfig := d.getConfig()

	url := fmt.Sprintf("%s/api/v2/torrents/add", qbitConfig.url)

	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", BOUNDARY)

	body := []string{}
	body = append(body, fmt.Sprintf("--%s\r\n", BOUNDARY))
	body = append(body, "Content-Disposition: form-data; name=\"urls\"\r\n\r\n")
	body = append(body, fmt.Sprintf("%s\r\n", link))
	body = append(body, fmt.Sprintf("--%s\r\n", BOUNDARY))
	body = append(body, "Content-Disposition: form-data; name=\"savepath\"\r\n\r\n")
	body = append(body, fmt.Sprintf("%s\r\n", storePath))
	body = append(body, fmt.Sprintf("--%s--\r\n", BOUNDARY))

	bodyStr := strings.Join(body, "")

	d.log.Infof("url=%s, contentType=%s, body=%s", url, contentType, bodyStr)

	if err := d.noResponseHttpRequest(ctx, http.MethodPost, url, bodyStr, http.Header{"content-type": []string{contentType}}); err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.Pause: %v", err)
		return err
	}

	return nil
}

func (d *QbittorrentDownloaderRepo) Delete(ctx context.Context, id string) error {
	qbitConfig := d.getConfig()

	url := fmt.Sprintf("%s/api/v2/torrents/delete", qbitConfig.url)

	if err := d.noResponseHttpRequest(ctx, http.MethodPost, url, fmt.Sprintf("hashes=%s&deleteFiles=false", id), http.Header{"content-type": []string{"application/x-www-form-urlencoded"}}); err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.Pause: %v", err)
		return err
	}

	return nil
}

func (d *QbittorrentDownloaderRepo) List(ctx context.Context) ([]*biz.DownloaderTaskInfo, error) {
	qbitConfig := d.getConfig()

	cookies, err := d.login(ctx)
	if err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.List:login %v", err)
		return nil, err
	}

	listUrl := fmt.Sprintf("%s/api/v2/torrents/info", qbitConfig.url)
	resp, err := util.SendHttpRequest(ctx, http.MethodGet, listUrl, "", cookies, nil)
	if err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.List:SendHttpRequest %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	buff := bytes.NewBuffer([]byte{})

	n, err := io.Copy(buff, resp.Body)
	if err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.List:io.Copy %v", err)
		return nil, err
	}
	d.log.Infof("read from listUrl %d bytes", n)

	torrentInfoList := []qbittorrentInfo{}

	if err := json.Unmarshal(buff.Bytes(), &torrentInfoList); err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.List:json.Unmarshal %v", err)
		return nil, err
	}

	taskInfos := []*biz.DownloaderTaskInfo{}
	for _, torrentInfo := range torrentInfoList {
		taskInfo := &biz.DownloaderTaskInfo{
			Name:      torrentInfo.Name,
			Status:    convertState(torrentInfo.State),
			Progress:  torrentInfo.Progress,
			Size:      torrentInfo.Size,
			StorePath: torrentInfo.StorePath,
			Id:        torrentInfo.Hash,
		}
		taskInfo.CreatedTime = time.Unix(torrentInfo.CreatedTime, 0).Format("2006-01-02 15:04:05")
		taskInfos = append(taskInfos, taskInfo)
	}

	return taskInfos, nil
}

func (d *QbittorrentDownloaderRepo) Pause(ctx context.Context, id string) error {
	qbitConfig := d.getConfig()

	url := fmt.Sprintf("%s/api/v2/torrents/pause", qbitConfig.url)

	if err := d.noResponseHttpRequest(ctx, http.MethodPost, url, fmt.Sprintf("hashes=%s", id), http.Header{"content-type": []string{"application/x-www-form-urlencoded"}}); err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.Pause: %v", err)
		return err
	}

	return nil
}

func (d *QbittorrentDownloaderRepo) Resume(ctx context.Context, id string) error {
	qbitConfig := d.getConfig()

	url := fmt.Sprintf("%s/api/v2/torrents/resume", qbitConfig.url)

	if err := d.noResponseHttpRequest(ctx, http.MethodPost, url, fmt.Sprintf("hashes=%s", id), http.Header{"content-type": []string{"application/x-www-form-urlencoded"}}); err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.Pause: %v", err)
		return err
	}

	return nil
}

func (d *QbittorrentDownloaderRepo) login(ctx context.Context) ([]*http.Cookie, error) {
	qbitConfig := d.getConfig()

	cookies, err := d.DownloaderBaseRepo.LoadCookies(ctx, d.getCookieCacheName())
	if err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.List:LoadCookies %v", err)
	}
	if len(cookies) > 0 {
		return cookies, nil
	}

	loginUrl := fmt.Sprintf("%s/api/v2/auth/login", qbitConfig.url)
	body := fmt.Sprintf("username=%s&password=%s", qbitConfig.userName, qbitConfig.password)

	resp, err := util.SendHttpRequest(ctx, http.MethodPost, loginUrl, body, nil, http.Header{"content-type": []string{"application/x-www-form-urlencoded"}})
	if err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.login: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	cookies = resp.Cookies()

	if err := d.SaveCookies(ctx, d.getCookieCacheName(), cookies); err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.SaveCookies: %v", err)
	}
	d.log.Infof("QbittorrentDownloaderRepo login succeeded")
	return cookies, nil
}

func (d *QbittorrentDownloaderRepo) getCookieCacheName() string {
	return fmt.Sprintf("downloader:qbittorrent:%s:cookie:login", d.GetComponentName())
}

func (d *QbittorrentDownloaderRepo) getConfig() *qbittorrentConfig {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	qbitConfig := *d.qbittorrentConfig
	return &qbitConfig
}

func (d *QbittorrentDownloaderRepo) watchConfig() {
	prefix := fmt.Sprintf("/downloader/%s", d.GetComponentName())
	watchChan := d.etcdCli.Watch(context.Background(), prefix, etcd.WithPrefix())

	for range watchChan {
		d.syncConfig()
	}
}

func (d *QbittorrentDownloaderRepo) syncConfig() error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	prefix := fmt.Sprintf("/downloader/%s", d.GetComponentName())
	d.log.Infof("prefix=%s\n", prefix)
	resp, err := d.etcdCli.Get(context.Background(), prefix, etcd.WithPrefix())
	if err != nil {
		d.log.Errorf("DownloaderConfig.Watch: etcdCli.Get %s error:%v", prefix, err)
		return err
	}
	for _, kv := range resp.Kvs {
		switch string(kv.Key) {
		case fmt.Sprintf("%s/%s", prefix, "username"):
			d.userName = string(kv.Value)
		case fmt.Sprintf("%s/%s", prefix, "password"):
			d.password = string(kv.Value)
		case fmt.Sprintf("%s/%s", prefix, "url"):
			d.url = string(kv.Value)
		}
	}
	if d.url == "" || d.userName == "" || d.password == "" {
		d.log.Errorf("DownloaderConfig.Watch: url or username or password is empty")
		return fmt.Errorf("DownloaderConfig.Watch: url or username or password is empty")
	}
	return nil
}

func (d *QbittorrentDownloaderRepo) noResponseHttpRequest(ctx context.Context, method, url, body string, headers http.Header) error {
	cookies, err := d.login(ctx)
	if err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.Pause:login %v", err)
		return err
	}

	resp, err := util.SendHttpRequest(ctx, method, url, body, cookies, headers)
	if err != nil {
		d.log.Errorf("QbittorrentDownloaderRepo.noResponseHttpRequest:SendHttpRequest %v", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

func convertState(state string) string {
	lowerState := strings.ToLower(state)
	if strings.Contains(lowerState, "queued") {
		return "queued"
	}

	if strings.Contains(lowerState, "checking") {
		return "checking"
	}

	if strings.Contains(lowerState, "paused") {
		return "paused"
	}

	if strings.Contains(lowerState, "paused") {
		return "paused"
	}

	if strings.Contains(lowerState, "up") {
		return "seeding"
	}

	if strings.Contains(lowerState, "dl") || strings.Contains(lowerState, "downloading") {
		return "downloading"
	}

	return state
}
