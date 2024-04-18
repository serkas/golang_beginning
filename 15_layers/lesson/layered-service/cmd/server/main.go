package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"go.uber.org/zap"

	"layered-service/internal/app"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	a := app.New(logger, initConfig())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	a.Run(ctx)
}

func initConfig() app.Config {
	// Config should be populated from flags, config file or (the best option) environment variables
	cfg := app.Config{
		Address: os.Getenv("SERVER_ADDRESS"),
	}

	if cfg.Address == "" {
		cfg.Address = ":8001"
	}

	return cfg
}
