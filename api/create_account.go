package api

import (
	"account-producer-service/cmd/middleware"
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

func (api *AccountApi) createAccount(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()

	metrics := echoContext.Get(middleware.CKeyMetrics).(*middleware.Metrics)

	validate := validator.New()

	var createAccountRequest models.AccountCreateRequest

	err := echoContext.Bind(&createAccountRequest)
	if err != nil {
		utils.Logger.Errorf("error on binding info: %v", err)
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr)
	}

	err = validate.Struct(&createAccountRequest)
	if err != nil {
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr)
	}

	err = api.service.CreateOrUpdateAccount(ctx, createAccountRequest)
	if err != nil {
		errorxErr := errorx.RejectedOperation.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr)
	}

	metrics.IncCustomCnt("any", "value")
	return echoContext.NoContent(http.StatusCreated)
}
