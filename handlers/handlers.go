package handlers

import (
	"encoding/json"
	"log"
	"os"
	"net/http"
	"sync/atomic"
	"time"
)

var requestCount uint64
var startTime = time.Now()

func incrementRequest() {
	atomic.AddUint64(&requestCount, 1)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	incrementRequest()
	log.Printf("Endpoint hit: %s %s", r.Method, r.URL.Path)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	incrementRequest()
	log.Printf("Endpoint hit: %s %s", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "running",
	}

	json.NewEncoder(w).Encode(response)
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	incrementRequest()
	log.Printf("Endpoint hit: %s %s", r.Method, r.URL.Path)

	uptime := time.Since(startTime).Seconds()

	response := map[string]interface{}{
		"uptime_seconds": uptime,
		"request_count":  atomic.LoadUint64(&requestCount),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CrashHandler(w http.ResponseWriter, r *http.Request){
	incrementRequest()
	log.Printf("Endpoint hit: %s %s", r.Method, r.URL.Path)

	log.Println("Simulating service crash..." )
	w.Write([]byte("server will crash now !"))
	os.Exit(1)
	//panic("intentional crash triggered")
}
