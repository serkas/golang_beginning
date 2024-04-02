package entities

import "time"

type Sensor struct {
	ID        int64
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
