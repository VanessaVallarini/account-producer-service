package services

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/clients"
	"account-producer-service/internal/pkg/kafka"
	"account-producer-service/internal/pkg/redis"
	"account-producer-service/internal/pkg/utils"
	"context"
	"encoding/json"
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

type IAccountService interface {
	CreateOrUpdateAccount(ctx context.Context, ae models.AccountCreateRequest) error
	Delete(ctx context.Context, ae models.AccountDeleteRequest) error
	GetByEmail(ctx context.Context, ade models.AccountGetRequest) (*models.Account, error)
}

type AccountService struct {
	producer kafka.IProducer
	viaCep   clients.ViaCepApiClient
	redis    redis.RedisClientInterface
}

func NewAccountService(kafkaProducer kafka.IProducer, redisClient redis.RedisClientInterface, viaCep clients.ViaCepApiClient) *AccountService {
	return &AccountService{
		producer: kafkaProducer,
		viaCep:   viaCep,
		redis:    redisClient,
	}
}

func (service *AccountService) CreateOrUpdateAccount(ctx context.Context, request models.AccountCreateRequest) error {

	viaCepRequest := models.ViaCepRequest{
		Cep: request.ZipCode,
	}

	viaCepResponse, err := service.viaCep.CallViaCepApi(ctx, viaCepRequest)
	if err != nil {
		utils.Logger.Error("error during call via cep api", err)
		return err
	}

	aCreate := models.AccountCreateOrUpdateEvent{
		Email:       request.Email,
		FullNumber:  request.FullNumber,
		Alias:       viaCepResponse.Uf,
		City:        viaCepResponse.Localidade,
		District:    viaCepResponse.Bairro,
		Name:        request.Name,
		PublicPlace: viaCepResponse.Logradouro,
		Status:      request.Status,
		ZipCode:     request.ZipCode,
	}

	service.producer.Send(aCreate, topic_account_createorupdate, models.AccountCreateOrUpdateSubject)
	return nil
}

func (service *AccountService) Delete(ctx context.Context, request models.AccountDeleteRequest) error {
	aDelete := models.AccountDeleteEvent{
		Email: request.Email,
	}

	service.producer.Send(aDelete, topic_account_delete, models.AccountDeleteSubject)

	return nil
}

func (service *AccountService) GetByEmail(ctx context.Context, request models.AccountGetRequest) (*models.Account, error) {
	aGet := models.AccountGetEvent{
		Email: request.Email,
	}

	service.producer.Send(aGet, topic_account_get, models.AccountGetSubject)

	var account *models.Account
	jsonAccount, err := service.redis.GetString(request.Email)
	if err == nil && jsonAccount != "" {
		if jsonErr := json.Unmarshal([]byte(jsonAccount), &account); jsonErr != nil {
			utils.Logger.Errorf("error unmarshal context json from Redis: %v", jsonErr)
			return nil, jsonErr
		}
		account.Email = account.Email
		return account, nil
	}

	return nil, nil
}
