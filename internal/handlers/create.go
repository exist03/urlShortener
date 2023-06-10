package handlers

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ozon/pb"
	"ozon/pkg/logger"
	"regexp"
)

func (h *Handlers) Create(ctx context.Context, req *pb.CreateUrlRequest) (*pb.CreateUrlResponse, error) {
	log := logger.GetLogger()
	originalUrl := req.Url
	log.Info().Msg(originalUrl)
	re := regexp.MustCompile(".+\\..+")
	if !re.MatchString(originalUrl) {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}
	shortUrl := h.service.Encrypt(originalUrl)
	if err := h.service.Create(ctx, shortUrl[5:], originalUrl); err != nil {
		log.Warn().Msg("problems with insert in db")
		return nil, err
	}
	return &pb.CreateUrlResponse{ShortUrl: shortUrl}, nil
}
