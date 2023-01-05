package kafka

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/redis"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

const (
	topic_account_createorupdate_dlq = "account_createorupdate_dlq"
	topic_account_delete_dlq         = "account_delete_dlq"
	topic_account_get_dlq            = "account_get_dlq"
	topic_account_get_response_dlq   = "account_get_response_dlq"
	topic_account_createorupdate     = "account_createorupdate"
	topic_account_delete             = "account_delete"
	topic_account_get                = "account_get"
	topic_account_get_response       = "account_get_response"
)

type Consumer struct {
	ready         chan bool
	dlqTopic      string
	consumerTopic []string
	sr            *SchemaRegistry
	producer      sarama.SyncProducer
	redis         *redis.RedisClient
}

func NewConsumer(ctx context.Context, cfg *models.KafkaConfig, kafkaClient *KafkaClient, redisClient *redis.RedisClient) error {
	consumer := Consumer{
		sr:            kafkaClient.SchemaRegistry,
		ready:         make(chan bool),
		dlqTopic:      cfg.DlqTopic,
		consumerTopic: cfg.ConsumerTopic,
		redis:         redisClient,
	}

	wg := &sync.WaitGroup{} //tratar erros em go rotinas de forma concorrente
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			ctx := context.Background()
			propagators := propagation.TraceContext{}
			handler := otelsarama.WrapConsumerGroupHandler(&consumer, otelsarama.WithPropagators(propagators))
			if err := kafkaClient.GroupClient.Consume(ctx, cfg.ConsumerTopic, handler); err != nil {
				zap.S().Errorf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				err := kafkaClient.GroupClient.Close()
				if err != nil {
					zap.S().Fatalf("Error from consumer: %v", err)
				}

				zap.S().Info("consume closed, consuming again")
			}

			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready
	zap.S().Info("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigterm:
		log.Println("terminating: via signal")
	}

	wg.Wait()
	if err := kafkaClient.GroupClient.Close(); err != nil {
		zap.S().Panicf("Error closing groupClient: %v", err)
	}
	return nil
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			ctx := context.Background()
			if err := consumer.processMessage(ctx, message); err != nil {
				consumer.sendToDlq(ctx, consumer.dlqTopic, message)
			}
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func (consumer *Consumer) sendToDlq(ctx context.Context, dlqTopic string, message *sarama.ConsumerMessage) {
	if topic := consumer.getTopicDlq(ctx, message); topic != "" {
		dlqTopic = topic
	}

	ctx, span := otel.GetTracerProvider().Tracer("consumer").Start(ctx, "sendToDlq")
	defer span.End()
	msg := &sarama.ProducerMessage{
		Topic:     dlqTopic,
		Key:       sarama.ByteEncoder(message.Key),
		Value:     sarama.ByteEncoder(message.Value),
		Timestamp: time.Now(),
	}
	for _, header := range message.Headers {
		msg.Headers = append(msg.Headers, *header)
	}

	partition, offset, err := consumer.producer.SendMessage(msg)
	if err != nil {
		zap.S().Error(err)
		span.SetStatus(codes.Error, err.Error())
		// change to retry queues instead of recursive approach
		consumer.sendToDlq(ctx, dlqTopic, message)
	}
	span.SetAttributes(attribute.String("topic", dlqTopic))
	span.SetAttributes(attribute.Int("partition", int(partition)))
	span.SetAttributes(attribute.Int64("offset", offset))
	zap.S().Infof("Message sent to dlq: topic = %s, partition = %v, offset = %v", dlqTopic, partition, offset)
}

func (consumer *Consumer) getTopicDlq(ctx context.Context, message *sarama.ConsumerMessage) string {
	switch message.Topic {
	case topic_account_createorupdate:
		return topic_account_createorupdate_dlq
	case topic_account_delete:
		return topic_account_delete_dlq
	case topic_account_get:
		return topic_account_get_dlq
	}

	return ""
}
