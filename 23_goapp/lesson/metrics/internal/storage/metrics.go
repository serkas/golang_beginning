package storage

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var dbOpDurationMetric = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "db_operation_seconds",
	Help: "Duration of db operation in seconds",
}, []string{"operation"})

func durationMetric(operation string) func() {
	startTime := time.Now()

	return func() {
		dbOpDurationMetric.WithLabelValues(operation).Observe(time.Since(startTime).Seconds())
	}
}
