package main

import (
	"log"
	"os"
	"net/http"
	"onecsc/handlers"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/status", handlers.StatusHandler)
 	http.HandleFunc("/metrics", handlers.MetricsHandler)
 	http.HandleFunc("/crash", handlers.CrashHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port=  "8080"
}

	log.Println("Starting server on port ",port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

