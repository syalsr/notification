package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/syalsr/notification/internal/app/service"
	"github.com/syalsr/notification/internal/config"
	"github.com/syalsr/notification/internal/kafka"
	"github.com/syalsr/notification/internal/usecase"
	"github.com/syalsr/notification/internal/usecase/emailer"
	api "github.com/syalsr/notification/pkg/v1"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type canceler func() error

// Run - run grpc server
func Run(ctx context.Context, cfg *config.App) error {
	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	listener, err := net.Listen("tcp", cfg.GrpcAddr)
	if err != nil {
		log.Err(err).Msgf("cant connected to %s", cfg.GrpcAddr)
	}

	log.Info().Msg("Create new gRPC server")
	server := grpc.NewServer()

	log.Info().Msg("Register gRPC server")

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

	client, err := kafka.NewConsumer(cfg)
	if err != nil {
		log.Err(err).Msgf("cant create new consumer: %w", err)
	}
	log.Info().Msgf("Start kafka consumer on %s", cfg.KafkaURL)

	commonEmail := make(chan string)
	personEmail := make(chan string)
	go client.Run(ctxCancel, commonEmail, personEmail)
	go notif.Run(ctxCancel, commonEmail, personEmail)

	canceler := []canceler{client.Close}
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
		err := item()
		if err != nil {
			log.Err(err).Msgf("cant close %w", err)
		}
	}
	s.GracefulStop()
	cancel()
}
