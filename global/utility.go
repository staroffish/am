package global

import (
	"time"
)

// FormatTime - 格式化时间
func FormatTime(t time.Time) string{
	return t.Format("2006-1-2 15:04:05")
}