package kafka

import (
	"account-producer-service/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKafkaClient(t *testing.T) {
	configKafka := &models.KafkaConfig{
		ClientId:               "account-producer-service",
		Hosts:                  []string{"localhost:9092"},
		UseAuthentication:      false,
		SaslMechanism:          "SCRAM-SHA-512",
		EnableTLS:              true,
		SchemaRegistryHost:     "http://localhost:8086",
		User:                   "",
		Password:               "",
		SchemaRegistryUser:     "",
		SchemaRegistryPassword: "",
	}

	kafkaClient, _ := NewKafkaClient(configKafka)

	assert.NotNil(t, kafkaClient)
}
