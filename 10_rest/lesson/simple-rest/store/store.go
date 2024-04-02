package store

import (
	"errors"
	"sync"

	"simple-rest/entities"
)

var ErrNotFound = errors.New("entity_not_found")

type MemStore struct {
	mx sync.Mutex

	sensors []*entities.Sensor
}

func NewMemStorage() *MemStore {
	return &MemStore{
		sensors: []*entities.Sensor{
			{
				ID:   1,
				Name: "sensor 1",
				Type: "air",
			},
		},
	}
}

func (s *MemStore) List() ([]*entities.Sensor, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	return s.sensors[:], nil
}

func (s *MemStore) Create(sensor *entities.Sensor) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.sensors = append(s.sensors, sensor)

	return nil
}

func (s *MemStore) Update(sensor *entities.Sensor) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	for i, e := range s.sensors {
		if e.ID == sensor.ID {
			s.sensors[i] = sensor
			return nil
		}
	}

	return ErrNotFound
}

func (s *MemStore) Get(id int64) (*entities.Sensor, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	for _, e := range s.sensors {
		if e.ID == id {
			return e, nil
		}
	}

	return nil, ErrNotFound
}

func (s *MemStore) Delete(id int64) error {
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
