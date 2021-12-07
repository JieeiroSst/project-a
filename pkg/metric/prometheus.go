package metric

import (
	"github.com/JieeiroSst/itjob/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type prometheusService struct {
	pHistogram           *prometheus.HistogramVec
	httpRequestHistogram *prometheus.HistogramVec
	config 				 *config.Config
}

type PrometheusService interface {
	SaveCLI(c *CLI) error
	SaveHTTP(h *HTTP)
}

func NewPrometheusService(config *config.Config) (PrometheusService, error) {
	cli := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "pushgateway",
		Name:      "cmd_duration_seconds",
		Help:      "CLI application execution in seconds",
		Buckets:   prometheus.DefBuckets,
	}, []string{"name"})
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	s := &prometheusService{
		pHistogram:           cli,
		httpRequestHistogram: http,
	}
	err := prometheus.Register(s.pHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	err = prometheus.Register(s.httpRequestHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	return &prometheusService{
		pHistogram:           s.pHistogram,
		httpRequestHistogram: s.httpRequestHistogram,
		config:               config,
	}, nil
}

//SaveCLI send metrics to server
func (s *prometheusService) SaveCLI(c *CLI) error {
	gatewayURL := s.config.Prometheus.PrometheusPushgateway
	s.pHistogram.WithLabelValues(c.Name).Observe(c.Duration)
	return push.New(gatewayURL, "cmd_job").Collector(s.pHistogram).Push()
}

//SaveHTTP send metrics to server
func (s *prometheusService) SaveHTTP(h *HTTP) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}


