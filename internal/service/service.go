package service

import (
	"crypto/sha1"
	"golang.org/x/net/context"
	"ozon/domain"
	"regexp"
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
func (s *Service) Hash(url string) string {
	var result string
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	h := sha1.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)[:10]
	for _, b := range bs {
		temp := int16(b)
		index := temp % 63
		result = result + string(alphabet[index])
	}
	return result
}
func (s *Service) Create(ctx context.Context, url string) (string, error) {
	re := regexp.MustCompile(".+\\..+")
	if !re.MatchString(url) {
		return "", domain.ErrInvalidArgument
	}
	shortUrl := s.Hash(url)
	return shortUrl, s.repo.Create(ctx, shortUrl, url)
}
func (s *Service) Get(ctx context.Context, shortUrl string) (string, error) {
	if len(shortUrl) != 10 || shortUrl == "" {
		return "", domain.ErrInvalidArgument
	}
	return s.repo.Get(ctx, shortUrl)
}
