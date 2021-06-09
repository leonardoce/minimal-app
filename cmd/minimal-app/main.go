package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/lib/pq"
)

var (
	color       string
	environment string
	query       string
	db          *sql.DB
)

var (
	// https://prometheus.io/docs/guides/go-application/
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minimal_app_requests",
		Help: "The total number of processed requests",
	})

	txProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minimal_app_tx",
		Help: "The total number of processed transactions",
	})

	errorsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minimal_app_errors",
		Help: "The total number of errors",
	})
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "<p style='color: %v; text-weight: 2000'>Environment: %v</p>",
		color,
		environment,
	)

	opsProcessed.Inc()
}

func checkPage(w http.ResponseWriter, r *http.Request) {
	err := db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintln(w, err.Error())
		errorsProcessed.Inc()
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "ok")
}

func txPage(w http.ResponseWriter, r *http.Request) {
	txProcessed.Inc()
	_, err := db.Exec(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintln(w, err.Error())
		errorsProcessed.Inc()
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "ok")
}

func main() {
	var err error

	color = os.Getenv("COLOR")
	environment = os.Getenv("ENVIRONMENT")
	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = ":5000"
	}

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	query = os.Getenv("QUERY")

	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/tx", txPage)
	router.HandleFunc("/.well-known/check", checkPage)
	router.Handle("/metrics", promhttp.Handler())

	// https://prometheus.io/docs/guides/go-application/

	log.Printf("Starting web server on %v", listenAddress)
	err = http.ListenAndServe(listenAddress, handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		panic(err)
	}
}
