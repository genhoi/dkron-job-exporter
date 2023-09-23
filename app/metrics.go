package app

import (
	"github.com/prometheus/client_golang/prometheus"
	"runtime"
)

type Metrics struct {
	Version              prometheus.GaugeFunc
	JobLastSuccessfulRun *prometheus.GaugeVec
	JobLastStart         *prometheus.GaugeVec
	JobExecutionSeconds  *prometheus.GaugeVec
	DrkonApiUp           *prometheus.GaugeVec
}

func CreateMetrics(version string) *Metrics {
	return &Metrics{
		Version: prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "build_info",
			Help: "A metric with a constant '1' value labeled by version from whichwas built.",
			ConstLabels: prometheus.Labels{
				"version":   version,
				"goversion": runtime.Version(),
			},
		},
			func() float64 { return 1 },
		),
		JobLastSuccessfulRun: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_job_last_successful_run_ts",
			Help: "Время последнего успешного запуска задачи",
		}, []string{"job"}),
		JobLastStart: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_job_last_start_ts",
			Help: "Время последнего старта задачи",
		}, []string{"job"}),
		JobExecutionSeconds: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_job_execution_seconds",
			Help: "Время успешного выполнения задачи",
		}, []string{"job"}),
		DrkonApiUp: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_api_up",
			Help: "Доступность api Dkron",
		}, []string{}),
	}
}
