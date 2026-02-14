package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint hit: %s %s", r.Method, r.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint hit: %s %s", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "running",
	}

	json.NewEncoder(w).Encode(response)
}

