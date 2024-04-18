package entities

type Measurement struct {
	SensorID   int64
	Timestamp  int64
	Parameters MeasurementParameters
}

type MeasurementParameters struct {
	Temperature float64
	Humidity    float64
}
