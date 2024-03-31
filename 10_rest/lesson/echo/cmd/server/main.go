package main

import (
	"context"
	"echo-server-demo/internal/app"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	// Config should be populated from flags, config file or (the best option) environment variables
	cfg := app.Config{
		Port: 8001,
	}

	a := app.New(logger, cfg)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	a.Run(ctx)
}
