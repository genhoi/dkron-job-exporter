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
	JobFailedCount       *prometheus.GaugeVec
	JobSuccessCount      *prometheus.GaugeVec
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
			Help: "Timestamp of the last successful job start",
		}, []string{"job"}),
		JobLastStart: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_job_last_start_ts",
			Help: "Timestamp of last job start",
		}, []string{"job"}),
		JobExecutionSeconds: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_job_execution_seconds",
			Help: "Time of successful job completion in seconds",
		}, []string{"job"}),
		JobFailedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_job_failed_count",
			Help: "Count of failed job execution",
		}, []string{"job"}),
		JobSuccessCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_job_success_count",
			Help: "Count of success job execution",
		}, []string{"job"}),
		DrkonApiUp: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "dkron_api_up",
			Help: "Availability of Dkron API",
		}, []string{}),
	}
}
