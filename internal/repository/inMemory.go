package repository

import (
	"context"
	"github.com/rs/zerolog"
	"ozon/domain"
	"ozon/pkg/logger"
	"sync"
)

type InMemoryRepo struct {
	mem    map[string]string
	mu     *sync.Mutex
	logger zerolog.Logger
}

func NewInMemoryRepo() *InMemoryRepo {
	log := logger.GetLogger()
	inMemoryStorage := make(map[string]string)
	return &InMemoryRepo{mem: inMemoryStorage, mu: &sync.Mutex{}, logger: log}
}

func (r *InMemoryRepo) Create(ctx context.Context, shortUrl, url string) (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.mem[shortUrl]; !ok {
		r.mem[shortUrl] = url
	}
	return nil
}

func (r *InMemoryRepo) Get(ctx context.Context, shortUrl string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	originalUrl, ok := r.mem[shortUrl]
	if !ok {
		return "", domain.ErrNotFound
	}
	return originalUrl, nil
}
