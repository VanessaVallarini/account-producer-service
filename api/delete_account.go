package api

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo"
)

func (api *AccountApi) deleteAccount(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	validate := validator.New()

	var deleteAccountRequest models.AccountDeleteRequest

	id := echoContext.Param("id")
	deleteAccountRequest.Id = id

	err := validate.Struct(&deleteAccountRequest)
	if err != nil {
		errorxErr := errorx.IllegalArgument.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr)
	}

	err = api.service.Delete(ctx, deleteAccountRequest)
	if err != nil {
		errorxErr := errorx.RejectedOperation.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr)
	}

	return echoContext.NoContent(http.StatusCreated)
}
