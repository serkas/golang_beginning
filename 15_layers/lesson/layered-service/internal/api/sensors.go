package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"layered-service/internal/entities"
)

// Using open layers, accessing the storage directly from handlers
// Is it acceptable? For simple CRUD where all validation can be done withing handler - probably yes
// If there are some additional logic or dependent objects, use of service layer is more appropriate

func (a *API) ListSensors(w http.ResponseWriter, r *http.Request) {
	sensorsList, err := a.sensorStorage.ListSensors(r.Context())
	if err != nil {
		a.handleError(w, err)
		return
	}

	a.returnJSON(w, http.StatusOK, sensorsList)
}

func (a *API) GetSensor(w http.ResponseWriter, r *http.Request) {
	// read and validate input
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s, err := a.sensorStorage.GetSensor(r.Context(), id)
	if err != nil {
		a.handleError(w, err)
		return
	}

	a.returnJSON(w, http.StatusOK, s)
}

func (a *API) CreateSensor(w http.ResponseWriter, r *http.Request) {
	var sensor entities.Sensor
	err := json.NewDecoder(r.Body).Decode(&sensor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.sensorStorage.CreateSensor(r.Context(), &sensor)
	if err != nil {
		a.handleError(w, err)
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

	err = a.sensorStorage.UpdateSensor(r.Context(), &sensor)
	if err != nil {
		a.handleError(w, err)
		return
	}
}

func (a *API) DeleteSensor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.sensorStorage.DeleteSensor(r.Context(), id)
	if err != nil {
		a.handleError(w, err)
		return
	}
}
