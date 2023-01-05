package listner

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/kafka"
	"account-producer-service/internal/pkg/redis"
	"account-producer-service/internal/pkg/utils"
	"context"
)

func Start(ctx context.Context, cfg *models.KafkaConfig, kafkaClient *kafka.KafkaClient, redisClient *redis.RedisClient) {
	err := kafka.NewConsumer(ctx, cfg, kafkaClient, redisClient)
	if err != nil {
		utils.Logger.Error("Error consumer msg: %v", err)
	}
}
