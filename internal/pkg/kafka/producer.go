package kafka

import (
	"account-producer-service/internal/pkg/utils"
	"time"

	"github.com/Shopify/sarama"
)

type IProducer struct {
	syncProducer sarama.SyncProducer
	schema       *SchemaRegistry
}

func (kc *KafkaClient) NewProducer() (*IProducer, error) {
	producer, err := sarama.NewSyncProducerFromClient(kc.Client)
	if err != nil {
		return nil, err
	}
	return &IProducer{producer, kc.SchemaRegistry}, nil
}

func (ip *IProducer) Send(msg interface{}, topic, subject string) {
	msgEncoder, err := ip.schema.Encode(msg, subject)
	if err != nil {
		utils.Logger.Error("Error send msg: %v", err)
	}

	m := sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.ByteEncoder(time.Now().String()),
		Value:     sarama.ByteEncoder(msgEncoder),
		Timestamp: time.Now(),
	}
	ip.syncProducer.SendMessage(&m)
}
