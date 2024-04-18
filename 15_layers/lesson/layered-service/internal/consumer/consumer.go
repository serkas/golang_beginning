package consumer

import (
	"context"
	"go.uber.org/zap"
	"layered-service/internal/entities"
	"layered-service/internal/services/measuring"
	"math/rand"
	"time"
)

type Consumer struct {
	log     *zap.Logger
	service *measuring.Service
}

func New(log *zap.Logger, s *measuring.Service) *Consumer {
	return &Consumer{
		log:     log,
		service: s,
	}
}

func (c *Consumer) Run(ctx context.Context) error {
	// Fake implementation of a queue consumer
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		m, err := c.consumeOne()
		if err != nil {
			c.log.Error("consuming error", zap.Error(err))
			continue
		}

		c.log.Info("adding measurement from consumer")
		err = c.service.AddMeasurement(ctx, m)
		if err != nil {
			c.log.Error("measurement adding error", zap.Error(err))
		}
	}
}

func (c *Consumer) consumeOne() (*entities.Measurement, error) {
	<-time.NewTimer(time.Second * 5).C

	m := entities.Measurement{
		SensorID:  1,
		Timestamp: time.Now().Add(-time.Minute).Unix(),
		Parameters: entities.MeasurementParameters{
			Temperature: 22 * rand.Float64(),
			Humidity:    40 + 60*rand.Float64(),
		},
	}

	return &m, nil
}
