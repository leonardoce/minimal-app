package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootPage)
	// https://prometheus.io/docs/guides/go-application/
	mux.Handle("/metrics", promhttp.Handler())

	log.Println("Starting web server")
	err := http.ListenAndServe(":5000", handlers.LoggingHandler(os.Stdout, mux))
	if err != nil {
		panic(err)
	}
}
