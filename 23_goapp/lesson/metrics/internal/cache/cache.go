package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"

	"proj/lessons/23_goapp/lesson/metrics/internal/model"
)

var ErrNotFound = errors.New("not_found")

type Cache struct {
	cli *redis.Client
}

func New(cli *redis.Client) *Cache {
	return &Cache{cli: cli}
}

func (c *Cache) GetItems(ctx context.Context, key string) ([]*model.Item, error) {
	defer func(start time.Time) {
		log.Printf("got top viewed from cache in %v", time.Since(start))
	}(time.Now())

	result, err := c.cli.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("getting from redis: %w", err)
	}

	var items []*model.Item
	err = json.Unmarshal([]byte(result), &items)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling cached value: %w", err)
	}

	return items, nil
}

func (c *Cache) SetItems(ctx context.Context, key string, items []*model.Item, exp time.Duration) error {
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}

	return c.cli.SetEx(ctx, key, data, exp).Err()
}
