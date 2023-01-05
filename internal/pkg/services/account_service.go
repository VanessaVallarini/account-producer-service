package services

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/clients"
	"account-producer-service/internal/pkg/kafka"
	"account-producer-service/internal/pkg/utils"
	"context"
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
}

func NewAccountService(kafkaProducer kafka.IProducer, viaCep clients.ViaCepApiClient) *AccountService {
	return &AccountService{
		producer: kafkaProducer,
		viaCep:   viaCep,
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
	service.producer.Send(request, topic_account_delete, models.AccountDeleteSubject)

	return nil
}

func (service *AccountService) GetByEmail(ctx context.Context, request models.AccountGetRequest) error {

	service.producer.Send(request, topic_account_get, models.AccountGetSubject)
	return nil
}
