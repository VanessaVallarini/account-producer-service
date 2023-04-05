package kafka

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"time"

	"github.com/Shopify/sarama"
	"github.com/VanessaVallarini/account-toolkit/avros"
)

const (
	Avro string = "AVRO"
)

type IKafkaProducer interface {
	Send(msg interface{}, topic, subject string) error
}

type KafkaProducer struct {
	syncProducer sarama.SyncProducer
	schema       *SchemaRegistry
}

func NewProducer(cfg *models.KafkaConfig) (*KafkaProducer, error) {
	client, err := NewKafkaClient(cfg)
	if err != nil {
		return nil, err
	}

	client.config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(cfg.Hosts, client.config)
	if err != nil {
		utils.Logger.Error("kafka producer failed to new sync producer: %v", err)
		return nil, err
	}

	if err := client.schemaRegistry.ValidateSchema(avros.AccountCreateOrUpdateAvro, avros.AccountCreateOrUpdateSubject, Avro); err != nil {
		return nil, err
	}

	if err := client.schemaRegistry.ValidateSchema(avros.AccountDeleteAvro, avros.AccountDeleteSubject, Avro); err != nil {
		return nil, err
	}

	return &KafkaProducer{producer, client.schemaRegistry}, nil
}

func (ip *KafkaProducer) Send(msg interface{}, topic, subject string) error {
	msgEncoder, err := ip.schema.Encode(msg, subject)
	if err != nil {
		utils.Logger.Error("kafka producer failed encode msg to send: %v", err)
		return err
	}

	key := utils.GenerateRandomUUID()

	m := sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.ByteEncoder(key.String()),
		Value:     sarama.ByteEncoder(msgEncoder),
		Timestamp: time.Now(),
	}
	ip.syncProducer.SendMessage(&m)

	return nil
}
