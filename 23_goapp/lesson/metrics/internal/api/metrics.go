package api

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var idRemoverRegexp = regexp.MustCompile(`\d+`)

func metricsMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer durationMetric(prepareEndpointForMetrics(r.URL.Path), r.Method)()

		next(w, r)
		//statusMetric(prepareEndpointForMetrics(r.URL.Path), )
	}
}

var apiDurationMetric = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "api_response_seconds",
	Help: "API request handling duration in seconds",
}, []string{"endpoint", "method"})

func durationMetric(endpoint, method string) func() {
	startTime := time.Now()

	return func() {
		apiDurationMetric.WithLabelValues(endpoint, method).Observe(time.Since(startTime).Seconds())
	}
}

var apiStatusMetric = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "api_response_status",
	Help: "API response statuses",
}, []string{"endpoint", "method"})

func statusMetric(endpoint, method string) func() {
	return func() {
		apiStatusMetric.WithLabelValues(endpoint, method).Inc()
	}
}

func prepareEndpointForMetrics(path string) string {
	path, _ = strings.CutPrefix(path, "/api/v1")
	return idRemoverRegexp.ReplaceAllString(path, "*")
}
