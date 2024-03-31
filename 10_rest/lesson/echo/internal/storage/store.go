package storage

import (
	"context"
	"echo-server-demo/internal/entities"
	"sync"
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

func (s *MemStore) StoreSensorMeasurement(_ context.Context, meas *entities.Measurement) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.measurements = append(s.measurements, meas)

	return nil
}
