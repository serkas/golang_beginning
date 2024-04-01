package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"rest-server-demo/internal/entities"
	"rest-server-demo/internal/services/sensors"
)

func (a *API) Measurements(w http.ResponseWriter, r *http.Request) {
	a.log.Info("got measurement")

	var meas entities.Measurement
	err := json.NewDecoder(r.Body).Decode(&meas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.sensors.StoreMeasurement(r.Context(), &meas)
	var dErr sensors.DataError
	if errors.As(err, &dErr) {
		a.log.Warn("sensors error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
