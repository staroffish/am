package config

import (
	"context"
	"fmt"
	"strconv"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServerConfig struct {
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout int64  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

type HttpServerConfig ServerConfig
type GrpcServerConfig ServerConfig
type ComponentName string
type ComponentType string
type Version string

type ReidsConfig struct {
	Addr         string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	User         string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Password     string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Db           int64  `protobuf:"varint,4,opt,name=db,proto3" json:"db,omitempty"`
	ReadTimeout  int64  `protobuf:"bytes,6,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout int64  `protobuf:"bytes,7,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
}

func NewGrpcServerConfig(client *clientv3.Client, componentName ComponentName) (*GrpcServerConfig, error) {
	prefix := fmt.Sprintf("/component/%s/grpc", componentName)

	grpcConfig, err := newServerConfig(client, prefix)
	if err != nil {
		return nil, err
	}

	return (*GrpcServerConfig)(grpcConfig), nil
}

func NewHTTPServerConfig(client *clientv3.Client, componentName ComponentName) (*HttpServerConfig, error) {
	prefix := fmt.Sprintf("/component/%s/http", componentName)

	httpConfig, err := newServerConfig(client, prefix)
	if err != nil {
		return nil, err
	}

	return (*HttpServerConfig)(httpConfig), nil
}

func newServerConfig(client *clientv3.Client, prefix string) (*ServerConfig, error) {
	serverConfig := &ServerConfig{}
	resp, err := client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("get %s error:%v", prefix, err)
	}
	for _, kv := range resp.Kvs {
		key := string(kv.Key)
		switch key {
		case fmt.Sprintf("%s/addr", prefix):
			serverConfig.Addr = string(kv.Value)
		case fmt.Sprintf("%s/timeout", prefix):
			timeout, err := strconv.ParseInt(string(kv.Value), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("convert %s to int error:%v", key, err)
			}
			serverConfig.Timeout = timeout
			if serverConfig.Timeout == 0 {
				serverConfig.Timeout = 5
			}
		case fmt.Sprintf("%s/network", prefix):
			serverConfig.Network = string(kv.Value)
		}
	}

	if serverConfig.Addr == "" {
		return nil, fmt.Errorf("http or grpc server addr is empty")
	}
	return serverConfig, nil
}

func NewRedisConfig(client *clientv3.Client) (*ReidsConfig, error) {
	ctx := context.Background()
	prefix := "/database/redis"

	redisConfig := &ReidsConfig{}
	redisConfig.ReadTimeout = 5
	redisConfig.WriteTimeout = 5

	resp, err := client.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("get %s error:%v", prefix, err)
	}
	for _, kv := range resp.Kvs {
		key := string(kv.Key)
		switch key {
		case fmt.Sprintf("%s/addr", prefix):
			redisConfig.Addr = string(kv.Value)
		case fmt.Sprintf("%s/password", prefix):
			redisConfig.Password = string(kv.Value)
		case fmt.Sprintf("%s/user", prefix):
			redisConfig.User = string(kv.Value)
		case fmt.Sprintf("%s/read_timeout", prefix):
			timeout, err := strconv.ParseInt(string(kv.Value), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("convert %s to int error:%v", key, err)
			}
			redisConfig.ReadTimeout = timeout
		case fmt.Sprintf("%s/write_timeout", prefix):
			timeout, err := strconv.ParseInt(string(kv.Value), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("convert %s to int error:%v", key, err)
			}
			redisConfig.WriteTimeout = timeout
		case fmt.Sprintf("%s/db", prefix):
			db, err := strconv.ParseInt(string(kv.Value), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("convert %s to int error:%v", key, err)
			}
			redisConfig.Db = db
		}
	}

	return redisConfig, nil
}
