package global

import (
	"encoding/json"
	"fmt"
	"os"
)

var Cfg *Config

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
	BTWebIP string `json:"BTWebIP"`
	// BT远程下载的端口
	BTWebPort string `json:"BTWebPort"`
	// 是否开启DEBUGLOG
	DebugOn bool `json:"DebugOn"`
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
		c.BTWebIP == "" || c.BTWebPort == "" ||
		c.DBName == "" || c.DBUser == "" ||
		c.DBPasswd == "" {
		return fmt.Errorf("Config file invalid")
	}
	return nil
}
