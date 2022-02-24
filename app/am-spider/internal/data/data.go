package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/staroffish/am/app/am-spider/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAmSpiderRepo)

// Data .
type Data struct {
	*redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Username:     c.Redis.User,
		Password:     c.Redis.Password,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		DB:           int(c.Redis.Db),
	})
	_, err := redisClient.Get(context.Background(), "").Result()
	if err != nil && err != redis.Nil {
		log.NewHelper(logger).Errorf("redisClient.Get error %v", err)
		panic(fmt.Sprintf("redisClient.Get error %v", err))
	}
	return &Data{
		Client: redisClient,
	}, cleanup, nil
}
