package utils

import (
	"account-producer-service/internal/metrics"
	"account-producer-service/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func BuildSuccessResponse(context echo.Context, statusCode int, method string, metrics *metrics.Metrics, account *models.Account, accounts []models.Account) error {

	statusCodeStr := strconv.Itoa(statusCode)
	path := context.Path()
	mv := []string{statusCodeStr, method, path}
	metrics.ApiStrategySuccessCounter.WithLabelValues(mv...).Inc()

	if account != nil {
		return context.JSON(statusCode, account)
	}

	if accounts != nil {
		return context.JSON(statusCode, accounts)
	}

	return context.NoContent(statusCode)
}
