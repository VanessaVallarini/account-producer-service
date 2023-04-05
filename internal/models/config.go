package models

type Config struct {
	AppName          string
	ServerHost       string
	HealthServerHost string
	Database         *DatabaseConfig
	Kafka            *KafkaConfig
	ViaCep           *ViaCepConfig
}

type DatabaseConfig struct {
	DatabaseUser                string
	DatabasePassword            string
	DatabaseKeyspace            string
	DatabaseHost                string
	DatabasePort                int
	DatabaseConnectionRetryTime int
	DatabaseRetryMinArg         int
	DatabaseRetryMaxArg         int
	DatabaseNumRetries          int
	DatabaseClusterTimeout      int
}

type KafkaConfig struct {
	ClientId               string
	Hosts                  []string
	SchemaRegistryHost     string
	Acks                   string
	Timeout                int
	UseAuthentication      bool
	EnableTLS              bool
	SaslMechanism          string
	User                   string
	Password               string
	SchemaRegistryUser     string
	SchemaRegistryPassword string
	EnableEvents           bool
	MaxMessageBytes        int
	RetryMax               int
	ConsumerTopic          []string
	ConsumerGroup          string
	ProducerTopic          []string
}

type ViaCepConfig struct {
	Url                   string
	MaxRetriesHttpRequest int
	MaxFailureRatio       float64
	Name                  string
}
