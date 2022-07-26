package util

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-redis/redis/v8"
	"github.com/staroffish/am/common/config"
	etcd "go.etcd.io/etcd/client/v3"
)

func NewTestLogger() log.Logger {
	return log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.Caller(4),
	)
}

func SendHttpRequest(ctx context.Context, method string, url, bodyStr string, cookies []*http.Cookie, headers http.Header) (*http.Response, error) {
	r, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(bodyStr))
	if err != nil {
		return nil, fmt.Errorf("DelugeDownloader.sendReq.NewRequest:%v", err)
	}

	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}

	for headerName, headerValues := range headers {
		for _, headerValue := range headerValues {
			r.Header.Set(headerName, headerValue)
		}
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("SendHttpRequest:err=%v, jsonStr=%s, url=%s", err, bodyStr, url)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SendHttpRequest:status %d is not OK", resp.StatusCode)
	}

	return resp, nil
}

func NewReidsClient(c *config.ReidsConfig) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Username:     c.User,
		Password:     c.Password,
		ReadTimeout:  time.Duration(c.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.WriteTimeout) * time.Second,
		DB:           int(c.Db),
	})
	_, err := redisClient.Get(context.Background(), "").Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	return redisClient, nil
}

func NewLogger(client *etcd.Client, componentName config.ComponentName, version config.Version) (log.Logger, error) {
	resp, err := client.Get(context.Background(), fmt.Sprintf("/component/%s/log_path", string(componentName)))
	logPath := ""
	if err != nil {
		return nil, fmt.Errorf("util:NewLogger error:%v. component name=%s", err, componentName)
	}
	for _, kv := range resp.Kvs {
		logPath = string(kv.Value)
	}

	id, _ := os.Hostname()

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.Caller(4),
		"service.id", id,
		"service.name", componentName,
		"service.version", version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	if logPath != "" {
		fileBase := filepath.Dir(logPath)
		if fileBase != "." {
			if err := os.MkdirAll(fileBase, 0755); err != nil {
				return nil, fmt.Errorf("util:NewLogger MkdirAll error:%v. component name=%s, file name=%s", err, componentName, logPath)
			}
		}

		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("util:NewLogger OpenFile error:%v. component name=%s, file name=%s", err, componentName, logPath)
		}
		logger = log.With(log.NewStdLogger(logFile),
			"ts", log.DefaultTimestamp,
			"caller", log.Caller(4),
			"service.id", id,
			"service.name", componentName,
			"service.version", version,
			"trace_id", tracing.TraceID(),
			"span_id", tracing.SpanID(),
		)
	}
	return logger, nil
}
