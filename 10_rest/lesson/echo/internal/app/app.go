package app

import (
	"context"
	"echo-server-demo/internal/api"
	"echo-server-demo/internal/services/sensors"
	"echo-server-demo/internal/storage"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
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

	e := newEchoServer(handlers)
	addr := fmt.Sprintf(":%d", a.cfg.Port)

	go func() {
		a.log.Info("starting echo server", zap.String("addr", addr))
		err := e.Start(addr)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Fatal("running server", zap.Error(err))
		}
	}()

	// Handle graceful shutdown
	shutdownTimeout := 10 * time.Second
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		a.log.Fatal("stopping server", zap.Error(err))
	}
}

func newEchoServer(apiHandlers *api.API) *echo.Echo {

	e := echo.New()
	e.GET("/", apiHandlers.Hello)

	e.POST("/api/measurements", apiHandlers.Measurements)

	// TODO: add sensor CRUD API

	return e
}
