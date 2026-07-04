package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/staroffish/am/internal/config"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/spider"
	"github.com/staroffish/am/internal/store"
)

type SpiderService struct {
	cfg           []config.SpiderConfig
	redis         *store.RedisClient
	dmSvc         *DownloadManagerService
	log           *log.Logger
	status        map[string]*SpiderStatus
	statusMu      sync.RWMutex
	magnetTimeout int
}

type SpiderStatus struct {
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Running    bool      `json:"running"`
	LastRunAt  time.Time `json:"last_run_at"`
	LastError  string    `json:"last_error,omitempty"`
}

func NewSpiderService(cfg []config.SpiderConfig, redis *store.RedisClient, dmSvc *DownloadManagerService, logger *log.Logger, magnetTimeout int) *SpiderService {
	status := make(map[string]*SpiderStatus)
	for _, c := range cfg {
		status[c.Name] = &SpiderStatus{Name: c.Name, Type: c.Type}
	}
	return &SpiderService{
		cfg:           cfg,
		redis:         redis,
		dmSvc:         dmSvc,
		log:           logger,
		status:        status,
		magnetTimeout: magnetTimeout,
	}
}

func (s *SpiderService) GetStatus() []*SpiderStatus {
	s.statusMu.RLock()
	defer s.statusMu.RUnlock()
	var result []*SpiderStatus
	for _, c := range s.cfg {
		if st, ok := s.status[c.Name]; ok {
			result = append(result, st)
		}
	}
	return result
}

func (s *SpiderService) CrawlAll(ctx context.Context) {
	for _, c := range s.cfg {
		go s.Crawl(ctx, c.Name)
	}
}

func (s *SpiderService) Crawl(ctx context.Context, name string) {
	cfg := s.findConfig(name)
	if cfg == nil {
		s.log.Printf("spider %s not found", name)
		return
	}

	s.setRunning(name, true, "")
	defer s.setRunning(name, false, "")

	s.log.Printf("spider %s: start crawling %s", name, cfg.URL)

	webContent, err := s.fetch(ctx, cfg)
	if err != nil {
		s.setRunning(name, true, fmt.Sprintf("fetch error: %v", err))
		s.log.Printf("spider %s: %v", name, err)
		return
	}

	sp := spider.New(cfg.Type, s.log)
	if sp == nil {
		s.setRunning(name, true, fmt.Sprintf("unknown spider type: %s", cfg.Type))
		return
	}

	magnets, err := sp.ExtractData(ctx, webContent)
	if err != nil {
		s.setRunning(name, true, fmt.Sprintf("extract error: %v", err))
		s.log.Printf("spider %s: %v", name, err)
		return
	}

	s.log.Printf("spider %s: found %d magnets", name, len(magnets))
	if err := s.saveMagnets(ctx, cfg, magnets); err != nil {
		s.setRunning(name, true, fmt.Sprintf("save error: %v", err))
		return
	}
}

func (s *SpiderService) fetch(ctx context.Context, cfg *config.SpiderConfig) (string, error) {
	req, err := http.NewRequestWithContext(ctx, cfg.Method, cfg.URL, nil)
	if err != nil {
		return "", err
	}
	if cfg.UserAgent != "" {
		req.Header.Set("User-Agent", cfg.UserAgent)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	if cfg.Proxy != "" {
		proxyURL, err := url.Parse(cfg.Proxy)
		if err != nil {
			return "", fmt.Errorf("invalid proxy: %w", err)
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *SpiderService) saveMagnets(ctx context.Context, cfg *config.SpiderConfig, magnets []*model.AnimeMagnet) error {
	today := time.Now().Format("2006-01-02")
	key := fmt.Sprintf("anime:link:%s", today)

	for _, m := range magnets {
		exists, err := s.redis.HExists(ctx, key, m.Name)
		if err != nil {
			return err
		}
		if exists {
			continue
		}
		if err := s.redis.HSet(ctx, key, m.Name, m.MagnetLink); err != nil {
			return err
		}
	}

	ttl, err := s.redis.TTL(ctx, key)
	if err != nil {
		return err
	}
	if ttl < 0 {
		expire := time.Duration(s.magnetTimeout*24) * time.Hour
		if err := s.redis.Expire(ctx, key, expire); err != nil {
			return err
		}
	}

	if s.dmSvc != nil {
		go func() {
			bgCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			tasks, err := s.dmSvc.ScanAndDownload(bgCtx)
			if err != nil {
				s.log.Printf("spider %s: scan and download error: %v", cfg.Name, err)
				return
			}
			s.log.Printf("spider %s: scan matched %d tasks", cfg.Name, len(tasks))
		}()
	}

	return nil
}

func (s *SpiderService) findConfig(name string) *config.SpiderConfig {
	for i := range s.cfg {
		if s.cfg[i].Name == name {
			return &s.cfg[i]
		}
	}
	return nil
}

func (s *SpiderService) setRunning(name string, running bool, errMsg string) {
	s.statusMu.Lock()
	defer s.statusMu.Unlock()
	if st, ok := s.status[name]; ok {
		st.Running = running
		if running {
			st.LastRunAt = time.Now()
		}
		st.LastError = errMsg
	}
}
