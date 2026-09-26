package downloader

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/staroffish/am/internal/config"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/util"
)

type QBittorrentClient struct {
	cfg          config.QBittorrentConfig
	log          *log.Logger
	cookieMu     sync.Mutex
	cookies      map[string]string
	cookieExpire time.Time
}

const boundary = "---------------------------6688794727912"

// cookieTTL 与原来 Redis 里 cookie 的过期时间保持一致。
const cookieTTL = 5 * time.Minute

func NewQBittorrent(cfg config.QBittorrentConfig, logger *log.Logger) *QBittorrentClient {
	return &QBittorrentClient{
		cfg: cfg,
		log: logger,
	}
}

func (c *QBittorrentClient) Add(ctx context.Context, link, storePath string) error {
	url := fmt.Sprintf("%s/api/v2/torrents/add", c.cfg.URL)
	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", boundary)

	body := []string{}
	body = append(body, fmt.Sprintf("--%s\r\n", boundary))
	body = append(body, "Content-Disposition: form-data; name=\"urls\"\r\n\r\n")
	body = append(body, fmt.Sprintf("%s\r\n", link))
	body = append(body, fmt.Sprintf("--%s\r\n", boundary))
	body = append(body, "Content-Disposition: form-data; name=\"savepath\"\r\n\r\n")
	body = append(body, fmt.Sprintf("%s\r\n", storePath))
	body = append(body, fmt.Sprintf("--%s--\r\n", boundary))

	return c.noResponseHTTPRequest(ctx, http.MethodPost, url, strings.Join(body, ""),
		http.Header{"content-type": []string{contentType}})
}

func (c *QBittorrentClient) Delete(ctx context.Context, hash string) error {
	url := fmt.Sprintf("%s/api/v2/torrents/delete", c.cfg.URL)
	return c.noResponseHTTPRequest(ctx, http.MethodPost, url,
		fmt.Sprintf("hashes=%s&deleteFiles=false", hash),
		http.Header{"content-type": []string{"application/x-www-form-urlencoded"}})
}

func (c *QBittorrentClient) Pause(ctx context.Context, hash string) error {
	url := fmt.Sprintf("%s/api/v2/torrents/pause", c.cfg.URL)
	return c.noResponseHTTPRequest(ctx, http.MethodPost, url,
		fmt.Sprintf("hashes=%s", hash),
		http.Header{"content-type": []string{"application/x-www-form-urlencoded"}})
}

func (c *QBittorrentClient) Resume(ctx context.Context, hash string) error {
	url := fmt.Sprintf("%s/api/v2/torrents/resume", c.cfg.URL)
	return c.noResponseHTTPRequest(ctx, http.MethodPost, url,
		fmt.Sprintf("hashes=%s", hash),
		http.Header{"content-type": []string{"application/x-www-form-urlencoded"}})
}

func (c *QBittorrentClient) List(ctx context.Context) ([]model.TorrentInfo, error) {
	cookies, err := c.login(ctx)
	if err != nil {
		return nil, err
	}

	listURL := fmt.Sprintf("%s/api/v2/torrents/info", c.cfg.URL)
	resp, err := util.SendHTTPRequest(ctx, http.MethodGet, listURL, "", toHTTPCookies(cookies), nil)
	if err != nil {
		c.clearCookies()
		return nil, err
	}
	defer resp.Body.Close()

	buff := bytes.NewBuffer([]byte{})
	if _, err := io.Copy(buff, resp.Body); err != nil {
		return nil, err
	}

	type qbInfo struct {
		Name        string  `json:"name"`
		Size        int64   `json:"size"`
		Progress    float32 `json:"progress"`
		State       string  `json:"state"`
		StorePath   string  `json:"save_path"`
		CreatedTime int64   `json:"added_on"`
		Hash        string  `json:"hash"`
	}

	var torrentInfoList []qbInfo
	if err := json.Unmarshal(buff.Bytes(), &torrentInfoList); err != nil {
		return nil, err
	}

	taskInfos := make([]model.TorrentInfo, 0, len(torrentInfoList))
	for _, info := range torrentInfoList {
		taskInfos = append(taskInfos, model.TorrentInfo{
			Hash:        info.Hash,
			Name:        info.Name,
			Size:        info.Size,
			Progress:    info.Progress,
			Status:      convertState(info.State),
			StorePath:   info.StorePath,
			CreatedTime: time.Unix(info.CreatedTime, 0).Format("2006-01-02 15:04:05"),
		})
	}

	return taskInfos, nil
}

func (c *QBittorrentClient) login(ctx context.Context) (map[string]string, error) {
	c.cookieMu.Lock()
	defer c.cookieMu.Unlock()

	if len(c.cookies) > 0 && time.Now().Before(c.cookieExpire) {
		return c.cookies, nil
	}
	c.cookies = nil

	loginURL := fmt.Sprintf("%s/api/v2/auth/login", c.cfg.URL)
	body := fmt.Sprintf("username=%s&password=%s", c.cfg.Username, c.cfg.Password)

	resp, err := util.SendHTTPRequest(ctx, http.MethodPost, loginURL, body, nil,
		http.Header{"content-type": []string{"application/x-www-form-urlencoded"}})
	if err != nil {
		c.log.Printf("qBittorrent login error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	respStr := strings.TrimSpace(string(b))
	c.log.Printf("qBittorrent login: %s", respStr)

	if respStr != "Ok." {
		return nil, fmt.Errorf("qBittorrent login failed: %s", respStr)
	}

	cookieMap := make(map[string]string)
	for _, ck := range resp.Cookies() {
		cookieMap[ck.Name] = ck.Value
	}

	c.cookies = cookieMap
	c.cookieExpire = time.Now().Add(cookieTTL)

	return cookieMap, nil
}

func (c *QBittorrentClient) clearCookies() {
	c.cookieMu.Lock()
	c.cookies = nil
	c.cookieExpire = time.Time{}
	c.cookieMu.Unlock()
}

func (c *QBittorrentClient) noResponseHTTPRequest(ctx context.Context, method, url, body string, headers http.Header) error {
	cookies, err := c.login(ctx)
	if err != nil {
		c.clearCookies()
		c.log.Printf("qBittorrent login error: %v", err)
		return err
	}

	resp, err := util.SendHTTPRequest(ctx, method, url, body, toHTTPCookies(cookies), headers)
	if err != nil {
		c.log.Printf("qBittorrent %s %s error: %v", method, url, err)
		return err
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	respStr := strings.TrimSpace(string(b))
	c.log.Printf("qBittorrent %s %s → %s", method, url, respStr)

	if respStr != "Ok." && respStr != "" {
		return fmt.Errorf("qBittorrent %s failed: %s", method, respStr)
	}
	return nil
}

func toHTTPCookies(cookies map[string]string) []*http.Cookie {
	var result []*http.Cookie
	for name, value := range cookies {
		result = append(result, &http.Cookie{Name: name, Value: value})
	}
	return result
}

func convertState(state string) string {
	lower := strings.ToLower(state)
	switch {
	case strings.Contains(lower, "queued"):
		return "queued"
	case strings.Contains(lower, "checking"):
		return "checking"
	case strings.Contains(lower, "paused"):
		return "paused"
	case strings.Contains(lower, "up"):
		return "seeding"
	case strings.Contains(lower, "dl") || strings.Contains(lower, "downloading"):
		return "downloading"
	default:
		return state
	}
}
