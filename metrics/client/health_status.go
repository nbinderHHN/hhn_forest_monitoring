package client

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RecordHealthStatusEntry
// Records the health status of a client.
func RecordHealthStatusEntry(clientName string, healthStatus bool) {
	if healthStatus {
		status.With(prometheus.Labels{clientNameLabel: clientName}).Set(1)
	} else {
		status.With(prometheus.Labels{clientNameLabel: clientName}).Set(0)
	}
}

var status = promauto.NewGaugeVec(
	prometheus.GaugeOpts{Name: "health_status", Help: "Health status of a client"},
	[]string{clientNameLabel})
