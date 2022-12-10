package config

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func NewConfig() *models.Config {
	viperConfig := initConfig()

	return &models.Config{
		AppName:          viperConfig.GetString("APP_NAME"),
		ServerHost:       viperConfig.GetString("SERVER_HOST"),
		HealthServerHost: viperConfig.GetString("HEALTH_SERVER_HOST"),
		Kafka:            buildKafkaClientConfig(viperConfig),
		ViaCep:           buildViaCepClientConfig(viperConfig),
	}
}

func initConfig() *viper.Viper {
	config := viper.New()

	config.SetConfigType("yml")
	config.SetConfigName("configuration")
	config.AddConfigPath("internal/config/")

	err := config.ReadInConfig()
	if err != nil {
		utils.Logger.Fatal("failed to read config file", err)
	}

	config.AutomaticEnv()

	return config
}

func buildKafkaClientConfig(config *viper.Viper) *models.KafkaConfig {
	return &models.KafkaConfig{
		ClientId:               config.GetString("KAFKA_CLIENT_ID"),
		Hosts:                  cast.ToStringSlice(config.GetString("KAFKA_HOSTS")),
		SchemaRegistryHost:     config.GetString("KAFKA_SCHEMA_REGISTRY_HOST"),
		Acks:                   config.GetString("KAFKA_ACKS"),
		Timeout:                config.GetInt("KAFKA_TIMEOUT"),
		UseAuthentication:      config.GetBool("KAFKA_USE_AUTEHNTICATION"),
		EnableTLS:              config.GetBool("KAFKA_ENABLE_TLS"),
		SaslMechanism:          config.GetString("KAFKA_SASL_MECHANISM"),
		User:                   config.GetString("KAFKA_USER"),
		Password:               config.GetString("KAFKA_PASSWORD"),
		SchemaRegistryUser:     config.GetString("KAFKA_SCHEMA_REGISTRY_USER"),
		SchemaRegistryPassword: config.GetString("KAFKA_SCHEMA_REGISTRY_PASSWORD"),
		EnableEvents:           config.GetBool("KAFKA_ENABLE_EVENTS"),
		MaxMessageBytes:        config.GetInt("KAFKA_MAX_MESSAGE_BYTES"),
		RetryMax:               config.GetInt("KAFKA_RETRY_MAX"),
		DlqTopic:               config.GetString("KAFKA_DLQ_TOPIC"),
		ConsumerTopic:          config.GetString("KAFKA_CONSUMER_TOPIC"),
		ConsumerGroup:          config.GetString("KAFKA_CONSUMER_GROUP"),
	}
}

func buildViaCepClientConfig(config *viper.Viper) *models.ViaCepConfig {
	return &models.ViaCepConfig{
		Url:                   config.GetString("VIA_CEP_URL"),
		MaxRetriesHttpRequest: config.GetInt("VIA_CEP_MAX_RETRIES_HTTP_REQUEST"),
		MaxFailureRatio:       config.GetFloat64("VIA_CEP_MAX_FAILURE_RATIO"),
		Name:                  config.GetString("VIA_CEP_NAME"),
	}
}
