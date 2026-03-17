package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DevOps Platform is running 🚀\n")
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Server started on :8080")

	http.ListenAndServe("0.0.0.0:8080", nil)
}
