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
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	scylla, err := db.NewScylla(config.Database)
	if err != nil {
		panic(err)
	}
	defer scylla.Close()

	kafkaClient, err := kafka.NewKafkaClient(
		config.Kafka,
	)
	if err != nil {
		panic(err)
	}
	defer kafkaClient.Close()

	kafkaProducer, err := kafkaClient.NewProducer(config.Kafka)
	if err != nil {
		panic(err)
	}

	server := server.NewServer()

	viaCepApiClient := clients.NewViaCepApiClient(config.ViaCep)
	accountRepository := repository.NewAccountRepository(scylla)
	accountService := services.NewAccountService(kafkaProducer, viaCepApiClient, accountRepository)
	accountApi := api.NewAccountApi(accountService)
	accountApi.Register(server.Server)

	setupHttpServer(server, config)

	utils.Logger.Info("start application")

	health.NewHealthServer()
}

func setupHttpServer(server *server.Server, config *models.Config) {
	go func() {
		server.Start(config)
	}()
}
