package repository

import (
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"golang.org/x/net/context"
	"ozon/config"
	"ozon/domain"
	"ozon/pkg/logger"
	"ozon/pkg/redisdb"
)

type RedisRepo struct {
	redis  *redis.Client
	logger zerolog.Logger
}

func NewRedis(ctx context.Context, config config.RedisStorage) *RedisRepo {
	client, err := redisdb.NewClient(ctx, config, 3)
	if err != nil {
		return nil
	}
	log := logger.GetLogger()
	return &RedisRepo{redis: client, logger: log}
}

func (r *RedisRepo) Create(ctx context.Context, shortURL, originalUrl string) error {
	err := r.redis.Set(ctx, shortURL, originalUrl, 0).Err()
	if err != nil {
		r.logger.Error().Err(err).Msg("create")
		return err
	}
	return nil
}
func (r *RedisRepo) Get(ctx context.Context, shortUrl string) (string, error) {
	originalUrl, err := r.redis.Get(ctx, shortUrl).Result()
	if err == redis.Nil {
		r.logger.Info().Err(err).Msg("404")
		return "", domain.ErrNotFound
	} else if err != nil {
		r.logger.Error().Err(err).Msg("get error")
		return "", err
	}
	return originalUrl, nil
}
