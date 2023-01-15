package api

import (
	"account-producer-service/internal/pkg/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	malformedCreateAccountJSON = `{
		"email": "",
		"full_number": "5591999194410",
		"name": "Lorem",
		"zip_code": "01001-000",
		"status": "ACTIVE"
	}`
	successCreateAccountJSON = `{
		"email": "teste@email.com",
		"full_number": "5591999194410",
		"name": "Lorem",
		"zip_code": "01001-000",
		"status": "ACTIVE"
	}`
)

func TestCreateAccountReturnError(t *testing.T) {
	t.Run("Expect to return 422 when email is missing", func(t *testing.T) {
		mockAccountService := mocks.NewIAccountService(t)
		mockApi := &AccountApi{
			service: mockAccountService,
		}
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/accounts", strings.NewReader(malformedCreateAccountJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockApi.createAccount(c)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("Expect to return 5xx when service returns error", func(t *testing.T) {
		mockAccountService := mocks.NewIAccountService(t)
		mockApi := &AccountApi{
			service: mockAccountService,
		}
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/accounts", strings.NewReader(successCreateAccountJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAccountService.On("CreateOrUpdateAccount",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("models.AccountCreateRequest")).
			Return(
				errors.New("something went wrong"),
			)

		mockApi.createAccount(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestCreateAccountReturnSuccess(t *testing.T) {
	t.Run("Expect to return 201", func(t *testing.T) {
		mockAccountService := mocks.NewIAccountService(t)
		mockApi := &AccountApi{
			service: mockAccountService,
		}
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/accounts", strings.NewReader(successCreateAccountJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAccountService.On("CreateOrUpdateAccount",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("models.AccountCreateRequest")).
			Return(
				nil,
			)

		mockApi.createAccount(c)

		assert.Equal(t, http.StatusCreated, rec.Code)
	})
}
