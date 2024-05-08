package services

import "time"

type NowTimeProvider func() time.Time
