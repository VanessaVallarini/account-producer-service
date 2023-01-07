package services

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/clients"
	"account-producer-service/internal/pkg/kafka"
	"account-producer-service/internal/pkg/redis"
	"account-producer-service/internal/pkg/repository"
	"account-producer-service/internal/pkg/utils"
	"context"
	"encoding/json"

	"github.com/VanessaVallarini/account-toolkit/avros"
)

const (
	topic_account_createorupdate_dlq = "account_createorupdate_dlq"
	topic_account_delete_dlq         = "account_delete_dlq"
	topic_account_createorupdate     = "account_createorupdate"
	topic_account_delete             = "account_delete"
)

type IAccountService interface {
	CreateOrUpdateAccount(ctx context.Context, ae models.AccountCreateRequest) error
	Delete(ctx context.Context, ae models.AccountRequestByEmail) error
	GetByEmail(ctx context.Context, ade models.AccountRequestByEmail) (*models.Account, error)
	GetAll(ctx context.Context) ([]models.Account, error)
}

type AccountService struct {
	producer   kafka.IProducer
	viaCep     clients.ViaCepApiClient
	redis      redis.RedisClientInterface
	repository repository.IAccountRepository
}

func NewAccountService(kafkaProducer kafka.IProducer, redisClient redis.RedisClientInterface, viaCep clients.ViaCepApiClient, repo repository.IAccountRepository) *AccountService {
	return &AccountService{
		producer:   kafkaProducer,
		viaCep:     viaCep,
		redis:      redisClient,
		repository: repo,
	}
}

func (service *AccountService) CreateOrUpdateAccount(ctx context.Context, request models.AccountCreateRequest) error {

	viaCepRequest := models.ViaCepRequest{
		Cep: request.ZipCode,
	}

	viaCepResponse, err := service.viaCep.CallViaCepApi(ctx, viaCepRequest)
	if err != nil {
		utils.Logger.Errorf("error during call via cep api", err)
		return err
	}

	aCreate := avros.AccountCreateOrUpdateEvent{
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

	service.producer.Send(aCreate, topic_account_createorupdate, avros.AccountCreateOrUpdateSubject)
	return nil
}

func (service *AccountService) Delete(ctx context.Context, request models.AccountRequestByEmail) error {
	aDelete := avros.AccountDeleteEvent{
		Email: request.Email,
	}

	service.producer.Send(aDelete, topic_account_delete, avros.AccountDeleteSubject)

	return nil
}

func (service *AccountService) GetByEmail(ctx context.Context, request models.AccountRequestByEmail) (*models.Account, error) {
	var account *models.Account

	//first query on redis
	jsonAccount, err := service.redis.GetString(request.Email)
	if err == nil && jsonAccount != "" {
		if jsonErr := json.Unmarshal([]byte(jsonAccount), &account); jsonErr != nil {
			utils.Logger.Errorf("error unmarshal context json from Redis: %v", jsonErr)
			return nil, jsonErr
		}
		account.Email = account.Email
		return account, nil
	}

	//If not found in redis, query the database and store it in redis
	account, err = service.repository.GetByEmail(ctx, request)
	if err != nil {
		utils.Logger.Errorf("error during get account", err)
		return nil, err
	}

	accountByte, marshallErr := json.Marshal(account)
	if marshallErr != nil {
		utils.Logger.Errorf("error during encode account for save in redis", err)
	}

	_, err = service.redis.Setex(account.Email, string(accountByte), 600)
	if err != nil {
		utils.Logger.Errorf("error during save account in redis", err)
	}

	return account, nil
}

func (service *AccountService) GetAll(ctx context.Context) ([]models.Account, error) {

	accounts, err := service.repository.List(ctx)
	if err != nil {
		utils.Logger.Errorf("error during get all accounts", err)
		return nil, err
	}

	return accounts, nil
}
