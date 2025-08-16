package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BootstrapServers string
	Username         string
	Password         string
	SaslMechanism    string
	SecurityProtocol string
	GroupID          string
	Topic            string
	TopicDLQ         string
	TargetServiceURL string
}

func Load() Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	viper.AutomaticEnv()

	return Config{
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
}
