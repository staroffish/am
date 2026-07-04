package cron

import (
	"context"
	"log"
	"time"

	"github.com/staroffish/am/internal/config"
	"github.com/staroffish/am/internal/service"
)

func Setup(cfg *config.Config, spiderSvc *service.SpiderService, autoDownload bool, cronLogger *log.Logger) *time.Ticker {
	if !autoDownload || cfg.AutoDownload.Interval <= 0 || len(cfg.Spiders) == 0 {
		return nil
	}

	interval := time.Duration(cfg.AutoDownload.Interval) * time.Second
	ticker := time.NewTicker(interval)

	go func() {
		idx := 0
		for range ticker.C {
			name := cfg.Spiders[idx].Name
			spiderSvc.Crawl(context.Background(), name)

			idx++
			if idx >= len(cfg.Spiders) {
				idx = 0
			}
		}
	}()

	return ticker
}
