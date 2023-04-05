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

	metrics, ok := echoContext.Get(middleware.CKeyMetrics).(*middleware.Metrics)
	if !ok {
		utils.Logger.Error("account producer createaccount middleware is nil")
	}

	validate := validator.New()

	var createAccountRequest models.AccountCreateRequest

	err := echoContext.Bind(&createAccountRequest)
	if err != nil {
		utils.Logger.Error("account producer createaccount error on binding: %v", err)
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr, "createAccount", metrics)
	}

	err = validate.Struct(&createAccountRequest)
	if err != nil {
		utils.Logger.Error("account producer createaccount error on validate struct: %v", err)
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr, "createAccount", metrics)
	}

	err = api.service.CreateOrUpdateAccount(ctx, createAccountRequest)
	if err != nil {
		errorxErr := errorx.RejectedOperation.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr, "createAccount", metrics)
	}

	return echoContext.NoContent(http.StatusCreated)
}
