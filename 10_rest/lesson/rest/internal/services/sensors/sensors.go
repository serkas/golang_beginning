package sensors

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"rest-server-demo/internal/entities"
)

type Store interface {
	ListSensors(ctx context.Context) ([]*entities.Sensor, error)
	GetSensor(ctx context.Context, id int64) (*entities.Sensor, error)
	CreateSensor(ctx context.Context, s *entities.Sensor) error
	UpdateSensor(ctx context.Context, s *entities.Sensor) error
	DeleteSensor(ctx context.Context, id int64) error
	CreateSensorMeasurement(ctx context.Context, measurement *entities.Measurement) error
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
	s.log.Info("storing new sensor measurement", zap.Int64("sensorID", measurement.SensorID))

	_, err := s.store.GetSensor(ctx, measurement.SensorID)
	if errors.Is(err, entities.ErrNotFound) {
		return ErrEntityNotFound
	}
	if err != nil {
		return fmt.Errorf("finding sensor: %w", err)
	}

	err = s.store.CreateSensorMeasurement(ctx, measurement)
	if err != nil {
		return fmt.Errorf("storing sensor measurement: %w", err)
	}

	return nil
}

func (s *Service) ListSensors(ctx context.Context) ([]*entities.Sensor, error) {
	return s.store.ListSensors(ctx)
}

func (s *Service) CreateSensor(ctx context.Context, sensor *entities.Sensor) error {
	return s.store.CreateSensor(ctx, sensor)
}

func (s *Service) UpdateSensor(ctx context.Context, sensor *entities.Sensor) error {
	return s.store.UpdateSensor(ctx, sensor)
}

func (s *Service) GetSensor(ctx context.Context, id int64) (*entities.Sensor, error) {
	sensor, err := s.store.GetSensor(ctx, id)
	if errors.Is(err, entities.ErrNotFound) {
		return nil, ErrEntityNotFound
	}

	return sensor, nil
}

func (s *Service) DeleteSensor(ctx context.Context, id int64) error {
	return s.store.DeleteSensor(ctx, id)
}
