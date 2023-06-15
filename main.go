package main

import (
	"github.com/nbinderHHN/hhn_forest_monitoring/metrics"
	"github.com/nbinderHHN/hhn_forest_monitoring/metrics/backend"
	"github.com/nbinderHHN/hhn_forest_monitoring/metrics/client"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	go randomMetrics()

	http.Handle("/metrics", metrics.GetPrometheusHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func randomMetrics() {
	for {
		randInt := rand.Intn(10)
		randTime := randInt * rand.Int()
		if randInt%2 == 0 {
			backend.RecordHttpCall("/test/my/path", http.StatusOK, int64(randTime))
			client.RecordSignalStrengthEntry("Client 2", "Client 1", float64(randTime))
		} else if randInt%3 == 0 {
			backend.RecordHttpCall("/test/my/path", http.StatusNotFound, int64(randTime))
			client.RecordSignalStrengthEntry("Client 2", "Client 5", float64(randTime))
		} else {
			backend.RecordHttpCall("/test/my/path", http.StatusUnauthorized, int64(randTime))
			client.RecordSignalStrengthEntry("Client 2", "Client 1", float64(randTime))
		}
		time.Sleep(10 * time.Second)
	}
}
