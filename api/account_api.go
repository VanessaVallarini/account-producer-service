package api

import (
	"account-producer-service/internal/metrics"
	"account-producer-service/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

type AccountApi struct {
	service services.IAccountService
	metrics *metrics.Metrics
}

func NewAccountApi(service services.IAccountService, metrics *metrics.Metrics) *AccountApi {
	return &AccountApi{
		service: service,
		metrics: metrics,
	}
}

func (c *AccountApi) Register(router *echo.Echo) {
	v1 := router.Group("/v1")
	v1.POST("/accounts", c.createAccount)
	v1.DELETE("/accounts/:email", c.deleteAccount)
	v1.GET("/accounts/:email", c.getAccount)
	v1.GET("/accounts", c.getAllAccounts)
}
