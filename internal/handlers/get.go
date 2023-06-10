package handlers

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ozon/domain"
	"ozon/pb"
	"ozon/pkg/logger"
)

func (h *Handlers) Get(ctx context.Context, req *pb.GetUrlRequest) (*pb.GetUrlResponse, error) {
	log := logger.GetLogger()
	shortUrl := req.Url
	log.Info().Msg("shorturl parsed " + shortUrl)
	if len(shortUrl) != 10 || shortUrl == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}
	originalUrl, err := h.service.Get(ctx, shortUrl)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, domain.ErrNotFound.Error())
		} else {
			log.Warn().Err(err).Msg(err.Error())
			return nil, err
		}
	}
	return &pb.GetUrlResponse{OriginalUrl: originalUrl}, nil
}
