package items

import (
	"context"
	"errors"
	"fmt"
	"log"
	"proj/lessons/21_di/lesson/service/internal/cache"
	"proj/lessons/21_di/lesson/service/internal/model"
	"proj/lessons/21_di/lesson/service/internal/services"
	"proj/lessons/21_di/lesson/service/internal/storage"
	"time"
)

//type Store interface {
//	ListItems(ctx context.Context) ([]*model.Item, error)
//	GetItem(ctx context.Context, id int) (*model.Item, error)
//	AddItem(ctx context.Context, item *model.Item) error
//
//	GetTopLikedItems(ctx context.Context, limit int) ([]*model.Item, error)
//}
//
//type Cache interface {
//	GetItems(ctx context.Context, key string) ([]*model.Item, error)
//	SetItems(ctx context.Context, key string, items []*model.Item, exp time.Duration) error
//}

type Service struct {
	store        *storage.Store
	cache        *cache.Cache
	viewsTracker *ViewsTracker
	now          services.NowTimeProvider
}

func New(store *storage.Store, cache *cache.Cache, viewsTracker *ViewsTracker, now services.NowTimeProvider) *Service {
	return &Service{
		store:        store,
		cache:        cache,
		viewsTracker: viewsTracker,
		now:          now,
	}
}

func (s *Service) List(ctx context.Context) (result []*model.Item, err error) {
	return s.store.ListItems(ctx)
}

func (s *Service) Get(ctx context.Context, id int) (item *model.Item, err error) {
	return s.store.GetItem(ctx, id)
}

func (s *Service) Add(ctx context.Context, item *model.Item) error {
	//item.Timestamp = int(time.Now().Unix())
	item.Timestamp = int(s.now().Unix())
	return s.store.AddItem(ctx, item)
}

func (s *Service) GetTopLiked(ctx context.Context, numTopItems int) (result []*model.Item, err error) {
	cacheKey := "top_liked_items"
	items, err := s.cache.GetItems(ctx, cacheKey)
	if err != nil && !errors.Is(err, cache.ErrNotFound) {
		return nil, fmt.Errorf("getting top liked items from cahce: %w", err)
	}

	if errors.Is(err, cache.ErrNotFound) {
		items, err = s.store.GetTopLikedItems(ctx, numTopItems)
		if err != nil {
			return nil, fmt.Errorf("getting top liked items: %w", err)
		}

		err = s.cache.SetItems(ctx, cacheKey, items, time.Minute)
		if err != nil {
			log.Printf("caching top liked items: %s", err)
		}
	}

	return items, nil
}

func (s *Service) CountView(ctx context.Context, itemID int) error {
	return s.viewsTracker.View(ctx, itemID)
}

func (s *Service) GetTopViewed(ctx context.Context, num int) (items []*model.Item, err error) {
	ids, err := s.viewsTracker.GetTopViewed(ctx, num)
	if err != nil {
		return nil, fmt.Errorf("getting top viewed item ids: %w", err)
	}

	for _, id := range ids {
		item, err := s.store.GetItem(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("getting item by id: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}
