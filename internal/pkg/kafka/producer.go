package kafka

import (
	"account-producer-service/internal/pkg/utils"
	"time"

	"github.com/Shopify/sarama"
)

type IKafkaProducer interface {
	Send(msg interface{}, topic, subject string) error
}

type KafkaProducer struct {
	syncProducer sarama.SyncProducer
	schema       *SchemaRegistry
}

func (kc *KafkaClient) NewProducer() (*KafkaProducer, error) {
	producer, err := sarama.NewSyncProducerFromClient(kc.Client)
	if err != nil {
		utils.Logger.Fatal("Error during kafka producer. Details: %v", err)
		return nil, err
	}
	return &KafkaProducer{producer, kc.SchemaRegistry}, nil
}

func (ip *KafkaProducer) Send(msg interface{}, topic, subject string) error {
	msgEncoder, err := ip.schema.Encode(msg, subject)
	if err != nil {
		utils.Logger.Error("Error encode msg to send: %v", err)
		return err
	}

	m := sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.ByteEncoder(time.Now().String()),
		Value:     sarama.ByteEncoder(msgEncoder),
		Timestamp: time.Now(),
	}
	ip.syncProducer.SendMessage(&m)

	return nil
}
