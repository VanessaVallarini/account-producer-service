package api

import (
	"account-producer-service/cmd/middleware"
	"account-producer-service/internal/pkg/utils"
	"net/http"

	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

func (api *AccountApi) getAllAccounts(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()

	metrics, ok := echoContext.Get(middleware.CKeyMetrics).(*middleware.Metrics)
	if !ok {
		utils.Logger.Error("account producer getallaccounts middleware is nil")
	}

	accounts, err := api.service.GetAll(ctx)
	if err != nil {
		errorxErr := errorx.RejectedOperation.New(err.Error())
		return utils.BuildErrorResponse(echoContext, errorxErr, "getAllAccounts", metrics)
	}

	if cap(accounts) == 0 && err == nil {
		return echoContext.JSON(http.StatusNotFound, "Account does not exist")
	}

	return echoContext.JSON(http.StatusOK, accounts)
}
