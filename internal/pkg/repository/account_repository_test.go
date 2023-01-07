package repository

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/db/mocks"
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewFormRepository(t *testing.T) {
	scyllaMock := mocks.NewIScylla(t)
	accountRepository := NewAccountRepository(scyllaMock)

	assert.NotNil(t, accountRepository)
}

func TestGetByEmail(t *testing.T) {
	t.Run("Expect to return success on get account by email", func(t *testing.T) {
		ctx := context.Background()
		scyllaMock := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(scyllaMock)

		request := models.AccountRequestByEmail{}

		scyllaMock.On("ScanMap",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			nil,
		)

		response, err := accountRepository.GetByEmail(ctx, request)

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})

	t.Run("Expect to return success during query on get account by email and account not exist", func(t *testing.T) {
		ctx := context.Background()
		scyllaMock := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(scyllaMock)

		request := models.AccountRequestByEmail{}

		scyllaMock.On("ScanMap",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			errors.New("not found"),
		)

		response, err := accountRepository.GetByEmail(ctx, request)

		assert.Nil(t, err)
		assert.Nil(t, response)
	})

	t.Run("Expect to return error during query on get account by email", func(t *testing.T) {
		ctx := context.Background()
		scyllaMock := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(scyllaMock)

		request := models.AccountRequestByEmail{}

		scyllaMock.On("ScanMap",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			errors.New("error during query get account by email"),
		)

		response, err := accountRepository.GetByEmail(ctx, request)

		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestGetAllAccount(t *testing.T) {
	t.Run("Expect to return success on get all account", func(t *testing.T) {
		ctx := context.Background()
		scyllaMock := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(scyllaMock)

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

		var requestAsMap []map[string]interface{}
		marshalledRequest, _ := json.Marshal(accountList)
		json.Unmarshal(marshalledRequest, &requestAsMap)

		scyllaMock.On("ScanMapSlice",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			requestAsMap,
			nil,
		)

		response, err := accountRepository.List(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})

	t.Run("Expect to return error during query on get all account", func(t *testing.T) {
		ctx := context.Background()
		scyllaMock := mocks.NewIScylla(t)
		accountRepository := NewAccountRepository(scyllaMock)

		scyllaMock.On("ScanMapSlice",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
		).Return(
			nil,
			errors.New("error during query get all account"),
		)

		response, err := accountRepository.List(ctx)

		assert.Error(t, err)
		assert.Nil(t, response)
	})
}
