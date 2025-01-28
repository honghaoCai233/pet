package clients

import (
	"context"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"

	"go.uber.org/zap"

	"pet/configs"
	libCron "pet/pkg/cron"
	"pet/pkg/logger"
)

var ProviderSet = wire.NewSet(
	NewLogger,
	NewCron,
	NewRedisClient,
)

func NewRedisClient(conf *configs.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}

func NewLogger(conf *configs.Config) *zap.SugaredLogger {
	return logger.NewLogger(&conf.Log)
}

func NewCron() (*cron.Cron, func(), error) {
	c := libCron.New()
	c.Start()
	return c, func() {
		c.Stop()
	}, nil
}
