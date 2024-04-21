package storage

import (
	"context"
	"sync"

	"proj/lessons/16_testing/lesson/functional/model"
)

type MemStore struct {
	mx sync.Mutex

	items []*model.Item
}

func NewMemStorage() *MemStore {
	return &MemStore{
		items: make([]*model.Item, 0),
	}
}

func (s *MemStore) ListItems(_ context.Context) ([]*model.Item, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	return s.items[:], nil
}

func (s *MemStore) AddItem(_ context.Context, item *model.Item) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.items = append(s.items, item)

	return nil
}
