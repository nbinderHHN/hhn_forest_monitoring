package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"hhn_forest_monitoring/metrics"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	go randomMetrics()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func randomMetrics() {
	client := metrics.Client{}
	for {
		randInt := rand.Intn(10)
		randTime := randInt * rand.Int()
		if randInt%2 == 0 {
			client.RecordHttpCall("/test/my/path", http.StatusOK, int64(randTime))
		} else if randInt%3 == 0 {
			client.RecordHttpCall("/test/my/path", http.StatusNotFound, int64(randTime))
		} else {
			client.RecordHttpCall("/test/my/path", http.StatusUnauthorized, int64(randTime))
		}
		time.Sleep(10 * time.Second)
	}
}
