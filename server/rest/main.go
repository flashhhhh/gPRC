package main

import (
	"log"
	"net/http"

	"flash/server/rest/handler"
)

func main() {
	http.HandleFunc("/add", handler.AddHandler)
	http.HandleFunc("/subtract", handler.SubtractHandler)
	http.HandleFunc("/multiply", handler.MultiplyHandler)
	http.HandleFunc("/division", handler.DivisionHandler)
	http.HandleFunc("/power", handler.PowerHandler)

	log.Println("Starting server on port :1906")
	if err := http.ListenAndServe(":1906", nil); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}