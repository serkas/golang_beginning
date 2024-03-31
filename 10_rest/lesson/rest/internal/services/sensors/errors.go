package sensors

var (
	ErrSensorNotFound = DataError("sensor_not_found")
)

type DataError string

func (e DataError) Error() string {
	return string(e)
}
