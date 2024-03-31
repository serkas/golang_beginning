package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"rest-server-demo/internal/api"
	"rest-server-demo/internal/services/sensors"
	"rest-server-demo/internal/storage"
	"time"
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
	handlers := api.New(a.log, sensorsService)

	s := newServer(handlers, a.cfg.Port)

	go func() {
		a.log.Info("starting server", zap.Int("port", a.cfg.Port))

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
}

func newServer(apiHandlers *api.API, port int) *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/", apiHandlers.Hello).Methods(http.MethodGet)
	r.HandleFunc("/api/measurements", apiHandlers.Measurements).Methods(http.MethodPost)

	// TODO: add sensor CRUD API

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	return s
}
