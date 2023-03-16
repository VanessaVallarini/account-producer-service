package api

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

func (api *AccountApi) getAccount(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	validate := validator.New()

	var getAccountRequest models.AccountRequestByEmail

	email := echoContext.Param("email")
	getAccountRequest.Email = email

	err := echoContext.Bind(&getAccountRequest)
	if err != nil {
		utils.Logger.Errorf("error on binding info: %v", err)
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr)
	}

	err = validate.Struct(&getAccountRequest)
	if err != nil {
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr)
	}

	account, err := api.service.GetByEmail(ctx, getAccountRequest)
	if err != nil {
		errorxErr := errorx.RejectedOperation.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr)
	}

	if account == nil && err == nil {
		return echoContext.JSON(http.StatusNotFound, "Account does not exist")
	}

	return echoContext.JSON(http.StatusOK, account)
}
