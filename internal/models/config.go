package models

type Config struct {
	AppName          string
	ServerHost       string
	HealthServerHost string
	Kafka            *KafkaConfig
	ViaCep           *ViaCepConfig
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
	DlqTopic               string
	ConsumerTopic          string
	ConsumerGroup          string
}

type ViaCepConfig struct {
	Url                   string
	MaxRetriesHttpRequest int
	MaxFailureRatio       float64
	Name                  string
}
