package api

import (
	"go.uber.org/zap"
	"rest-server-demo/internal/services/sensors"
)

type API struct {
	log     *zap.Logger
	sensors *sensors.Service
}

func New(logger *zap.Logger, sensorsService *sensors.Service) *API {
	return &API{
		log:     logger,
		sensors: sensorsService,
	}
}
