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

		fmt.Println("Mensagem recebida:", string(msg.Value))
		//logger

		var payload map[string]any
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			fmt.Println("Erro ao deserializar JSON. Enviando para DLQ.")
			service.SendDLQ(cfg, msg.Value)
			continue
		}

		_, hasOrdem := payload["ordemDeVenda"]
		_, hasEtapa := payload["etapaAtual"]

		if !hasOrdem || !hasEtapa {
			fmt.Println("Campos obrigatórios ausentes. Enviando para DLQ.")
			service.SendDLQ(cfg, msg.Value)
			//logger
			continue
		}

		// Tudo certo, processa
		go service.SendToTarget(cfg, payload)
		fmt.Println("Enviando para o serviço:", cfg.TargetServiceURL, "...")
	}
}
