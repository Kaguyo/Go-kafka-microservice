package service

import (
	"fmt"
	"go-kafka-microservice/internal/config"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func SendDLQ(cfg config.Config, value []byte) error {
	// Cria o producer e verifica erros
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"security.protocol": cfg.SecurityProtocol,
		"sasl.mechanisms":   cfg.SaslMechanism,
		"sasl.username":     cfg.Username,
		"sasl.password":     cfg.Password,
	})
	if err != nil {
		return fmt.Errorf("falha ao criar Producer Kafka em DLQ: %w", err, "\n")
	}

	defer p.Close()

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &cfg.TopicDLQ, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	return nil
}
