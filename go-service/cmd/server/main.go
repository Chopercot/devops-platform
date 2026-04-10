package main

import (
 "fmt"
 "net/http"
 "os"

 "github.com/prometheus/client_golang/prometheus"
 "github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequests = prometheus.NewCounter(
 prometheus.CounterOpts{
  Name: "http_requests_total",
  Help: "Total number of HTTP requests",
 },
)

func handler(w http.ResponseWriter, r *http.Request) {
 httpRequests.Inc()

 version := os.Getenv("APP_VERSION")
 if version == "" {
  version = "unknown"
 }

 fmt.Fprintf(w, "DevOps Platform is running 🚀 (%s)\n", version)
}

func main() {
 prometheus.MustRegister(httpRequests)

 http.HandleFunc("/", handler)
 http.Handle("/metrics", promhttp.Handler())

 fmt.Println("Server started on :8080")
 http.ListenAndServe("0.0.0.0:8080", nil)
}
