package storage

import (
	"context"
	"sync"

	"rest-server-demo/internal/entities"
)

type MemStore struct {
	mx sync.Mutex

	sensors      []*entities.Sensor
	measurements []*entities.Measurement
}

func NewMemStorage() *MemStore {
	return &MemStore{}
}

func (s *MemStore) FindSensorByID(_ context.Context, id int) (*entities.Sensor, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	for _, sensor := range s.sensors {
		if sensor.ID == id {
			return sensor, nil
		}
	}

	return nil, entities.ErrNotFound
}
