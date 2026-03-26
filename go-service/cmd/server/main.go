package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "unknown"
	}

	// 🔥 отдельный handler для /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "DevOps Platform is running 🚀 (%s)\n", version)
	})

	// ✅ метрики отдельно
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server started on :8080")

	http.ListenAndServe("0.0.0.0:8080", nil)
}
