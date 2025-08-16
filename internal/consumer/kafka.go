package consumer

import (
	"encoding/json"

	"../config"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Start(cfg config.Config) {
	c, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"group.id":          cfg.GroupID,
		"security.protocol": cfg.SecurityProtocol,
		"sasl.mechanisms":   cfg.SaslMechanism,
		"sasl.username":     cfg.Username,
		"sasl.password":     cfg.Password,
		"auto.offset.reset": "earliest",
	})
	c.SubscribeTopics([]string{cfg.Topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			continue
		}

		var payload map[string]any
		if json.Unmarshal(msg.Value, &payload) != nil {
			sendDLQ(cfg, msg.Value)
			continue
		}

		go sendToTarget(cfg, payload)
	}
}
