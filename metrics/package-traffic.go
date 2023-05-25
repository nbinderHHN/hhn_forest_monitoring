package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Record the number of bytes sent and received by a client.
func RecordBytesSentEntry(senderClientName string, bytesSent float64) {
	bytesSentCounter.With(prometheus.Labels{senderLabel: senderClientName}).Add(bytesSent)
}

func RecordBytesReceivedEntry(receiverClientName string, bytesReceived float64) {
	bytesReceivedCounter.With(prometheus.Labels{receiverLabel: receiverClientName}).Add(bytesReceived)
}

// Maybe we should fusionate the two counters into one, with a label "direction" that can be "sent" or "received"?

var bytesSentCounter = promauto.NewCounterVec(
	prometheus.CounterOpts{Name: "bytes_sent", Help: "Number of bytes sent by a client"},
	[]string{senderLabel})

var bytesReceivedCounter = promauto.NewCounterVec(
	prometheus.CounterOpts{Name: "bytes_received", Help: "Number of bytes received by a client"},
	[]string{receiverLabel})
