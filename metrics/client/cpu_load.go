package client

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RecordCpuLoad
// Records the CPU load of a specific client.
func RecordCpuLoad(clientName string, cpuLoadValue float64) {
	cpuLoad.With(prometheus.Labels{clientNameLabel: clientName}).Set(cpuLoadValue)
}

var cpuLoad = promauto.NewGaugeVec(
	prometheus.GaugeOpts{Name: "cpu_load", Help: "CPU load of a given client"},
	[]string{clientNameLabel},
)
