package main

import (
	"log"
	"net/http"
	"onecsc/handlers"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/status", handlers.StatusHandler)

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

