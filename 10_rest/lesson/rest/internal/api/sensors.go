package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"rest-server-demo/internal/entities"
	"rest-server-demo/internal/services/sensors"
)

func (a *API) ListSensors(w http.ResponseWriter, r *http.Request) {
	sensorsList, err := a.sensors.ListSensors(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}

	resp, err := json.Marshal(sensorsList)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		a.log.Error("sending response", zap.Error(err))
	}
}

func (a *API) GetSensor(w http.ResponseWriter, r *http.Request) {
	// read and validate input
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// do some logic
	s, err := a.sensors.GetSensor(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}

	// show the result
	resp, err := json.Marshal(s)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		a.log.Error("sending response", zap.Error(err))
	}
}

func (a *API) CreateSensor(w http.ResponseWriter, r *http.Request) {
	var sensor entities.Sensor
	err := json.NewDecoder(r.Body).Decode(&sensor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.sensors.CreateSensor(r.Context(), &sensor)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *API) UpdateSensor(w http.ResponseWriter, r *http.Request) {
	var sensor entities.Sensor
	err := json.NewDecoder(r.Body).Decode(&sensor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil || id != sensor.ID {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.sensors.UpdateSensor(r.Context(), &sensor)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) DeleteSensor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.sensors.DeleteSensor(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleError(w http.ResponseWriter, err error) {
	var entityError sensors.DataError
	if errors.As(err, &entityError) {
		http.Error(w, entityError.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
