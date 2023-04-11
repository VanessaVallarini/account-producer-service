package kafka

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"errors"
	"time"

	"github.com/Shopify/sarama"
)

const (
	_saramaTimeoutMs      = 10000
	_saramaTimeoutFlushMs = 500
)

type KafkaClient struct {
	saramaConfig         *sarama.Config
	saramaClient         sarama.Client
	saramaClusterAdmin   sarama.ClusterAdmin
	saramaSchemaRegistry *SchemaRegistry
}

func NewKafkaClient(cfg *models.KafkaConfig) (*KafkaClient, error) {
	saramaConfig, err := generateSaramaConfig(cfg)
	if err != nil {
		utils.Logger.Error("kafka client failed to generate sarama config: %v", err)
		return nil, err
	}

	saramaClient, err := sarama.NewClient(cfg.Hosts, saramaConfig)
	if err != nil {
		utils.Logger.Error("kafka client failed to new kafka client: %v", err)
		return nil, err
	}

	saramaClusterAdmin, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		utils.Logger.Error("kafka client failed to new cluster admin from client: %v", err)
		return nil, err
	}

	if err := createTopic(cfg.ProducerTopic, saramaClusterAdmin); err != nil {
		utils.Logger.Error("kafka client create producer topic %v: %v", cfg.ProducerTopic, err)
	}

	saramaSchemaRegistry, err := NewSchemaRegistry(cfg.SchemaRegistryHost, cfg.SchemaRegistryUser, cfg.SchemaRegistryPassword)
	if err != nil {
		utils.Logger.Error("kafka client failed to new schema registry: %v", err)
		return nil, err
	}

	return &KafkaClient{
		saramaConfig:         saramaConfig,
		saramaClient:         saramaClient,
		saramaClusterAdmin:   saramaClusterAdmin,
		saramaSchemaRegistry: saramaSchemaRegistry,
	}, nil
}

func generateSaramaConfig(cfg *models.KafkaConfig) (*sarama.Config, error) {
	defaultSaramaConfig := sarama.NewConfig()
	saramaDefaultTimeout := time.Duration(_saramaTimeoutMs) * time.Millisecond

	defaultSaramaConfig.ClientID = cfg.ClientId
	defaultSaramaConfig.Version = sarama.V2_1_0_0
	defaultSaramaConfig.Net.DialTimeout = saramaDefaultTimeout
	defaultSaramaConfig.Net.ReadTimeout = saramaDefaultTimeout
	defaultSaramaConfig.Net.WriteTimeout = saramaDefaultTimeout
	defaultSaramaConfig.Metadata.Timeout = saramaDefaultTimeout
	defaultSaramaConfig.Producer.RequiredAcks = sarama.WaitForLocal
	defaultSaramaConfig.Producer.Flush.Frequency = _saramaTimeoutFlushMs * time.Millisecond

	if cfg.UseAuthentication {
		defaultSaramaConfig.Net.SASL.Mechanism = sarama.SASLMechanism(cfg.SaslMechanism)
		defaultSaramaConfig.Net.SASL.User = cfg.User
		defaultSaramaConfig.Net.SASL.Password = cfg.Password
		defaultSaramaConfig.Net.TLS.Enable = cfg.EnableTLS
		defaultSaramaConfig.Net.SASL.Enable = cfg.UseAuthentication
		if err := setAuthentication(defaultSaramaConfig); err != nil {
			return nil, err
		}
	}
	return defaultSaramaConfig, nil
}

func setAuthentication(conf *sarama.Config) error {
	switch conf.Net.SASL.Mechanism {
	case sarama.SASLTypeSCRAMSHA512:
		scram512Fn := func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
		conf.Net.SASL.SCRAMClientGeneratorFunc = scram512Fn
	case sarama.SASLTypeSCRAMSHA256:
		scram256Fn := func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
		conf.Net.SASL.SCRAMClientGeneratorFunc = scram256Fn
	case sarama.SASLTypePlaintext:
	default:
		return errors.New("kafka client invalid sasl mechanism")
	}
	return nil
}

func createTopic(topics []string, saramaClusterAdmin sarama.ClusterAdmin) error {
	for _, topic := range topics {
		err := saramaClusterAdmin.CreateTopic(topic,
			&sarama.TopicDetail{
				NumPartitions:     4,
				ReplicationFactor: 1,
			},
			false)
		if err != nil {
			utils.Logger.Error("kafka client create topic %v: %v", topic, err)
		}
	}
	return nil
}

// Close closes broker connections.
func (kc *KafkaClient) Close() error {
	kc.saramaClusterAdmin.Close()
	return kc.saramaClient.Close()
}
