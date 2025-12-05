package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type response struct {
	Service string `json:"service"`
	Message string `json:"message"`
	Host    string `json:"host"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		resp := response{
			Service: "service-a",
			Message: "Hello from service-a (Go demo app)!",
			Host:    host,
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	addr := ":8081"
	log.Printf("service-a listening on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("service-a failed: %v", err)
	}
}
