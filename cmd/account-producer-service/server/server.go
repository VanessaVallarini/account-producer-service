package server

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

type Server struct {
	Server *echo.Echo
}

func NewServer(appName string) *Server {
	svr := echo.New()
	svr.Use(otelecho.Middleware(appName))
	svr.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			requestUri := c.Request().RequestURI
			return requestUri == "/metrics" || requestUri == "/healthcheck/liveness" || requestUri == "/healthcheck/readiness"
		},
	}))
	svr.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	return &Server{
		Server: svr,
	}
}

func (s *Server) Start(c *models.Config) {
	utils.Logger.Info("starting server in port " + c.ServerHost)
	err := s.Server.Start(c.ServerHost)

	if err != nil {
		utils.Logger.Fatal(context.Background(), err, "unable to start server")
		panic(s.Server.Start(c.ServerHost))
	}
}
