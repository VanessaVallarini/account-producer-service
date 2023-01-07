package services

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAccountService(t *testing.T) {
	mockKafkaProducer := mocks.NewIKafkaProducer(t)
	mockViaCepClient := mocks.NewIViaCepApiClient(t)
	mockAccountRepository := mocks.NewIAccountRepository(t)
	accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

	assert.NotNil(t, accountService)
}

func TestCreateOrUpdateAccount(t *testing.T) {
	t.Run("Expect to return success on create account", func(t *testing.T) {
		ctx := context.Background()
		mockKafkaProducer := mocks.NewIKafkaProducer(t)
		mockViaCepClient := mocks.NewIViaCepApiClient(t)
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

		request := models.AccountCreateRequest{
			Email:      "lorem1@email.com",
			FullNumber: "5591999194410",
			Name:       "Lorem",
			Status:     models.Active.String(),
			ZipCode:    "01001-000",
		}

		viaCepResponse := models.ViaCepResponse{
			Cep:         "01001-000",
			Logradouro:  "Praça da Sé",
			Complemento: "lado ímpar",
			Bairro:      "Sé",
			Localidade:  "São Paulo",
			Uf:          "SP",
			Ibge:        "3550308",
			Gia:         "1004",
			Ddd:         "11",
			Siafi:       "7107",
		}

		mockViaCepClient.On("CallViaCepApi",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			&viaCepResponse,
			nil,
		)

		mockKafkaProducer.On("Send",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			nil,
		)

		err := accountService.CreateOrUpdateAccount(ctx, request)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error on send msg", func(t *testing.T) {
		ctx := context.Background()
		mockKafkaProducer := mocks.NewIKafkaProducer(t)
		mockViaCepClient := mocks.NewIViaCepApiClient(t)
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

		request := models.AccountCreateRequest{
			Email:      "lorem1@email.com",
			FullNumber: "5591999194410",
			Name:       "Lorem",
			Status:     models.Active.String(),
			ZipCode:    "01001-000",
		}

		viaCepResponse := models.ViaCepResponse{
			Cep:         "01001-000",
			Logradouro:  "Praça da Sé",
			Complemento: "lado ímpar",
			Bairro:      "Sé",
			Localidade:  "São Paulo",
			Uf:          "SP",
			Ibge:        "3550308",
			Gia:         "1004",
			Ddd:         "11",
			Siafi:       "7107",
		}

		mockViaCepClient.On("CallViaCepApi",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			&viaCepResponse,
			nil,
		)

		mockKafkaProducer.On("Send",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			errors.New("error during call via cep api"),
		)

		err := accountService.CreateOrUpdateAccount(ctx, request)

		assert.Error(t, err)
	})

	t.Run("Expect to return error on call cep api", func(t *testing.T) {
		ctx := context.Background()
		mockKafkaProducer := mocks.NewIKafkaProducer(t)
		mockViaCepClient := mocks.NewIViaCepApiClient(t)
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

		request := models.AccountCreateRequest{
			Email:      "lorem1@email.com",
			FullNumber: "5591999194410",
			Name:       "Lorem",
			Status:     models.Active.String(),
			ZipCode:    "01001-000",
		}

		mockViaCepClient.On("CallViaCepApi",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			nil,
			errors.New("error during call via cep api"),
		)

		err := accountService.CreateOrUpdateAccount(ctx, request)

		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Expect to return success on delete account", func(t *testing.T) {
		ctx := context.Background()
		mockKafkaProducer := mocks.NewIKafkaProducer(t)
		mockViaCepClient := mocks.NewIViaCepApiClient(t)
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

		request := models.AccountRequestByEmail{
			Email: "lorem1@email.com",
		}

		mockKafkaProducer.On("Send",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			nil,
		)

		err := accountService.Delete(ctx, request)

		assert.Nil(t, err)
	})

	t.Run("Expect to return error on send msg", func(t *testing.T) {
		ctx := context.Background()
		mockKafkaProducer := mocks.NewIKafkaProducer(t)
		mockViaCepClient := mocks.NewIViaCepApiClient(t)
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

		request := models.AccountRequestByEmail{
			Email: "lorem1@email.com",
		}

		mockKafkaProducer.On("Send",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			errors.New("error during call via cep api"),
		)

		err := accountService.Delete(ctx, request)

		assert.Error(t, err)
	})
}

