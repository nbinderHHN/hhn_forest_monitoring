package backend

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RecordHttpCall
// Records the internal response time for a given endpoint with the given statusCode.
// Increases the amount of http calls to a given endpoint with the given statusCode.
func RecordHttpCall(endpoint string, statusCode int, responseTime float64) {
	httpResponseTime.With(prometheus.Labels{statusCodeLabel: strconv.Itoa(statusCode), endpointLabel: endpoint}).Set(responseTime)
	httpCalls.With(prometheus.Labels{statusCodeLabel: strconv.Itoa(statusCode), endpointLabel: endpoint}).Inc()
}

var httpResponseTime = promauto.NewGaugeVec(
	prometheus.GaugeOpts{Name: "http_response_time"},
	[]string{statusCodeLabel, endpointLabel},
)

var httpCalls = promauto.NewCounterVec(
	prometheus.CounterOpts{Name: "http_calls"},
	[]string{statusCodeLabel, endpointLabel},
)
