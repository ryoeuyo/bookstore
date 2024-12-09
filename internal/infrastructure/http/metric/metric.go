package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics is used to store metrics
type Metrics struct {
	ErrorsCount     prometheus.Counter
	RequestCount    prometheus.Counter
	RequestDuration *prometheus.SummaryVec
}

// NewMetrics returns new metrics
func NewMetrics() *Metrics {
	return &Metrics{
		ErrorsCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "todo_service",
			Subsystem: "http",
			Name:      "errors_count",
			Help:      "The total number of HTTP errors",
		}),
		RequestCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "todo_service",
			Subsystem: "http",
			Name:      "request_count",
			Help:      "The total number of HTTP requests",
		}),
		RequestDuration: promauto.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:  "todo_service",
			Subsystem:  "http",
			Name:       "request_duration",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"status"}),
	}
}
