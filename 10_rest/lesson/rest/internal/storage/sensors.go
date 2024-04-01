package storage

import (
	"context"
	"fmt"

	"rest-server-demo/internal/entities"
)

func (s *MemStore) ListSensors(_ context.Context) ([]*entities.Sensor, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	return s.sensors[:], nil
}

func (s *MemStore) CreateSensor(_ context.Context, sensor *entities.Sensor) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	for _, e := range s.sensors {
		if e.ID == sensor.ID {
			return fmt.Errorf("id=%d: %w", sensor.ID, entities.ErrConflict)
		}
	}
	s.sensors = append(s.sensors, sensor)

	return nil
}

func (s *MemStore) GetSensor(_ context.Context, id int) (*entities.Sensor, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	for _, e := range s.sensors {
		if e.ID == id {
			return e, nil
		}
	}

	return nil, entities.ErrNotFound
}

func (s *MemStore) DeleteSensor(_ context.Context, id int) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	for i, e := range s.sensors {
		if e.ID == id {
			s.sensors = append(s.sensors[:i], s.sensors[i+1:]...)
			return nil
		}
	}

	return nil
}
