package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Health endpoint hit")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Status endpoint hit")
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "running",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/status", statusHandler)

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

