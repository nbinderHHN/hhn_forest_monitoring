package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func RecordBatteryStatusEntry(clientName string, batteryLevel float64) {
	batteryStatus.With(prometheus.Labels{clientNameLabel: clientName}).Set(batteryLevel)
}

var batteryStatus = promauto.NewGaugeVec(
	prometheus.GaugeOpts{Name: "battery_status", Help: "Battery status of a client"},
	[]string{clientNameLabel},
)

const clientNameLabel = "client"
