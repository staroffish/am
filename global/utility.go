package global

import (
	"encoding/base64"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// FormatTime - 格式化时间
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// ObjectIdToString - 将bson.ObjectId转换成16进制字符串
func ObjectIdToString(id bson.ObjectId) string {
	return id.Hex()
}

// EncodeBase64 - base64编码
func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

const (
	kilo = 1024.0
	mega = 1024.0 * 1024.0
	giga = 1024.0 * 1024.0 * 1024.0
)

// SizeToString 将大小转换成对应的KB MB GB字符串
func SizeToString(size int64) string {
	switch {
	case size < kilo:
		return fmt.Sprintf("%dB", size)
	case size < mega:
		return fmt.Sprintf("%.1fKB", float64(size)/kilo)
	case size < giga:
		return fmt.Sprintf("%.1fMB", float64(size)/mega)
	default:
		return fmt.Sprintf("%.1fGB", float64(size)/giga)
	}
}

// 计算现在属于哪一季度
func GetNowSeason() (int, int) {
	now := time.Now()
	nowMonth := int(now.Month())
	var season int

	// 计算添加的是哪一季度的动画
	if dateDiff := nowMonth - 1; dateDiff < 2 {
		season = 1
	} else if dateDiff = nowMonth - 4; dateDiff < 2 {
		season = 4
	} else if dateDiff = nowMonth - 7; dateDiff < 2 {
		season = 7
	} else if dateDiff = nowMonth - 10; dateDiff < 2 {
		season = 10
	}

	return now.Year(), season
}
