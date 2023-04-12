package api

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

func (api *AccountApi) deleteAccount(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()

	validate := validator.New()

	var deleteAccountRequest models.AccountRequestByEmail

	email := echoContext.Param("email")
	deleteAccountRequest.Email = email

	err := echoContext.Bind(&deleteAccountRequest)
	if err != nil {
		utils.Logger.Error("account producer deleteaccount error on binding: %v", err)
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr, "deleteAccount", api.metrics)
	}

	err = validate.Struct(&deleteAccountRequest)
	if err != nil {
		utils.Logger.Error("account producer deleteaccount error on validate struct: %v", err)
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr, "deleteAccount", api.metrics)
	}

	err = api.service.Delete(ctx, deleteAccountRequest)
	if err != nil {
		errorxErr := errorx.RejectedOperation.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr, "deleteAccount", api.metrics)
	}

	return utils.BuildSuccessResponse(echoContext, http.StatusOK, "deleteAccount", api.metrics, nil, nil)
}
