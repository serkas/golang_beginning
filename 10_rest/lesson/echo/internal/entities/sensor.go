package entities

import "time"

type Sensor struct {
	ID        int
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