func TestGetAccountByEmail(t *testing.T) {
	t.Run("Expect to return success on get account by email", func(t *testing.T) {
		ctx := context.Background()
		mockKafkaProducer := mocks.NewIKafkaProducer(t)
		mockViaCepClient := mocks.NewIViaCepApiClient(t)
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

		request := models.AccountRequestByEmail{
			Email: "lorem1@email.com",
		}

		account := models.Account{
			Email:       "lorem1@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			DateTime:    "2023-01-07 15:59:00.715669 -0300 -03 m=+88.440179745",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}

		mockAccountRepository.On("GetByEmail",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			&account,
			nil,
		)

		response, err := accountService.GetByEmail(ctx, request)

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})

	t.Run("Expect to return account nil on get account by email", func(t *testing.T) {
		ctx := context.Background()
		mockKafkaProducer := mocks.NewIKafkaProducer(t)
		mockViaCepClient := mocks.NewIViaCepApiClient(t)
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

		request := models.AccountRequestByEmail{
			Email: "lorem1@email.com",
		}

		mockAccountRepository.On("GetByEmail",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			nil,
			nil,
		)

		response, err := accountService.GetByEmail(ctx, request)

		assert.Nil(t, err)
		assert.Nil(t, response)
	})

	t.Run("Expect to return error on get account by email", func(t *testing.T) {
		ctx := context.Background()
		mockKafkaProducer := mocks.NewIKafkaProducer(t)
		mockViaCepClient := mocks.NewIViaCepApiClient(t)
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

		request := models.AccountRequestByEmail{
			Email: "lorem1@email.com",
		}

		mockAccountRepository.On("GetByEmail",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			nil,
			errors.New("error during get account by email"),
		)

		response, err := accountService.GetByEmail(ctx, request)

		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestGetAllAccounts(t *testing.T) {
	t.Run("Expect to return success on get all accounts", func(t *testing.T) {
		ctx := context.Background()
		mockKafkaProducer := mocks.NewIKafkaProducer(t)
		mockViaCepClient := mocks.NewIViaCepApiClient(t)
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

		var accountList []models.Account
		account := models.Account{
			Email:       "lorem1@email.com",
			FullNumber:  "5591999194410",
			Alias:       "SP",
			City:        "São Paulo",
			DateTime:    "2023-01-07 15:59:00.715669 -0300 -03 m=+88.440179745",
			District:    "Sé",
			Name:        "Lorem",
			PublicPlace: "Praça da Sé",
			Status:      models.Active.String(),
			ZipCode:     "01001-000",
		}

		accountList = append(accountList, account)
		accountList = append(accountList, account)

		mockAccountRepository.On("List",
			mock.AnythingOfType("*context.emptyCtx"),
		).Return(
			accountList,
			nil,
		)

		response, err := accountService.GetAll(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})

	t.Run("Expect to return error on get all accounts", func(t *testing.T) {
		ctx := context.Background()
		mockKafkaProducer := mocks.NewIKafkaProducer(t)
		mockViaCepClient := mocks.NewIViaCepApiClient(t)
		mockAccountRepository := mocks.NewIAccountRepository(t)
		accountService := NewAccountService(mockKafkaProducer, mockViaCepClient, mockAccountRepository)

		mockAccountRepository.On("List",
			mock.AnythingOfType("*context.emptyCtx"),
		).Return(
			nil,
			errors.New("error during get all accounts"),
		)

		response, err := accountService.GetAll(ctx)

		assert.Error(t, err)
		assert.Nil(t, response)
	})
}
