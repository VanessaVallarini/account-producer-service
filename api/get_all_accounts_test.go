package api

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllAccountsReturnError(t *testing.T) {
	t.Run("Expect to return 5xx when service returns error", func(t *testing.T) {
		mockAccountService := mocks.NewIAccountService(t)
		mockApi := &AccountApi{
			service: mockAccountService,
		}
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/accounts", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAccountService.On("GetAll",
			mock.AnythingOfType("*context.emptyCtx")).
			Return(
				nil, errors.New("something went wrong"),
			)

		mockApi.getAllAccounts(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("Expect to return 404 when account does not exists", func(t *testing.T) {
		mockAccountService := mocks.NewIAccountService(t)
		mockApi := &AccountApi{
			service: mockAccountService,
		}
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/accounts", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAccountService.On("GetAll",
			mock.AnythingOfType("*context.emptyCtx")).
			Return(
				nil, nil,
			)

		mockApi.getAllAccounts(c)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestGetAllAccountsReturnSuccess(t *testing.T) {
	t.Run("Expect to return account", func(t *testing.T) {
		mockAccountService := mocks.NewIAccountService(t)
		mockApi := &AccountApi{
			service: mockAccountService,
		}
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/accounts", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

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

		mockAccountService.On("GetAll",
			mock.AnythingOfType("*context.emptyCtx")).
			Return(
				accountList, nil,
			)

		mockApi.getAllAccounts(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
