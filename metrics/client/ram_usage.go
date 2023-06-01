package client

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RecordRamUsage
// Records the RAM usage of a specific client.
func RecordRamUsage(clientName string, ramUsageValue float64) {
	ramUsage.With(prometheus.Labels{clientNameLabel: clientName}).Set(ramUsageValue)
}

var ramUsage = promauto.NewGaugeVec(
	prometheus.GaugeOpts{Name: "ram_usage", Help: "RAM usage of a given client"},
	[]string{clientNameLabel},
)
