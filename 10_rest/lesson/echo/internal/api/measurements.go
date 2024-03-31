package api

import (
	"echo-server-demo/internal/entities"
	"echo-server-demo/internal/services/sensors"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *API) Measurements(c echo.Context) error {
	a.log.Info("got measurement")

	var meas entities.Measurement
	err := json.NewDecoder(c.Request().Body).Decode(&meas)
	if err != nil {
		return err
	}

	err = a.sensors.StoreMeasurement(c.Request().Context(), &meas)
	var dErr sensors.DataError
	if errors.As(err, &dErr) {
		return echo.NewHTTPError(http.StatusBadRequest, dErr)
	}
	return err
}
