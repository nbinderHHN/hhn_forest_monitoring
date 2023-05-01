package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
)

func (client Client) IncreaseHttpCall(statusCode int) {
	httpCalls.With(prometheus.Labels{statusCodeLabel: strconv.Itoa(statusCode)}).Inc()
}

var httpCalls = promauto.NewCounterVec(
	prometheus.CounterOpts{Name: "http_calls"},
	[]string{statusCodeLabel},
)

var statusCodeLabel = "status_code"
