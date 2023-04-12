package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	onceConfig    sync.Once
	metricConfigs = &Metrics{}
)

type Metrics struct {
	KafkaStrategyProducedMessagesSuccessCounter *prometheus.CounterVec
	KafkaStrategyProducedMessagesFailCounter    *prometheus.CounterVec
	ApiStrategyErrosCounter                     *prometheus.CounterVec
	ApiStrategySuccessCounter                   *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	onceConfig.Do(func() {
		metricConfigs.KafkaStrategyProducedMessagesSuccessCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "kafka_produced_messages_success",
			Help: "Kafka produced messages success counter",
		}, []string{"topic", "partition", "offset"})
		prometheus.MustRegister(metricConfigs.KafkaStrategyProducedMessagesSuccessCounter)

		metricConfigs.KafkaStrategyProducedMessagesFailCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "kafka_produced_messages_fail",
			Help: "Kafka produced messages fail counter",
		}, []string{"topic", "partition", "offset"})
		prometheus.MustRegister(metricConfigs.KafkaStrategyProducedMessagesFailCounter)

		metricConfigs.ApiStrategyErrosCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "api_errors",
			Help: "api errors counter",
		}, []string{"statusCodeStr", "method", "path"})
		prometheus.MustRegister(metricConfigs.ApiStrategyErrosCounter)

		metricConfigs.ApiStrategySuccessCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "api_success",
			Help: "api success counter",
		}, []string{"statusCodeStr", "method", "path"})
		prometheus.MustRegister(metricConfigs.ApiStrategySuccessCounter)
	})
	return metricConfigs
}
