package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/syalsr/notification/internal/app/service"
	"github.com/syalsr/notification/internal/config"
	"github.com/syalsr/notification/internal/usecase"
	"github.com/syalsr/notification/internal/usecase/emailer"
	api "github.com/syalsr/notification/pkg/v1"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type canceler func()

// Run - run grpc server
func Run(ctx context.Context, cfg *config.App) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	listener, err := net.Listen("tcp", cfg.GrpcAddr)
	if err != nil {
		log.Err(err).Msgf("cant connected to %s", cfg.GrpcAddr)
	}

	log.Info().Msg("Create new gRPC server")
	server := grpc.NewServer()

	log.Info().Msg("Register gRPC server")

	canceler := make([]canceler, 0)
	email := emailer.NewEmailer(cfg)
	notif := usecase.NewNotificator(email)
	serviceNotificator := service.NewNotificator(cfg, notif)

	api.RegisterNotificationServiceServer(server, serviceNotificator)
	go func() {
		log.Info().Msgf("Start gRPC server on %s", cfg.GrpcAddr)
		if err = server.Serve(listener); err != nil {
			log.Fatal().Msgf("cant start gRPC server: %w", err)
		}
	}()

	gracefulShutDown(server, cancel, canceler)

	return nil
}

func gracefulShutDown(s *grpc.Server, cancel context.CancelFunc, canceler []canceler) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	c := <-ch
	log.Info().Msgf("Called graceful shutdown: %v", c)

	for _, item := range canceler {
		item()
	}
	s.GracefulStop()
	cancel()
}
