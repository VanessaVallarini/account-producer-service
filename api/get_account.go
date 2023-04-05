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

func (api *AccountApi) getAccount(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()

	metrics, ok := echoContext.Get(middleware.CKeyMetrics).(*middleware.Metrics)
	if !ok {
		utils.Logger.Error("account producer getaccount middleware is nil")
	}

	validate := validator.New()

	var getAccountRequest models.AccountRequestByEmail

	email := echoContext.Param("email")
	getAccountRequest.Email = email

	err := echoContext.Bind(&getAccountRequest)
	if err != nil {
		utils.Logger.Error("account producer getaccount error on binding: %v", err)
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr, "getAccount", metrics)
	}

	err = validate.Struct(&getAccountRequest)
	if err != nil {
		utils.Logger.Error("account producer getaccount error on validate struct: %v", err)
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr, "getAccount", metrics)
	}

	account, err := api.service.GetByEmail(ctx, getAccountRequest)
	if err != nil {
		errorxErr := errorx.RejectedOperation.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr, "getAccount", metrics)
	}

	if account == nil && err == nil {
		return echoContext.JSON(http.StatusNotFound, "Account does not exist")
	}

	return echoContext.JSON(http.StatusOK, account)
}
