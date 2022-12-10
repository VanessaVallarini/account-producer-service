package main

import (
	"account-producer-service/api"
	"account-producer-service/cmd/account-producer-service/health"
	"account-producer-service/cmd/account-producer-service/server"
	"account-producer-service/internal/config"
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/clients"
	"account-producer-service/internal/pkg/kafka"
	"account-producer-service/internal/pkg/services"
	"account-producer-service/internal/pkg/utils"

	"github.com/labstack/echo"
)

func main() {

	config := config.NewConfig()

	kafkaClient, err := kafka.NewKafkaClient(config.Kafka)
	if err != nil {
		utils.Logger.Warn("error during create kafka client")
	}

	kafkaProducer, err := kafkaClient.NewProducer()
	if err != nil {
		utils.Logger.Warn("error during kafka producer")
	}

	viaCepApiClient, err := clients.NewViaCepApiClient(config.ViaCep)
	if err != nil {
		utils.Logger.Warn("error during kafka producer")
	}

	accountServiceProducer := services.NewAccountServiceProducer(*kafkaProducer, *viaCepApiClient)

	go func() {
		setupHttpServer(accountServiceProducer, config)
	}()

	utils.Logger.Info("start application")

	health.NewHealthServer()
}

func setupHttpServer(asp *services.AccountServiceProducer, config *models.Config) *echo.Echo {

	accountApi := api.NewAccountApi(asp)
	s := server.NewServer()
	accountApi.Register(s.Server)

	s.Start(config)

	return s.Server
}
