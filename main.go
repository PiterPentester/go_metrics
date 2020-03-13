package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cmpErr = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "compilation_error",
		Help: "Count of compilation errors",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cmpErr)

}

func main() {
	cmpErr.Inc()

	for i := 0; i < 10; i++ {
		cmpErr.Inc()
	}
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
