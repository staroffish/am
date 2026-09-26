package util

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "[am] ", log.LstdFlags|log.Lshortfile)
}

func SendHTTPRequest(ctx context.Context, method, url, bodyStr string, cookies []*http.Cookie, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(bodyStr))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	for name, values := range headers {
		for _, val := range values {
			req.Header.Set(name, val)
		}
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}
	return resp, nil
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormatSize(size int64) string {
	const (
		kb = 1024
		mb = 1024 * kb
		gb = 1024 * mb
	)
	switch {
	case size < kb:
		return fmt.Sprintf("%dB", size)
	case size < mb:
		return fmt.Sprintf("%.1fKB", float64(size)/kb)
	case size < gb:
		return fmt.Sprintf("%.1fMB", float64(size)/mb)
	default:
		return fmt.Sprintf("%.1fGB", float64(size)/gb)
	}
}

// NearestSeason 返回离 t 最近的季度月份(1/4/7/10 月)。
// 例如 9 月取 10 月、8 月取 7 月、12 月取次年 1 月。
func NearestSeason(t time.Time) (int, int) {
	y, m := t.Year(), int(t.Month())

	switch {
	case m == 12: // 离次年 1 月最近
		return y + 1, 1
	case m >= 9: // 9,10,11 -> 10 月
		return y, 10
	case m >= 6: // 6,7,8 -> 7 月
		return y, 7
	case m >= 3: // 3,4,5 -> 4 月
		return y, 4
	default: // 1,2 -> 1 月
		return y, 1
	}
}

func GetNowSeason() (int, int) {
	return NearestSeason(time.Now())
}

func GetNextSeason() (int, int) {
	y, s := GetNowSeason()
	switch s {
	case 1:
		return y, 4
	case 4:
		return y, 7
	case 7:
		return y, 10
	default:
		return y + 1, 1
	}
}
