package client

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RecordHealthStatusEntry
// Records the health status of a client.
// uses the same senderLabel as signal_strength.go
func RecordHealthStatusEntry(clientName string, healthStatus float64) {
	signal.With(prometheus.Labels{clientNameLabel: clientName}).Set(healthStatus)
}

var status = promauto.NewGaugeVec(
	prometheus.GaugeOpts{Name: "health_status", Help: "Health status of a client"},
	[]string{clientNameLabel})
