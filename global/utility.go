package global

import (
	"encoding/base64"
	"gopkg.in/mgo.v2/bson"
	"time"
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