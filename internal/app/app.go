package app

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"ozon/config"
	"ozon/internal/handlers"
	"ozon/internal/repository"
	"ozon/internal/service"
	"ozon/pb"
	"ozon/pkg/logger"
)

const (
	grpcPort = ":8090"
	httpPort = ":8080"
)

type Repository interface {
	Get(ctx context.Context, shortUrl string) (string, error)
	Create(ctx context.Context, shortURL, url string) error
}
type App struct {
	handlers   *handlers.Handlers
	service    *service.Service
	repository Repository
}

func Run(ctx context.Context, cfg *config.Config) {
	storage := os.Getenv("STORAGE_TYPE") //TODO: change to config
	a := &App{}
	log := logger.GetLogger()
	switch storage {
	case "psql":
		a.repository = repository.NewPsql(ctx, cfg.PsqlStorage)
		log.Info().Msg("Psql storage")
	case "redis":
		a.repository = repository.NewRedis(ctx, cfg.RedisStorage)
		log.Info().Msg("Redis storage")
	case "inMemo":
		a.repository = repository.NewInMemoryRepo()
		log.Info().Msg("in-memory storage")
	default:
		log.Fatal().Msg("No database has chosen")
	}
	a.service = service.New(a.repository)

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}
	a.handlers = handlers.New(a.service)

	s := grpc.NewServer()
	pb.RegisterGatewayServer(s, a.handlers)
	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatal().Msg("failed to serve: " + err.Error())
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		grpcPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal().Msg("Failed to dial server: " + err.Error())
	}
	defer conn.Close()

	gwmux := runtime.NewServeMux()
	err = pb.RegisterGatewayHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatal().Msg("Failed to register gateway:" + err.Error())
	}

	gwServer := &http.Server{
		Addr:    httpPort,
		Handler: gwmux,
	}

	log.Info().Msg("Serving gRPC-Gateway on port " + httpPort)
	if err = gwServer.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Warn().Msg("Server closed: " + err.Error())
			os.Exit(0)
		}
		log.Fatal().Msg("Failed to listen and serve: " + err.Error())
	}
}
