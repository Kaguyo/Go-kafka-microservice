package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	BootstrapServers string `validate:"required"`
	Username         string `validate:"required"`
	Password         string `validate:"required"`
	SaslMechanism    string `validate:"required,oneof=PLAIN SCRAM-SHA-256 SCRAM-SHA-512"`
	SecurityProtocol string `validate:"required,oneof=SASL_SSL SASL_PLAINTEXT"`
	GroupID          string `validate:"required"`
	Topic            string `validate:"required,min=3"`
	TopicDLQ         string `validate:"required,min=3"`
	TargetServiceURL string `validate:"required,url"`
}

func Load() Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("erro ao ler o arquivo .env: %w", err))
	}

	viper.AutomaticEnv()

	cfg := Config{
		BootstrapServers: viper.GetString("KAFKA_BOOTSTRAP_SERVERS"),
		Username:         viper.GetString("KAFKA_USERNAME"),
		Password:         viper.GetString("KAFKA_PASSWORD"),
		SaslMechanism:    viper.GetString("KAFKA_SASL_MECHANISM"),
		SecurityProtocol: viper.GetString("KAFKA_SECURITY_PROTOCOL"),
		GroupID:          viper.GetString("KAFKA_GROUP_ID"),
		Topic:            viper.GetString("KAFKA_TOPIC"),
		TopicDLQ:         viper.GetString("KAFKA_TOPIC_DLQ"),
		TargetServiceURL: viper.GetString("TARGET_SERVICE_URL"),
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		panic(fmt.Errorf("erro de validação no arquivo de configuração: %w", err))
	}

	return cfg
}
