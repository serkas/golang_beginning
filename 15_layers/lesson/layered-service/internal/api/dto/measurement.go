package dto

type Measurement struct {
	SensorID    int     `json:"sensor_id"`
	Timestamp   int     `json:"timestamp"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
