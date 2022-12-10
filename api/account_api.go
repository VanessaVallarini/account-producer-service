package api

import (
	"account-producer-service/internal/pkg/services"

	"github.com/labstack/echo"
)

type AccountApi struct {
	service *services.AccountServiceProducer
}

func NewAccountApi(service *services.AccountServiceProducer) *AccountApi {
	return &AccountApi{
		service: service,
	}
}

func (c *AccountApi) Register(router *echo.Echo) {
	v1 := router.Group("/v1")
	v1.POST("/accounts", c.createAccount)
}
