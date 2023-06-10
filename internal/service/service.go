package service

import (
	"golang.org/x/net/context"
	"ozon/pkg/utils"
)

type Repository interface {
	Get(ctx context.Context, shortUrl string) (string, error)
	Create(ctx context.Context, shortURL, url string) error
}
type Service struct {
	repo Repository
}

func New(repository Repository) *Service {
	return &Service{repo: repository}
}
func (s *Service) Encrypt(url string) string {
	return "xy.z/" + utils.Encrypt(url)
}
func (s *Service) Create(ctx context.Context, shortURL, url string) error {
	return s.repo.Create(ctx, shortURL, url)
}
func (s *Service) Get(ctx context.Context, shortUrl string) (string, error) {
	return s.repo.Get(ctx, shortUrl)
}
