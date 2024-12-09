package metric

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartMetricServer(host string, port int) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(host+":"+strconv.Itoa(port), mux)
}
