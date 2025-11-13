package main

import (
	"eda-payments/internal"
	"eda-payments/internal/web"
	"log"
	"net/http"
)

const (
	PORT = ":3002"
)

func main() {

	client := internal.NewKafkaClient()
	defer client.Close()

	server := web.NewServer(client)

	log.Println("Starting payment service...")

	err := http.ListenAndServe(PORT, server)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
