package server

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"context"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type server struct {
	Server *echo.Echo
}

func NewServer() *server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.PUT},
	}))

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
