package redisdb

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"ozon/config"
	"ozon/pkg/logger"
	"ozon/pkg/utils"
	"time"
)

func NewClient(ctx context.Context, config config.RedisStorage, maxAttempts int) (client *redis.Client, err error) {
	log := logger.GetLogger()
	dsn := fmt.Sprintf("redis://%s:@%s:%s/%s", config.Username, config.Host, config.Port, config.Database)
	log.Info().Msg(dsn)
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}
	err = utils.DoWithTries(func() error {
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		client = redis.NewClient(opt)
		if err != nil {
			log.Err(err)
			return err
		}
		return nil
	}, maxAttempts, 5*time.Second)
	//one more check
	if err = client.Ping(ctx).Err(); err != nil {
		log.Fatal().Err(err).Msg("Redis is unable to connect.")
	}
	log.Info().Msg("success connection redis")
	return client, nil
}
