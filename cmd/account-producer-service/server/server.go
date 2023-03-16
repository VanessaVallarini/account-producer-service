package server

import (
	"account-producer-service/cmd/middleware"
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"context"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

type server struct {
	Server *echo.Echo
}

func NewServer() *server {
	m := middleware.NewMetrics()
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	p := prometheus.NewPrometheus("echo", nil, m.MetricList())
	p.Use(e)
	e.Use(m.AddCustomMetricsMiddleware)
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//Skipper: func(c echo.Context) bool {
	//requestUri := c.Request().RequestURI
	//return requestUri == "/metrics"
	//},
	//}))

	return &server{
		Server: e,
	}
}

func (s *server) Start(c *models.Config) {
	utils.Logger.Info("starting server in port " + c.ServerHost)
	err := s.Server.Start(c.ServerHost)

	if err != nil {
		utils.Logger.Fatal(context.Background(), err, "unable to start server")
		panic(s.Server.Start(c.ServerHost))
	}
}
