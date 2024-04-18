package storage

import (
	"sync"
	"time"

	"layered-service/internal/entities"
)

type MemStore struct {
	mx sync.Mutex

	sensors      []*entities.Sensor
	measurements []*entities.Measurement
}

func NewMemStorage() *MemStore {
	return &MemStore{
		sensors: []*entities.Sensor{
			{
				ID:        10,
				Name:      "sensor 10",
				Type:      "air",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}
}
