package main

import (
	"apigateway/gateway"
	"log"
)

func main() {

	server := gateway.NewServer()
	err := server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
