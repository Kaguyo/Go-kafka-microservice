package main

import (
	"./internal/config"
	"./internal/consumer"
)

func main() {
	cfg := config.Load()
	consumer.Start(cfg)
}
