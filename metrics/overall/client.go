package overall

import "github.com/prometheus/client_golang/prometheus"

type client struct {
	metrics map[string]*prometheus.GaugeVec
}
