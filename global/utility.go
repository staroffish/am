package global

import (
	"encoding/base64"
	"gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
)

// FormatTime - 格式化时间
func FormatTime(t time.Time) string{
	return t.Format("2006-1-2 15:04:05")
}

// ObjectIdToString - 将bson.ObjectId转换成16进制字符串
func ObjectIdToString(id bson.ObjectId) string{
	return id.Hex()
}

// EncodeBase64 - base64编码
func EncodeBase64(b []byte) string{
	return base64.StdEncoding.EncodeToString(b)
}

const (
	kilo = 1024.0
	mega = 1024.0 * 1024.0
	giga = 1024.0 * 1024.0 * 1024.0
)

// SizeToString 将大小转换成对应的KB MB GB字符串
func SizeToString(size int64) string{
	switch {
	case size < kilo:
		return fmt.Sprintf("%dB", size);
	case size < mega:
		return fmt.Sprintf("%.1fKB", float64(size) / kilo);
	case size < giga:
		return fmt.Sprintf("%.1fMB", float64(size) / mega);
	default:
		return fmt.Sprintf("%.1fGB", float64(size) / giga);
	}
}