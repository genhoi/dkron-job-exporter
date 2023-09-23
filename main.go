package main

import (
	"github.com/genhoi/dkron-job-exporter/app"
	"github.com/genhoi/dkron-job-exporter/module/dkron"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	version string
)

func main() {
	container, err := app.CreateContainer(version)
	if err != nil {
		log.Fatal("Error CreateContainer ", err)
	}

	srv := container.GetHttpServer()
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("Error ListenAndServe ", err)
		}
	}()

	go func() {
		for {
			metrics := container.GetMetrics()
			err := updateMetrics(container.DkronClient(), metrics)
			if err != nil {
				metrics.DrkonApiUp.WithLabelValues().Set(0)
				log.Print("Error updateMetrics ", err)
			} else {
				metrics.DrkonApiUp.WithLabelValues().Set(1)
			}

			time.Sleep(5 * time.Second)
		}
	}()

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt)
	<-cancelChan
}

func updateMetrics(dkronClient *dkron.Client, metrics *app.Metrics) error {
	jobs, err := dkronClient.GetJobs()
	if err != nil {
		return err
	}
	for _, job := range jobs {
		executions, err := dkronClient.ListExecutionsByJob(job.Name)
		if err != nil {
			return err
		}

		for _, execution := range executions {
			if execution.Success {
				metrics.JobLastSuccessfulRun.With(prometheus.Labels{"job": job.Name}).Set(float64(execution.FinishedAt.Unix()))
				metrics.JobExecutionSeconds.With(prometheus.Labels{"job": job.Name}).Set(
					float64(execution.FinishedAt.Unix()) - float64(execution.StartedAt.Unix()),
				)
				break
			}
		}
		for _, execution := range executions {
			metrics.JobLastStart.With(prometheus.Labels{"job": job.Name}).Set(float64(execution.StartedAt.Unix()))
			break
		}

		metrics.JobFailedCount.With(prometheus.Labels{"job": job.Name}).Set(float64(job.ErrorCount))
		metrics.JobSuccessCount.With(prometheus.Labels{"job": job.Name}).Set(float64(job.SuccessCount))
	}

	return nil
}
