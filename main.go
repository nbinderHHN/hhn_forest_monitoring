package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	go randomMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

func randomMetrics() {
	for {
		if rand.Int()%2 == 0 {
			fmt.Println("Increase")
			randomCounter.Inc()
		} else {
			fmt.Println("No Increase")
		}
		time.Sleep(2 * time.Second)
	}
}

var randomCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "random_counter",
})
