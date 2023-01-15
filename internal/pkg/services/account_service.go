package services

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/clients"
	"account-producer-service/internal/pkg/kafka"
	"account-producer-service/internal/pkg/repository"
	"account-producer-service/internal/pkg/utils"
	"context"
	"strings"

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
	producer   kafka.IKafkaProducer
	viaCep     clients.IViaCepApiClient
	repository repository.IAccountRepository
}

func NewAccountService(kafkaProducer kafka.IKafkaProducer, viaCep clients.IViaCepApiClient, repo repository.IAccountRepository) *AccountService {
	return &AccountService{
		producer:   kafkaProducer,
		viaCep:     viaCep,
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
		Status:      models.AccountStatusString(request.Status).String(),
		ZipCode:     request.ZipCode,
	}

	err = service.producer.Send(aCreate, topic_account_createorupdate, avros.AccountCreateOrUpdateSubject)
	if err != nil {
		utils.Logger.Errorf("error during send msg", err)
		return err
	}

	return nil
}

func (service *AccountService) Delete(ctx context.Context, request models.AccountRequestByEmail) error {
	aDelete := avros.AccountDeleteEvent{
		Email: request.Email,
	}

	err := service.producer.Send(aDelete, topic_account_delete, avros.AccountDeleteSubject)
	if err != nil {
		utils.Logger.Errorf("error during send msg", err)
		return err
	}

	return nil
}

func (service *AccountService) GetByEmail(ctx context.Context, request models.AccountRequestByEmail) (*models.Account, error) {
	account, err := service.repository.GetByEmail(ctx, request)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		}
		utils.Logger.Errorf("error during get account by email", err)
		return nil, err
	}

	return account, nil
}

func (service *AccountService) GetAll(ctx context.Context) ([]models.Account, error) {

	accounts, err := service.repository.List(ctx)
	if err != nil {
		utils.Logger.Errorf("error during get account all acounts", err)
		return nil, err
	}

	return accounts, nil
}
