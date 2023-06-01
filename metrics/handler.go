package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func GetPrometheusHandler() http.Handler {
	return promhttp.Handler()
}
