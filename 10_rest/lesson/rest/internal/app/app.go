package app

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"

	"rest-server-demo/internal/api"
	"rest-server-demo/internal/services/sensors"
	"rest-server-demo/internal/storage"
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

func (a *App) Run(ctx context.Context) {
	store := storage.NewMemStorage()

	sensorsService := sensors.New(a.log, store)

	router := api.New(a.log, sensorsService).CreateRouter()

	s := newServer(router, a.cfg.Address)

	go func() {
		a.log.Info("starting server", zap.String("addr", a.cfg.Address))

		err := s.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Fatal("running server", zap.Error(err))
		}
	}()

	// Handle graceful shutdown
	shutdownTimeout := 10 * time.Second
	<-ctx.Done()
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
