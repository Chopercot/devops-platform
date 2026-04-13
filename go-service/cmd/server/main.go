package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// метрика с label'ами
var httpRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "path", "status"},
)

// регистрация метрики
func init() {
	prometheus.MustRegister(httpRequests)
}

// структура для перехвата статус-кода
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

// перехват WriteHeader
func (r *statusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

// основной handler
func handler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "unknown"
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "DevOps Platform is running 🚀 (%s)\n", version)
}

// middleware для метрик
func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     200,
		}

		next.ServeHTTP(recorder, r)

		httpRequests.WithLabelValues(
			r.Method,
			r.URL.Path,
			strconv.Itoa(recorder.statusCode),
		).Inc()
	})
}

func main() {

	http.Handle("/", metricsMiddleware(http.HandlerFunc(handler)))
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server started on :8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
