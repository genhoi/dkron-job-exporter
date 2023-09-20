package app

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	JobLastSuccessfulRun *prometheus.GaugeVec
	JobLastStart         *prometheus.GaugeVec
}

func CreateMetrics() *Metrics {
	return &Metrics{
		JobLastSuccessfulRun: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_job_last_successful_run_ts",
			Help: "Время последнего успешного запуска задачи",
		}, []string{"job"}),
		JobLastStart: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_job_last_start_ts",
			Help: "Время последнего старта задачи",
		}, []string{"job"}),
	}
}
