package handlers

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"net/http"
	"ozon/domain"
	"ozon/pb"
	"ozon/pkg/logger"
)

type Service interface {
	Hash(url string) string
	Create(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, shortUrl string) (string, error)
}

type Handlers struct {
	service Service
	pb.UnimplementedGatewayServer
}

func New(service Service) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) Create(ctx context.Context, req *pb.CreateUrlRequest) (*pb.CreateUrlResponse, error) {
	log := logger.GetLogger()
	originalUrl := req.Url
	shortUrl, err := h.service.Create(ctx, originalUrl)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidArgument) {
			return nil, status.Error(http.StatusBadRequest, "Bad request")
		}
		log.Warn().Err(err).Msg("Inserting problems")
		return nil, err
	}
	return &pb.CreateUrlResponse{ShortUrl: shortUrl}, nil
}
func (h *Handlers) Get(ctx context.Context, req *pb.GetUrlRequest) (*pb.GetUrlResponse, error) {
	log := logger.GetLogger()
	shortUrl := req.Url
	originalUrl, err := h.service.Get(ctx, shortUrl)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(http.StatusNotFound, "Not found")
		} else if errors.Is(err, domain.ErrInvalidArgument) {
			return nil, status.Error(http.StatusBadRequest, "Bad request")
		} else {
			log.Warn().Err(err).Msg("Get problems")
			return nil, err
		}
	}
	return &pb.GetUrlResponse{OriginalUrl: originalUrl}, nil
}
