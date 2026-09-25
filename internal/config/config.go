package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server       ServerConfig       `yaml:"server"`
	MongoDB      MongoDBConfig      `yaml:"mongodb"`
	Redis        RedisConfig        `yaml:"redis"`
	QBittorrent  QBittorrentConfig  `yaml:"qbittorrent"`
	Spiders      []SpiderConfig     `yaml:"spiders"`
	AutoDownload AutoDownloadConfig `yaml:"auto_download"`
	Anime        AnimeConfig        `yaml:"anime"`
	AI           AIConfig           `yaml:"ai"`
	Log          LogConfig          `yaml:"log"`
}

type LogConfig struct {
	File string `yaml:"file"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (s ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type MongoDBConfig struct {
	URI      string `yaml:"uri"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisConfig struct {
	Addr         string `yaml:"addr"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type QBittorrentConfig struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type SpiderConfig struct {
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	URL       string `yaml:"url"`
	Method    string `yaml:"method"`
	Proxy     string `yaml:"proxy"`
	UserAgent string `yaml:"user_agent"`
}

type AutoDownloadConfig struct {
	Enabled        bool `yaml:"enabled"`
	ScanAfterCrawl bool `yaml:"scan_after_crawl"`
	MagnetTimeout  int  `yaml:"magnet_timeout"`
	Interval        int  `yaml:"interval"`
}

type AnimeConfig struct {
	DefaultStoreDirPrefix string `yaml:"default_store_dir_prefix"`
	MainPageCount         int    `yaml:"main_page_count"`
}

type AIConfig struct {
	BaseURL string `yaml:"base_url"`
	APIKey  string `yaml:"api_key"`
	Model   string `yaml:"model"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parse config file: %w", err)
	}

	cfg.setDefaults()
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) setDefaults() {
	if c.Server.Host == "" {
		c.Server.Host = "0.0.0.0"
	}
	if c.Server.Port == 0 {
		c.Server.Port = 8080
	}
	if c.Redis.ReadTimeout == 0 {
		c.Redis.ReadTimeout = 5
	}
	if c.Redis.WriteTimeout == 0 {
		c.Redis.WriteTimeout = 5
	}
	if c.Anime.MainPageCount == 0 {
		c.Anime.MainPageCount = 50
	}
	if c.AutoDownload.MagnetTimeout == 0 {
		c.AutoDownload.MagnetTimeout = 30
	}
	if c.AutoDownload.Interval == 0 {
		c.AutoDownload.Interval = 3600
	}
	for i := range c.Spiders {
		if c.Spiders[i].Method == "" {
			c.Spiders[i].Method = "GET"
		}
	}
}

func (c *Config) validate() error {
	if c.MongoDB.URI == "" {
		return fmt.Errorf("mongodb.uri is required")
	}
	if c.MongoDB.Database == "" {
		return fmt.Errorf("mongodb.database is required")
	}
	if c.Redis.Addr == "" {
		return fmt.Errorf("redis.addr is required")
	}
	if c.QBittorrent.URL == "" {
		return fmt.Errorf("qbittorrent.url is required")
	}
	if c.QBittorrent.Username == "" {
		return fmt.Errorf("qbittorrent.username is required")
	}
	return nil
}
