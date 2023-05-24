package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func RecordSignalStrengthEntry(receiverClientName, senderClientName string, signalStrength float64) {
	signal.With(prometheus.Labels{receiverLabel: receiverClientName, senderLabel: senderClientName}).Set(signalStrength)
}

var signal = promauto.NewGaugeVec(
	prometheus.GaugeOpts{Name: "signal_strength", Help: "Strength of the signal from receiver to sender"},
	[]string{receiverLabel, senderLabel})

const receiverLabel = "receiver"
const senderLabel = "sender"
