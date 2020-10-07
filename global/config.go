package global

import (
	"encoding/json"
	"fmt"
	"log/syslog"
	"os"
	"strings"
)

var Cfg *Config

const (
	LOG_TYPE_SYSLOG = "SYSLOG"
	LOG_TYPE_FILE   = "FILE"
)

// Config 配置文件结构体定义
type Config struct {
	// mongodb的IP
	DBIP string `json:"DBIP"`
	// mongodb的端口
	DBPort string `json:"DBPort"`
	// mongodb的DB名
	DBName string `json:"DBName"`
	// mongodb的用户名
	DBUser string `json:"DBUser"`
	// mongodb的密码
	DBPasswd string `json:"DBPasswd"`
	// 本程序绑定的地址
	BindAddr string `json:"BindAddr"`
	// 是否开启自动下载
	AdOn bool `json:"AdOn"`
	// 自动下载轮循间隔
	AdInter int `json:"AdInter"`
	// BT远程下载的地址
	BTWebUrl string `json:"BTWebUrl"`
	// BT远程下载的密码
	BTWebPass string `json:"BTWebPass"`
	// 是否开启DEBUGLOG
	DebugOn bool `json:"DebugOn"`
	// 连接远程WEB服务器的超时时间
	WebTimeout int `json:"WebTimeout"`
	// 主页显示的动漫的数量
	AnimeCntInMain int `json:"AnimeCntInMain"`
	// redis ip
	RedisAddr string `json:"RedisAddr"`
	// redis password
	RedisPasswd string `json:"RedisPasswd"`
	// 默认自动下载的页面
	DefaultAdUrl string `json:"DefaultAdUrl"`
	// 默认抓取磁连的正则
	DefaultAdMagExp string `json:"DefaultAdMagExp"`
	// 文件服务器的目录(如果配置错误,网页中的动画文件链接可能无法打开)
	SaveDir string `json:"SaveDir"`
	// 动漫保存默认路径前缀
	AnimeDefaultDirPre string `json:"AnimeDefaultDirPre"`
	// 日志类型(SYSLOG,FILE)
	LogType string `json:"LogType"`
	// 日志路径 当制定SYSLOG时无效
	LogFile string `json:"LogFile"`
}

const (
	configLen = 8191
)

// NewConfig 读入配置文件
func NewConfig(filePath string) error {
	Cfg = &Config{}

	fileCtx := make([]byte, 8192)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// 只读取8192个字节
	n, err := file.Read(fileCtx)
	if err != nil && n > configLen {
		return err
	}
	fileCtx = fileCtx[:n]

	err = json.Unmarshal(fileCtx, Cfg)
	if err != nil {
		return err
	}
	return Cfg.check()
}

func (c *Config) check() error {
	if c.DBIP == "" || c.DBPort == "" ||
		c.BindAddr == "" ||
		c.BTWebUrl == "" ||
		c.DBName == "" || c.DBUser == "" ||
		c.DBPasswd == "" || c.SaveDir == "" {
		return fmt.Errorf("Config file invalid")
	}
	if c.WebTimeout == 0 {
		c.WebTimeout = 10
	}
	if c.AnimeCntInMain == 0 {
		c.AnimeCntInMain = 30
	}

	if !strings.HasPrefix(c.SaveDir, "/") {
		c.SaveDir += "/"
	}
	if !strings.HasPrefix(c.AnimeDefaultDirPre, "/") {
		c.AnimeDefaultDirPre += "/"
	}
	var logger Logger
	var err error
	switch c.LogType {
	case LOG_TYPE_SYSLOG:
		logger, err = syslog.NewLogger(syslog.LOG_USER|syslog.LOG_INFO, 0)
		if err != nil {
			return fmt.Errorf("New sys logger error", err)
		}
	case LOG_TYPE_FILE:
		logger, err = NewFileLogger(c.LogFile)
		if err != nil {
			return fmt.Errorf("New file logger error, %v", err)
		}
	default:
		return fmt.Errorf("invaild log type:%s. log type must be:%s,%s", c.LogType, LOG_TYPE_SYSLOG, LOG_TYPE_FILE)
	}
	NewLogger(c.DebugOn, logger)
	return nil
}
