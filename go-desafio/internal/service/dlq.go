package service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/your/project/internal/config"
)

func sendDLQ(cfg config.Config, value []byte) {
	p, _ := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"security.protocol": cfg.SecurityProtocol,
		"sasl.mechanisms":   cfg.SaslMechanism,
		"sasl.username":     cfg.Username,
		"sasl.password":     cfg.Password,
	})
	defer p.Close()

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &cfg.TopicDLQ, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)
}
