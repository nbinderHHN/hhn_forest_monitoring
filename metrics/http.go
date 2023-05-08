package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
)

// IncreaseHttpCall Increases the amount of http calls to a given endpoint with the statusCode
func (client Client) IncreaseHttpCall(endpoint string, statusCode int) {
	httpCalls.With(prometheus.Labels{statusCodeLabel: strconv.Itoa(statusCode), endpointLabel: endpoint}).Inc()
}

var httpCalls = promauto.NewCounterVec(
	prometheus.CounterOpts{Name: "http_calls"},
	[]string{statusCodeLabel, endpointLabel},
)

const statusCodeLabel = "status_code"
const endpointLabel = "endpoint"
