package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

const (
	updatePeriod = 1 * time.Second
	serverURL    = "http://localhost:8001/api/measurements"
)

func main() {
	sens := newSensor(12)
	ticker := time.NewTicker(updatePeriod)
	for range ticker.C {
		slog.Info("now sending one measurement")
		err := sens.sendOneMeasurement()
		if err != nil {
			slog.Error("sending one measurement", "err", err)
		}
	}
}

type sensor struct {
	id  int64
	cli *http.Client
}

func newSensor(id int64) *sensor {
	return &sensor{
		id:  id,
		cli: http.DefaultClient,
	}
}

func (s *sensor) sendOneMeasurement() error {
	m := Measurement{
		SensorID:  s.id,
		Timestamp: time.Now().Unix(),
		Parameters: map[string]float64{
			"temperature": 23.4,
			"humidity":    65.4,
		},
	}

	payload, err := json.Marshal(&m)
	if err != nil {
		return fmt.Errorf("marshaling payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, serverURL, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	resp, err := s.cli.Do(req)
	if err != nil {
		return fmt.Errorf("doing request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

type Measurement struct {
	SensorID   int64              `json:"sensor_id"`
	Timestamp  int64              `json:"timestamp"`
	Parameters map[string]float64 `json:"parameters"`
}
