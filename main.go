package main

import (
	"go-kafka-microservice/internal/config"
	"go-kafka-microservice/internal/consumer"
)

func main() {
	cfg := config.Load()
	//logger
	consumer.Start(cfg)
}
