package overall

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func (c client) ExportClientData(typeName, clientName string, value float64) {
	if c.metrics[typeName] == nil {
		c.metrics[typeName] = promauto.NewGaugeVec(
			prometheus.GaugeOpts{Name: typeName},
			[]string{clientNameLabel},
		)
	}
	c.metrics[typeName].With(prometheus.Labels{clientNameLabel: clientName}).Set(value)
}
