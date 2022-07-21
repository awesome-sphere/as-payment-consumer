package main

import (
	"log"

	"github.com/awesome-sphere/as-payment-consumer/db"
	"github.com/awesome-sphere/as-payment-consumer/internal"
	"github.com/awesome-sphere/as-payment-consumer/kafka"
)

func main() {
	internal.InitializeInternalServices()
	db.InitializeDatabase()
	log.Println("Starting kafka...")
	kafka.InitializeKafka()
	for {
	}
}
