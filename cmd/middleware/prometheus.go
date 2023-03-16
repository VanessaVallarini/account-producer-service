package middleware

import (
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	prom "github.com/prometheus/client_golang/prometheus"
)

const CKeyMetrics = "custom_metrics"

// See the NewMetrics func for proper descriptions and prometheus names!
// In case you add a metric here later, make sure to include it in the
// MetricsList method or you'll going to have a bad time.
type Metrics struct {
	customCnt *prometheus.Metric
	customDur *prometheus.Metric
}

// Needed by echo-contrib so echo can register and collect these metrics
func (m *Metrics) MetricList() []*prometheus.Metric {
	return []*prometheus.Metric{
		// ADD EVERY METRIC HERE!
		m.customCnt,
		m.customDur,
	}
}

// Creates and populates a new Metrics struct
// This is where all the prometheus metrics, names and labels are specified
func NewMetrics() *Metrics {
	return &Metrics{
		customCnt: &prometheus.Metric{
			Name:        "custom_total",
			Description: "Custom counter events.",
			Type:        "counter_vec",
			Args:        []string{"label_one", "label_two"},
		},
		customDur: &prometheus.Metric{
			Name:        "custom_duration_seconds",
			Description: "Custom duration observations.",
			Type:        "histogram_vec",
			Args:        []string{"label_one", "label_two"},
			Buckets:     prom.DefBuckets, // or your Buckets
		},
	}
}

// This will push your metrics object into every request context for later use
func (m *Metrics) AddCustomMetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(CKeyMetrics, m)
		return next(c)
	}
}

func (m *Metrics) IncCustomCnt(labelOne, labelTwo string) {
	labels := prom.Labels{"label_one": labelOne, "label_two": labelTwo}
	m.customCnt.MetricCollector.(*prom.CounterVec).With(labels).Inc()
}

func (m *Metrics) ObserveCustomDur(labelOne, labelTwo string, d time.Duration) {
	labels := prom.Labels{"label_one": labelOne, "label_two": labelTwo}
	m.customDur.MetricCollector.(*prom.HistogramVec).With(labels).Observe(d.Seconds())
}
