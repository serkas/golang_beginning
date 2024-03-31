package sensors

import (
	"context"
	"echo-server-demo/internal/entities"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type Store interface {
	FindSensorByID(ctx context.Context, id int) (*entities.Sensor, error)
	StoreSensorMeasurement(ctx context.Context, measurement *entities.Measurement) error
}

type Service struct {
	log   *zap.Logger
	store Store
}

func New(logger *zap.Logger, store Store) *Service {
	return &Service{
		log:   logger,
		store: store,
	}
}

func (s *Service) StoreMeasurement(ctx context.Context, measurement *entities.Measurement) error {
	s.log.Info("storing new sensor measurement", zap.Int("sensorID", measurement.SensorID))

	_, err := s.store.FindSensorByID(ctx, measurement.SensorID)
	if errors.Is(err, entities.ErrNotFound) {
		return ErrSensorNotFound
	}
	if err != nil {
		return fmt.Errorf("finding sensor: %w", err)
	}

	err = s.store.StoreSensorMeasurement(ctx, measurement)
	if err != nil {
		return fmt.Errorf("storing sensor measurement: %w", err)
	}

	return nil
}
