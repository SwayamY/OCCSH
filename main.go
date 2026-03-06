package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"onecsc/handlers"
	"onecsc/middleware"
)

var crashChan = make(chan struct{})

func createServer() *http.Server {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/status", handlers.StatusHandler)
	mux.HandleFunc("/metrics", handlers.MetricsHandler)
	mux.HandleFunc("/crash", crashHandler)

	handler := middleware.Logging(mux)

	return &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
}

func crashHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Crash triggered")

	w.Write([]byte("Server crashing..."))

	go func() {
		crashChan <- struct{}{}
	}()
}

func runServer() error {

	server := createServer()

	go func() {
		log.Println("Server starting on", server.Addr)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("Server error:", err)
		}
	}()

	select {

	case <-crashChan:
		log.Println("Crash signal received")

	case <-time.After(24 * time.Hour):
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}

func main() {

	log.Println("Starting OCCSH service supervisor")

	for {

		err := runServer()

		if err != nil {
			log.Println("Server stopped:", err)
		}

		log.Println("Restarting service in 2 seconds...")

		time.Sleep(2 * time.Second)
	}
}
