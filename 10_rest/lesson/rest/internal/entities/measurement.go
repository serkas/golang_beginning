package entities

type Measurement struct {
	SensorID   int                   `json:"sensor_id"`
	Timestamp  uint64                `json:"timestamp"`
	Parameters MeasurementParameters `json:"parameters"`
}

type MeasurementParameters struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
