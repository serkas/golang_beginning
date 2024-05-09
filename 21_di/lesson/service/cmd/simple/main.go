package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"log"
	"proj/lessons/21_di/lesson/service/internal/api"
	"proj/lessons/21_di/lesson/service/internal/app"
	"proj/lessons/21_di/lesson/service/internal/cache"
	"proj/lessons/21_di/lesson/service/internal/services"
	"proj/lessons/21_di/lesson/service/internal/services/items"
	"proj/lessons/21_di/lesson/service/internal/storage"
	"time"
)

// References:
// https://uber-go.github.io/fx/get-started/minimal.html
// https://medium.com/@erez.levi/using-uber-fx-to-simplify-dependency-injection-875363245c4c

func main() {
	fxApp := fx.New(
		fx.Provide(app.ReadConfig),
		fx.Provide(newDBClient),
		fx.Provide(storage.New),

		fx.Provide(
			newRedisClient,
			cache.New,
			items.NewViewsTracker,
			fx.Annotate(items.New, fx.As(new(api.ItemsService))), // dependency is an interface, hint which implementation to use. Link: https://uber-go.github.io/fx/annotate.html#casting-structs-to-interfaces
		),
		fx.Provide(func(conf *app.Config, is api.ItemsService) *api.Server {
			return api.NewServer(conf.ServerAddress, is) // NewServer requires string argument for address. Of course, we can rewrite the server constructor to accept config type, but this can be impossible with external dependency
		}),
		fx.Provide(func() services.NowTimeProvider {
			return func() time.Time { return time.Now() }
		}),
		fx.Invoke(runAPIServer),
	)

	fxApp.Run()
}

func runAPIServer(lifecycle fx.Lifecycle, s *api.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return s.Start()
		},
		OnStop: func(ctx context.Context) error {
			log.Println("shutdown")
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			return s.Shutdown(ctx)
		},
	})
}

func newDBClient(lc fx.Lifecycle, conf *app.Config) (*sql.DB, error) {
	sqlDB, err := sql.Open("mysql", conf.DB)
	if err != nil {
		return nil, fmt.Errorf("db initialization: %w", err)
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("closing DB connection")
			return sqlDB.Close()
		},
	})

	return sqlDB, nil
}

func newRedisClient(lc fx.Lifecycle, conf *app.Config) (*redis.Client, error) {
	redisCli := redis.NewClient(&redis.Options{Addr: conf.RedisAddress})
	err := redisCli.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("connecting to redis: %w", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("closing Redis connection")
			return redisCli.Close()
		},
	})

	return redisCli, nil
}
