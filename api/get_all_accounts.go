package api

import (
	"account-producer-service/internal/pkg/utils"
	"net/http"

	"github.com/joomcode/errorx"
	"github.com/labstack/echo"
)

func (api *AccountApi) getAllAccounts(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()

	accounts, err := api.service.GetAll(ctx)
	if err != nil {
		errorxErr := errorx.RejectedOperation.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr)
	}

	return echoContext.JSON(http.StatusOK, accounts)
}
