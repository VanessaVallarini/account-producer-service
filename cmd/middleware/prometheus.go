package middleware

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	prom "github.com/prometheus/client_golang/prometheus"
)

const CKeyMetrics = "custom_metrics"

// See the NewMetrics func for proper descriptions and prometheus names!
// In case you add a metric here later, make sure to include it in the
// MetricsList method or you'll going to have a bad time.
type Metrics struct {
	errorResponse *prometheus.Metric
}

// Needed by echo-contrib so echo can register and collect these metrics
func (m *Metrics) MetricList() []*prometheus.Metric {
	return []*prometheus.Metric{
		// ADD EVERY METRIC HERE!
		m.errorResponse,
	}
}

// Creates and populates a new Metrics struct
// This is where all the prometheus metrics, names and labels are specified
func NewMetrics() *Metrics {
	return &Metrics{
		errorResponse: &prometheus.Metric{
			Name:        "error_response",
			Description: "Custom error response",
			Type:        "counter_vec",
			Args:        []string{"status", "method", "path"},
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

func (m *Metrics) IncErrorResponse(status, method, path string) {
	labels := prom.Labels{"status": status, "method": method, "path": path}
	m.errorResponse.MetricCollector.(*prom.CounterVec).With(labels).Inc()
}
