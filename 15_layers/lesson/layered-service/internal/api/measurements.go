package api

import (
	"encoding/json"
	"layered-service/internal/api/dto"
	"layered-service/internal/entities"
	"net/http"
)

func (a *API) CreateMeasurement(w http.ResponseWriter, r *http.Request) {
	a.log.Info("got measurement")

	var meas dto.Measurement
	err := json.NewDecoder(r.Body).Decode(&meas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.measurements.AddMeasurement(r.Context(), mapMeasurement(&meas))
	if err != nil {
		a.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func mapMeasurement(m *dto.Measurement) *entities.Measurement {
	return &entities.Measurement{
		SensorID:  int64(m.SensorID),
		Timestamp: int64(m.Timestamp),
		Parameters: entities.MeasurementParameters{
			Temperature: m.Temperature,
			Humidity:    m.Humidity,
		},
	}
}
