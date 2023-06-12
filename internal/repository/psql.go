package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"ozon/config"
	"ozon/domain"
	"ozon/pkg/logger"
	"ozon/pkg/postgresql"
)

const sqlDuplicatePK = "ERROR: duplicate key value violates unique constraint \"urls_pkey\" (SQLSTATE 23505)"

type PsqlRepo struct {
	pool   *pgxpool.Pool
	logger zerolog.Logger
}

func NewPsql(ctx context.Context, config config.PsqlStorage) *PsqlRepo {
	log := logger.GetLogger()
	pool, err := postgresql.NewClient(ctx, config, 3)
	if err != nil {
		log.Fatal().Err(err).Msg("Can`t create psql client")
	}
	return &PsqlRepo{pool: pool, logger: log}
}

func (r *PsqlRepo) Create(ctx context.Context, shortURL, url string) (err error) {
	_, err = r.pool.Exec(ctx, "INSERT INTO urls VALUES ($1, $2)", shortURL, url)
	if err != nil {
		if err.Error() == sqlDuplicatePK {
			r.logger.Info().Err(err).Msg("duplicate PK")
			return nil
		}
		r.logger.Warn().Err(err).Msg("create error")
		return err
	}
	return nil
}

func (r *PsqlRepo) Get(ctx context.Context, shortUrl string) (string, error) {
	var originalUrl string
	err := r.pool.QueryRow(ctx, "SELECT long FROM urls WHERE short=$1", shortUrl).Scan(&originalUrl)
	if err == pgx.ErrNoRows {
		r.logger.Info().Err(err).Msg("404")
		return "", domain.ErrNotFound
	}
	if err != nil {
		r.logger.Error().Err(err).Msg("get error")
		return "", err
	}
	return originalUrl, nil
}
