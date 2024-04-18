package app

import (
	"context"
	"errors"
	"layered-service/internal/consumer"
	"layered-service/internal/services/measuring"
	"net/http"
	"time"

	"go.uber.org/zap"

	"layered-service/internal/api"
	"layered-service/internal/storage"
)

type App struct {
	log *zap.Logger
	cfg Config
}

func New(logger *zap.Logger, cfg Config) *App {
	return &App{
		log: logger,
		cfg: cfg,
	}
}

// Run is the main execution function of entire application
func (a *App) Run(ctx context.Context) {
	// creating components, injecting dependencies

	// starting from the least dependent (persistence)
	store := storage.NewMemStorage()
	// moving to services
	measuringService := measuring.NewService(a.log, store)

	// finishing with presentation layer and connectors (depending on underlying layers)
	router := api.New(a.log, measuringService, store).CreateRouter()
	s := newServer(router, a.cfg.Address)

	measurementConsumer := consumer.New(a.log, measuringService)

	// starting application components
	go func() {
		err := measurementConsumer.Run(ctx)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Error("running consumer", zap.Error(err))
		}
	}()

	go func() {
		a.log.Info("starting server", zap.String("addr", a.cfg.Address))
		err := s.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Error("running server", zap.Error(err))
		}
	}()

	<-ctx.Done()
	// Handle graceful shutdown
	shutdownTimeout := 10 * time.Second
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := s.Shutdown(shutdownCtx); err != nil {
		a.log.Fatal("stopping server", zap.Error(err))
	}
	a.log.Info("stopped gracefully")
}

func newServer(handler http.Handler, address string) *http.Server {
	s := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	return s
}
