package main

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
	)
)

func main() {
	// Register our metric
	prometheus.MustRegister(requestsTotal)

	// Health endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		requestsTotal.Inc()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Metrics endpoint for Prometheus to scrape
	http.Handle("/metrics", promhttp.Handler())

	// Run server
	http.ListenAndServe(":8080", nil)
}
