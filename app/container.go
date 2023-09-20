package app

import (
	"github.com/genhoi/dkron-job-exporter/module/dkron"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"sync"
)

var (
	onceDkronClient sync.Once
	onceHttpServer  sync.Once
	onceMetrics     sync.Once
)

type Container struct {
	config      config
	dkronClient *dkron.Client
	httpServer  *http.Server
	metrics     *Metrics
}

type config struct {
	listenAddr string
	dkronApi   string
}

func (c *Container) GetMetrics() *Metrics {
	onceMetrics.Do(func() {
		c.metrics = CreateMetrics()
	})

	return c.metrics
}

func (c *Container) GetHttpServer() *http.Server {
	onceHttpServer.Do(func() {
		mux := http.NewServeMux()
		mux.Handle("/", c.getPrometheusHandler())

		c.httpServer = &http.Server{
			Addr:    c.config.listenAddr,
			Handler: mux,
		}
	})

	return c.httpServer
}

func (c *Container) getPrometheusHandler() http.Handler {
	registry := prometheus.NewRegistry()
	registry.MustRegister(c.GetMetrics().JobLastSuccessfulRun)
	registry.MustRegister(c.GetMetrics().JobLastStart)

	return promhttp.InstrumentMetricHandler(
		registry,
		promhttp.HandlerFor(registry, promhttp.HandlerOpts{}),
	)
}

func (c *Container) DkronClient() *dkron.Client {
	onceDkronClient.Do(func() {
		c.dkronClient = dkron.CreateClient(c.config.dkronApi)
	})

	return c.dkronClient
}

func CreateContainer() (*Container, error) {
	err := godotenv.Load(".env.local")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	listen := os.Getenv("DKRON_JOB_EXPORTER_ADDR")
	dkronApi := os.Getenv("DKRON_JOB_EXPORTER_DKRON_API")

	conf := config{
		listenAddr: listen,
		dkronApi:   dkronApi,
	}

	return &Container{config: conf}, nil
}
