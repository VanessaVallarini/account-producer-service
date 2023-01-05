package kafka

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
)

func (consumer *Consumer) processMessage(ctx context.Context, message *sarama.ConsumerMessage) error {
	var account models.AccountGetResponseEvent

	if err := consumer.sr.Decode(message.Value, &account, models.AccountGetResponseSubject); err != nil {
		utils.Logger.Error("error during decode message consumer kafka")
		return err
	}

	accountByte, marshallErr := json.Marshal(account)
	if marshallErr != nil {
		panic(marshallErr)
	}

	_, err := consumer.redis.Setex(account.Email, string(accountByte), 600)
	if err != nil {
		panic(err)
	}

	return nil
}
