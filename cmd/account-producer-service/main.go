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
	"account-producer-service/internal/pkg/redis"
	"account-producer-service/internal/pkg/repository"
	"account-producer-service/internal/pkg/services"
	"account-producer-service/internal/pkg/utils"

	"github.com/labstack/echo"
)

func main() {
	config := config.NewConfig()

	scylla := db.NewScylla(config.Database)
	defer scylla.Close()

	redisClient := redis.NewRedisClient(config.Redis)
	if redisErr := redisClient.Ping(); redisErr != nil {
		utils.Logger.Fatal("error during create redis client", redisErr)
		panic(redisErr)
	}

	kafkaClient, err := kafka.NewKafkaClient(config.Kafka)

	kafkaProducer, err := kafkaClient.NewProducer()
	if err != nil {
		utils.Logger.Fatal("Error during kafka producer. Details: %v", err)
		panic(kafkaProducer)
	}

	viaCepApiClient, err := clients.NewViaCepApiClient(config.ViaCep)
	if err != nil {
		utils.Logger.Fatal("Error during via Cep Api Client. Details: %v", err)
		panic(viaCepApiClient)
	}

	accountRepository := repository.NewAccountRepository(scylla)
	accountServiceProducer := services.NewAccountService(*kafkaProducer, redisClient, *viaCepApiClient, accountRepository)

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
