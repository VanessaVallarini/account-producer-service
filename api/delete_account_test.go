package api

import (
	"account-producer-service/internal/pkg/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteAccountReturnError(t *testing.T) {
	t.Run("Expect to return 422 when email is missing", func(t *testing.T) {
		mockAccountService := mocks.NewIAccountService(t)
		mockApi := &AccountApi{
			service: mockAccountService,
		}
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/accounts/:email")
		c.SetParamNames("email")
		c.SetParamValues("")

		mockApi.deleteAccount(c)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("Expect to return 5xx when service returns error", func(t *testing.T) {
		mockAccountService := mocks.NewIAccountService(t)
		mockApi := &AccountApi{
			service: mockAccountService,
		}
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/accounts/:email")
		c.SetParamNames("email")
		c.SetParamValues("teste@email.com")

		mockAccountService.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("models.AccountRequestByEmail")).
			Return(
				errors.New("something went wrong"),
			)

		mockApi.deleteAccount(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestDeleteAccountReturnSuccess(t *testing.T) {
	t.Run("Expect to return 200", func(t *testing.T) {
		mockAccountService := mocks.NewIAccountService(t)
		mockApi := &AccountApi{
			service: mockAccountService,
		}
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/accounts/:email")
		c.SetParamNames("email")
		c.SetParamValues("teste@email.com")

		mockAccountService.On("Delete",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("models.AccountRequestByEmail")).
			Return(
				nil,
			)

		mockApi.deleteAccount(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
