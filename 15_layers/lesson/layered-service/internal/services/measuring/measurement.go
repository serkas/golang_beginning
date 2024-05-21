package measuring

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"layered-service/internal/entities"
)

type SensorStore interface {
	ListSensors(ctx context.Context) ([]*entities.Sensor, error)
	GetSensor(ctx context.Context, id int64) (*entities.Sensor, error)
	CreateSensor(ctx context.Context, s *entities.Sensor) error
	UpdateSensor(ctx context.Context, s *entities.Sensor) error
	DeleteSensor(ctx context.Context, id int64) error
}

type Store interface {
	SensorStore

	CreateSensorMeasurement(ctx context.Context, measurement *entities.Measurement) error
}

type Service struct {
	log   *zap.Logger
	store Store
}

func NewService(logger *zap.Logger, store Store) *Service {
	return &Service{
		log:   logger,
		store: store,
	}
}

// AddMeasurement creates a new measurement record in the system
// The corresponding sensor object must be already registered in the system
func (s *Service) AddMeasurement(ctx context.Context, measurement *entities.Measurement) error {
	s.log.Info("storing new sensor measurement", zap.Int64("sensorID", measurement.SensorID))

	_, err := s.store.GetSensor(ctx, measurement.SensorID)
	if err != nil {
		return fmt.Errorf("finding sensor: %w", err)
	}
	// other business logic...

	err = s.store.CreateSensorMeasurement(ctx, measurement)
	if err != nil {
		return fmt.Errorf("storing sensor measurement: %w", err)
	}

	return nil
}
