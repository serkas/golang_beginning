package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"proj/lessons/19_redis/lesson/cached/internal/api"
	"proj/lessons/19_redis/lesson/cached/internal/cache"
	"proj/lessons/19_redis/lesson/cached/internal/services/items"
	"proj/lessons/19_redis/lesson/cached/internal/storage"
	"time"
)

type App struct {
	conf Config
}

func New(conf Config) (*App, error) {
	return &App{
		conf: conf,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	sqlDB, err := sql.Open("mysql", a.conf.DB)
	if err != nil {
		return fmt.Errorf("db initialization: %w", err)
	}
	defer sqlDB.Close()

	store := storage.New(sqlDB)

	redisCli := redis.NewClient(&redis.Options{
		Addr:     a.conf.RedisAddress,
		DB:       0, // use default DB
		Protocol: 3, // specify 2 for RESP 2 or 3 for RESP 3
	})
	err = redisCli.Ping(context.Background()).Err()
	if err != nil {
		return fmt.Errorf("connecting to redis: %w", err)
	}

	defer redisCli.Close()

	itemsService := items.New(store, cache.New(redisCli), items.NewViewsTracker(redisCli))
	server := api.NewServer(a.conf.ServerAddress, itemsService)

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			log.Printf("shutdonw: %s", err)
		}
	}()

	err = server.Serve()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
