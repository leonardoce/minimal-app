package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	color       string
	environment string
)

var (
	// https://prometheus.io/docs/guides/go-application/
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minimal_app_requests",
		Help: "The total number of processed requests",
	})
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "<p style='color: %v; text-weight: 2000'>Environment: %v</p>",
		color,
		environment,
	)

	opsProcessed.Inc()
}

func main() {
	color = os.Getenv("COLOR")
	environment = os.Getenv("ENVIRONMENT")

	http.HandleFunc("/", rootPage)

	// https://prometheus.io/docs/guides/go-application/
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}
