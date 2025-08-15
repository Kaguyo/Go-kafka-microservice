package main

import (
	"github.com/your/project/internal/config"
	"github.com/your/project/internal/consumer"
)

func main() {
	cfg := config.Load()
	consumer.Start(cfg)
}
