package storage

import (
	"context"

	"rest-server-demo/internal/entities"
)

func (s *MemStore) CreateSensorMeasurement(_ context.Context, meas *entities.Measurement) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.measurements = append(s.measurements, meas)

	return nil
}
