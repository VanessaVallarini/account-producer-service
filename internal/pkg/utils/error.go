package utils

import (
	"account-producer-service/internal/metrics"
	"net/http"
	"strconv"

	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

var NotFound = errorx.NewType(errorx.NewNamespace("common"), "not_found", errorx.NotFound())

type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func newError(status int, message string) *Error {
	return &Error{
		StatusCode: status,
		Message:    message,
	}
}

func handleErrorResponse(errx *errorx.Error) *Error {
	var errorResp *Error
	switch {
	case errorx.IsOfType(errx, errorx.IllegalArgument):
		errorResp = newError(http.StatusUnprocessableEntity, errx.Error())
	case errorx.IsOfType(errx, errorx.IllegalFormat):
		errorResp = newError(http.StatusBadRequest, errx.Error())
	case errorx.HasTrait(errx, errorx.NotFound()):
		errorResp = newError(http.StatusNotFound, errx.Error())
	default:
		errorResp = newError(http.StatusInternalServerError, errx.Error())
	}
	return errorResp
}

func BuildErrorResponse(context echo.Context, errx *errorx.Error, method string, metrics *metrics.Metrics) error {
	errorResponse := handleErrorResponse(errx)

	statusCodeStr := strconv.Itoa(errorResponse.StatusCode)
	path := context.Path()
	mv := []string{statusCodeStr, method, path}
	metrics.ApiStrategyErrosCounter.WithLabelValues(mv...).Inc()

	return context.JSON(errorResponse.StatusCode, errorResponse)
}
