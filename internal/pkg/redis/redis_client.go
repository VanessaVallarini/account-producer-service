package redis

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var (
	onceConfigs  sync.Once
	url          string
	db           int
	read_timeout int
	redisClient  *RedisClient
)

type RedisClientInterface interface {
	GetString(key string) (string, error)
	Setex(key string, value string, expiration int) (string, error)
}

type RedisClient struct {
	wrapper RedisWrapper
}
type RedisWrapper struct {
	client *redis.Client
}

func NewRedisClient(cfg *models.RedisConfig) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:        cfg.Url,
		DB:          cfg.Db,
		ReadTimeout: time.Duration(cfg.ReadTimeout) * time.Millisecond,
	})

	redisClient = &RedisClient{
		wrapper: RedisWrapper{
			client: client,
		},
	}

	return redisClient
}

func (client *RedisClient) Ping() error {
	_, err := client.wrapper.client.Ping().Result()
	if err != nil {
		utils.Logger.Errorf("error trying to Ping redis")
		return err
	}
	utils.Logger.Debugf("Redis PONG!")

	return nil
}

func (wrapper *RedisWrapper) Get(key string) (string, error) {
	return wrapper.client.Get(key).Result()
}

func (wrapper *RedisWrapper) setex(key string, value string, expiration int) (string, error) {
	return wrapper.client.Set(key, value, time.Duration(expiration)*time.Second).Result()
}

func (client *RedisClient) GetBool(key string) (bool, error) {
	value, err := client.wrapper.Get(key)
	if err != nil {
		utils.Logger.Errorf("Failed to fetch value, key: %s, err: %v", key, err)
		return false, err
	}
	valueAsBool, err := strconv.ParseBool(value)
	if err != nil {
		utils.Logger.Warnf("Failed to convert value to boolean, key: %s, value: %v, err: %v", key, value, err)
		return false, err
	}
	return valueAsBool, nil
}

func (client *RedisClient) GetString(key string) (string, error) {
	if key == "" {
		return "", nil
	}
	value, err := client.wrapper.Get(key)
	if err != nil {
		utils.Logger.Warnf("Failed to fetch value, key: %s, err: %v", key, err)
		return "", err
	}

	return value, nil
}

func (client *RedisClient) Setex(key string, value string, expiration int) (string, error) {
	return client.wrapper.setex(key, value, expiration)
}
