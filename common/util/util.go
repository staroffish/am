package util

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/staroffish/am/common/config"
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
