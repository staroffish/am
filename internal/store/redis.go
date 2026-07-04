package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/staroffish/am/internal/config"
)

type RedisClient struct {
	cli *redis.Client
}

func NewRedisClient(cfg config.RedisConfig) (*RedisClient, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis connect: %w", err)
	}

	return &RedisClient{cli: cli}, nil
}

func (r *RedisClient) Client() *redis.Client {
	return r.cli
}

func (r *RedisClient) Close() error {
	return r.cli.Close()
}

func (r *RedisClient) SaveCookies(ctx context.Context, key string, cookies map[string]string) error {
	pipe := r.cli.Pipeline()
	for k, v := range cookies {
		pipe.HSet(ctx, key, k, v)
	}
	pipe.Expire(ctx, key, 5*time.Minute)
	_, err := pipe.Exec(ctx)
	return err
}

func (r *RedisClient) LoadCookies(ctx context.Context, key string) (map[string]string, error) {
	vals, err := r.cli.HGetAll(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return vals, nil
}

func (r *RedisClient) ClearCookies(ctx context.Context, key string) error {
	return r.cli.Del(ctx, key).Err()
}

func (r *RedisClient) HSet(ctx context.Context, key, field, value string) error {
	return r.cli.HSet(ctx, key, field, value).Err()
}

func (r *RedisClient) HExists(ctx context.Context, key, field string) (bool, error) {
	return r.cli.HExists(ctx, key, field).Result()
}

func (r *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.cli.HGetAll(ctx, key).Result()
}

func (r *RedisClient) Expire(ctx context.Context, key string, d time.Duration) error {
	return r.cli.Expire(ctx, key, d).Err()
}

func (r *RedisClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	return r.cli.TTL(ctx, key).Result()
}
