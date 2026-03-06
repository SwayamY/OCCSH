package handlers

import (
	"encoding/json"
//	"log"
	"net/http"
	//"os"
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
	w.Write([]byte("OK"))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	incrementRequest()

	response := map[string]string{
		"status": "running",
	}

	json.NewEncoder(w).Encode(response)
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {

	incrementRequest()

	uptime := time.Since(startTime).Seconds()

	response := map[string]interface{}{
		"uptime_seconds": uptime,
		"request_count":  atomic.LoadUint64(&requestCount),
	}

	json.NewEncoder(w).Encode(response)
}
//func CrashHandler(w http.ResponseWriter, r *http.Request) {
///
///	log.Println("Crash triggered")

	// simulate real crash
//	os.Exit(1)
//}
