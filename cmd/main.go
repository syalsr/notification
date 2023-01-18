package main

import (
	"context"

	"github.com/syalsr/notification/internal/app"
	"github.com/syalsr/notification/internal/config"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := &config.App{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal().Msgf("failed to retrieve env variables, %v", err)
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		log.Fatal().Msgf("error running grpc server ", err)
	}
}
