package handlers

import (
	"golang.org/x/net/context"
	"ozon/pb"
)

type Service interface {
	Encrypt(url string) string
	Create(ctx context.Context, shortURL, url string) error
	Get(ctx context.Context, shortUrl string) (string, error)
}

type Handlers struct {
	service Service
	pb.UnimplementedGatewayServer
}

func New(service Service) *Handlers {
	return &Handlers{service: service}
}
