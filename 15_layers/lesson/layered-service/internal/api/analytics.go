package api

import "net/http"

func (a *API) GetGeneralAnalytics(w http.ResponseWriter, r *http.Request) {
	sensorsList, err := a.sensorStorage.ListSensors(r.Context())
	if err != nil {
		a.handleError(w, err)
		return
	}

	a.returnJSON(w, http.StatusOK, sensorsList)
}
