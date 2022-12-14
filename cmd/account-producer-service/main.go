package main

import (
	"account-producer-service/api"
	"account-producer-service/cmd/account-producer-service/health"
	"account-producer-service/cmd/account-producer-service/server"
	"account-producer-service/internal/config"
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/clients"
	"account-producer-service/internal/pkg/db"
	"account-producer-service/internal/pkg/kafka"
	"account-producer-service/internal/pkg/repository"
	"account-producer-service/internal/pkg/services"
	"account-producer-service/internal/pkg/utils"

	"github.com/labstack/echo"
)

func main() {
	config := config.NewConfig()

	scylla := db.NewScylla(config.Database)
	defer scylla.Close()

	kafkaClient := kafka.NewKafkaClient(config.Kafka)

	kafkaProducer := kafkaClient.NewProducer()

	viaCepApiClient := clients.NewViaCepApiClient(config.ViaCep)

	accountRepository := repository.NewAccountRepository(scylla)
	accountServiceProducer := services.NewAccountService(kafkaProducer, viaCepApiClient, accountRepository)

	go func() {
		setupHttpServer(accountServiceProducer, config)
	}()

	utils.Logger.Info("start application")

	health.NewHealthServer()
}

func setupHttpServer(asp *services.AccountService, config *models.Config) *echo.Echo {

	accountApi := api.NewAccountApi(asp)
	s := server.NewServer()
	accountApi.Register(s.Server)

	s.Start(config)

	return s.Server
}
