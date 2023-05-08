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
		randInt := rand.Int()
		if randInt%2 == 0 {
			client.IncreaseHttpCall(200, "/test/my/path")
		} else if randInt%3 == 0 {
			client.IncreaseHttpCall(404, "/test/my/path")
		} else {
			client.IncreaseHttpCall(401, "/test/my/path")
		}
		time.Sleep(2 * time.Second)
	}
}
