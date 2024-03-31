package api

import (
	"echo-server-demo/internal/services/sensors"
	"go.uber.org/zap"
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
