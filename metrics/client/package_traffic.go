package client

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RecordBytesSentEntry
// Record the number of bytes sent by a client.
func RecordBytesSentEntry(clientName string, bytesSent float64) {
	packageTraffic.With(prometheus.Labels{clientNameLabel: clientName, directionLabel: "sent"}).Add(bytesSent)
}

// RecordBytesReceivedEntry
// Record the number of bytes received by a client.
func RecordBytesReceivedEntry(clientName string, bytesReceived float64) {
	packageTraffic.With(prometheus.Labels{clientNameLabel: clientName, directionLabel: "received"}).Add(bytesReceived)
}

var packageTraffic = promauto.NewCounterVec(
	prometheus.CounterOpts{Name: "package_traffic", Help: "Number of bytes received and sent by a client"},
	[]string{clientNameLabel, directionLabel})
