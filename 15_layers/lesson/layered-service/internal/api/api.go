package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"layered-service/internal/entities"
	"layered-service/internal/services/measuring"
	"net/http"
)

type API struct {
	log           *zap.Logger
	measurements  *measuring.Service
	sensorStorage measuring.SensorStore
}

func New(logger *zap.Logger, measurements *measuring.Service, sensorStorage measuring.SensorStore) *API {
	return &API{
		log:           logger,
		measurements:  measurements,
		sensorStorage: sensorStorage,
	}
}

func (a *API) CreateRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/measurements", a.CreateMeasurement).Methods(http.MethodPost)

	r.HandleFunc("/api/sensors", a.ListSensors).Methods(http.MethodGet)
	r.HandleFunc("/api/sensors", a.CreateSensor).Methods(http.MethodPost)
	r.HandleFunc("/api/sensors/{id}", a.GetSensor).Methods(http.MethodGet)
	r.HandleFunc("/api/sensors/{id}", a.UpdateSensor).Methods(http.MethodPut)
	r.HandleFunc("/api/sensors/{id}", a.DeleteSensor).Methods(http.MethodDelete)

	r.HandleFunc("/api/analytics/general", a.GetGeneralAnalytics).Methods(http.MethodGet)

	return r
}

func (a *API) returnJSON(w http.ResponseWriter, status int, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		a.log.Error("marshaling response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(resp)
	if err != nil {
		a.log.Error("sending response", zap.Error(err))
	}
}

func (a *API) handleError(w http.ResponseWriter, err error) {
	var status int
	if errors.Is(err, entities.ErrNotFound) {
		status = http.StatusNotFound
	} else {
		status = http.StatusInternalServerError
		a.log.Error("unexpected error", zap.Error(err))
	}

	a.returnJSON(w, status, &Error{Error: err.Error()})
}

type Error struct {
	Error string `json:"error"`
}
