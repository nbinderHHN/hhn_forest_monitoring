package client

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RecordSignalStrengthEntry
// Records the signal strength from a sender to a receiver.
func RecordSignalStrengthEntry(senderClientName, receiverClientName string, signalStrength float64) {
	signal.With(prometheus.Labels{senderLabel: senderClientName, receiverLabel: receiverClientName}).Set(signalStrength)
}

var signal = promauto.NewGaugeVec(
	prometheus.GaugeOpts{Name: "signal_strength", Help: "Strength of the signal from receiver to sender"},
	[]string{senderLabel, receiverLabel})
