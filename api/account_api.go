package api

import (
	"account-producer-service/internal/pkg/services"

	"github.com/labstack/echo"
)

type AccountApi struct {
	service services.IAccountService
}

func NewAccountApi(service services.IAccountService) *AccountApi {
	return &AccountApi{
		service: service,
	}
}

func (c *AccountApi) Register(router *echo.Echo) {
	v1 := router.Group("/v1")
	v1.POST("/accounts", c.createAccount)
	v1.DELETE("/accounts/:email", c.deleteAccount)
	v1.GET("/accounts/:email", c.getAccount)
	v1.GET("/accounts", c.getAllAccounts)
}
