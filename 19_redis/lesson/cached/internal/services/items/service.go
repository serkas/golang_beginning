package items

import (
	"context"
	"errors"
	"fmt"
	"log"
	"proj/lessons/19_redis/lesson/cached/internal/cache"
	"proj/lessons/19_redis/lesson/cached/internal/model"
	"time"
)

type Store interface {
	ListItems(ctx context.Context) ([]*model.Item, error)
	GetItem(ctx context.Context, id int) (*model.Item, error)
	AddItem(ctx context.Context, item *model.Item) error

	GetTopViewedItems(ctx context.Context, limit int) ([]*model.Item, error)
}

type Cache interface {
	GetItems(ctx context.Context, key string) ([]*model.Item, error)
	SetItems(ctx context.Context, key string, items []*model.Item, exp time.Duration) error
}

type Service struct {
	store Store
	cache Cache
}

func New(store Store, cache Cache) *Service {
	return &Service{
		store: store,
		cache: cache,
	}
}

func (s *Service) List(ctx context.Context) (result []*model.Item, err error) {
	return s.store.ListItems(ctx)
}

func (s *Service) Get(ctx context.Context, id int) (item *model.Item, err error) {
	return s.store.GetItem(ctx, id)
}

func (s *Service) Add(ctx context.Context, item *model.Item) error {
	return s.store.AddItem(ctx, item)
}

func (s *Service) GetTopViewed(ctx context.Context, numTopItems int) (result []*model.Item, err error) {
	items, err := s.cache.GetItems(ctx, "top_viewed_items")
	if err != nil && !errors.Is(err, cache.ErrNotFound) {
		return nil, fmt.Errorf("getting top viewed items from cahce: %w", err)
	}

	if errors.Is(err, cache.ErrNotFound) {
		items, err = s.store.GetTopViewedItems(ctx, numTopItems)
		if err != nil {
			return nil, fmt.Errorf("getting top viewed items: %w", err)
		}

		err = s.cache.SetItems(ctx, "top_viewed_items", items, time.Minute)
		if err != nil {
			log.Printf("setting top viewed items: %s", err)
		}
	}

	return items, nil
}
