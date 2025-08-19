package consumer

import (
	"encoding/json"
	"fmt"
	"go-kafka-microservice/internal/config"
	"go-kafka-microservice/internal/service"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Start(cfg config.Config) error {
	fmt.Printf("Iniciando consumidor Kafka para o tópico: %s\n", cfg.Topic)
	cm := kafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"group.id":          cfg.GroupID,
		"security.protocol": cfg.SecurityProtocol,
		"auto.offset.reset": "earliest",
	}

	if cfg.SaslMechanism != "" {
		cm["sasl.mechanisms"] = cfg.SaslMechanism
		cm["sasl.username"] = cfg.Username
		cm["sasl.password"] = cfg.Password
	}

	c, err := kafka.NewConsumer(&cm)

	if err != nil {
		fmt.Printf("Erro ao criar consumidor Kafka: %v\n", err)
		return err
	}

	c.SubscribeTopics([]string{cfg.Topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			continue
		}

		var payload map[string]any
		fmt.Print("Mensagem recebida: ", string(msg.Value), "\n")
		if json.Unmarshal(msg.Value, &payload) != nil {
			service.SendDLQ(cfg, msg.Value)
			fmt.Printf("Erro ao deserializar a mensagem, enviando para DLQ", err, "\n")
			continue
		}

		go service.SendToTarget(cfg, payload)
	}
}
