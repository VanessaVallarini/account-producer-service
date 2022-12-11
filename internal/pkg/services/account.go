package services

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/clients"
	"account-producer-service/internal/pkg/kafka"
	"account-producer-service/internal/pkg/utils"
	"context"
)

const (
	topic_create = "account_create"
	topic_update = "account_update"
)

type IAccountService interface {
	Create(ctx context.Context, ae models.AccountCreateRequest) error
}

type AccountService struct {
	producer kafka.IProducer
	viaCep   clients.ViaCepApiClient
}

func NewAccountService(p kafka.IProducer, v clients.ViaCepApiClient) *AccountService {
	return &AccountService{
		producer: p,
		viaCep:   v,
	}
}

func (service *AccountService) Create(ctx context.Context, ae models.AccountCreateRequest) error {

	viaCepRequest := models.ViaCepRequest{
		Cep: ae.ZipCode,
	}

	viaCepResponse, err := service.viaCep.CallViaCepApi(ctx, viaCepRequest)
	if err != nil {
		utils.Logger.Error("error during call via cep api", err)
		return err
	}

	aCreate := models.AccountCreateEvent{
		Alias:       viaCepResponse.Uf,
		City:        viaCepResponse.Localidade,
		District:    viaCepResponse.Bairro,
		Email:       ae.Email,
		FullNumber:  ae.FullNumber,
		Name:        ae.Name,
		PublicPlace: viaCepResponse.Logradouro,
		ZipCode:     ae.ZipCode,
	}

	service.producer.Send(aCreate, topic_create, models.AccountCreateSubject)
	return nil
}

func (asp *AccountService) Update(ctx context.Context, ae models.AccountUpdateRequest) error {

	viaCepRequest := models.ViaCepRequest{
		Cep: ae.ZipCode,
	}

	viaCepResponse, err := asp.viaCep.CallViaCepApi(ctx, viaCepRequest)
	if err != nil {
		utils.Logger.Error("error during call via cep api", err)
		return err
	}

	aUpdate := models.AccountUpdateEvent{
		Id:          ae.Id,
		Alias:       viaCepResponse.Uf,
		City:        viaCepResponse.Localidade,
		District:    viaCepResponse.Bairro,
		Email:       ae.Email,
		FullNumber:  ae.FullNumber,
		Name:        ae.Name,
		PublicPlace: viaCepResponse.Logradouro,
		ZipCode:     ae.ZipCode,
	}

	asp.producer.Send(aUpdate, topic_update, models.AccountUpdateSubject)
	return nil
}
