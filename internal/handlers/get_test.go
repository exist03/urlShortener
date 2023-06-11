package handlers

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ozon/domain"
	mock_handlers "ozon/internal/handlers/mocks"
	"ozon/pb"
	"testing"
)

func TestHandlers_Get(t *testing.T) {
	type mockBehavior func(r *mock_handlers.MockService, shortUrl string)
	tests := []struct {
		name    string
		request *pb.GetUrlRequest
		mockBehavior
		expectedResponse *pb.GetUrlResponse
		expectedErr      error
	}{
		{
			name:    "ok",
			request: &pb.GetUrlRequest{Url: "9Txmx1CPWV"},
			mockBehavior: func(r *mock_handlers.MockService, shortUrl string) {
				r.EXPECT().Get(context.Background(), shortUrl).Return("google.com", nil)
			},
			expectedResponse: &pb.GetUrlResponse{OriginalUrl: "google.com"},
			expectedErr:      nil,
		},
		{
			name:             "Bad request",
			request:          &pb.GetUrlRequest{Url: "9Txmx1CPWVT"},
			mockBehavior:     func(r *mock_handlers.MockService, shortUrl string) {},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.InvalidArgument, "Bad request"),
		},
		{
			name:             "Bad request",
			request:          &pb.GetUrlRequest{Url: ""},
			mockBehavior:     func(r *mock_handlers.MockService, shortUrl string) {},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.InvalidArgument, "Bad request"),
		},
		{
			name:             "Bad request",
			request:          &pb.GetUrlRequest{Url: " "},
			mockBehavior:     func(r *mock_handlers.MockService, shortUrl string) {},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.InvalidArgument, "Bad request"),
		},
		{
			name:    "Not found",
			request: &pb.GetUrlRequest{Url: "9Txmj1CPWV"},
			mockBehavior: func(r *mock_handlers.MockService, shortUrl string) {
				r.EXPECT().Get(context.Background(), shortUrl).Return("", domain.ErrNotFound)
			},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.NotFound, domain.ErrNotFound.Error()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := mock_handlers.NewMockService(ctrl)
			tt.mockBehavior(s, tt.request.Url)
			handlers := New(s)
			response, err := handlers.Get(context.Background(), tt.request)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedResponse, response)
		})
	}
}
