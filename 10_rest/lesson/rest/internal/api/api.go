package api

import (
	"net/http"

	"github.com/gorilla/mux"
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

func (a *API) CreateRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", a.Hello).Methods(http.MethodGet)
	r.HandleFunc("/api/measurements", a.CreateMeasurement).Methods(http.MethodPost)

	r.HandleFunc("/api/sensors", a.ListSensors).Methods(http.MethodGet)
	r.HandleFunc("/api/sensors/{id}", a.GetSensor).Methods(http.MethodGet)
	r.HandleFunc("/api/sensors", a.CreateSensor).Methods(http.MethodPost)
	r.HandleFunc("/api/sensors/{id}", a.UpdateSensor).Methods(http.MethodPut)
	r.HandleFunc("/api/sensors/{id}", a.DeleteSensor).Methods(http.MethodDelete)

	return r
}
