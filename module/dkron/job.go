package dkron

import "time"

type Job struct {
	// Job id. Must be unique, it's a copy of name.
	ID string `json:"id"`

	// Job name. Must be unique, acts as the id.
	Name string `json:"name"`

	// Display name of the job. If present, displayed instead of the name
	DisplayName string `json:"displayname"`

	// The timezone where the cron expression will be evaluated in.
	// Empty means local time.
	Timezone string `json:"timezone"`

	// Cron expression for the job. When to run the job.
	Schedule string `json:"schedule"`

	// Arbitrary string indicating the owner of the job.
	Owner string `json:"owner"`

	// Email address to use for notifications.
	OwnerEmail string `json:"owner_email"`

	// Number of successful executions of this job.
	SuccessCount int `json:"success_count"`

	// Number of errors running this job.
	ErrorCount int `json:"error_count"`

	// Last time this job executed successfully.
	LastSuccess time.Time `json:"last_success"`

	// Last time this job failed.
	LastError time.Time `json:"last_error"`

	// Is this job disabled?
	Disabled bool `json:"disabled"`

	// Tags of the target servers to run this job against.
	Tags map[string]string `json:"tags"`

	// Job metadata describes the job and allows filtering from the API.
	Metadata map[string]string `json:"metadata"`

	// Number of times to retry a job that failed an execution.
	Retries uint `json:"retries"`

	// Jobs that are dependent upon this one will be run after this job runs.
	DependentJobs []string `json:"dependent_jobs"`

	// Job id of job that this job is dependent upon.
	ParentJob string `json:"parent_job"`

	// Concurrency policy for this job (allow, forbid).
	Concurrency string `json:"concurrency"`

	// Executor plugin to be used in this job.
	Executor string `json:"executor"`

	// Computed job status.
	Status string `json:"status"`

	// Computed next execution.
	Next time.Time `json:"next"`

	// Delete the job after the first successful execution.
	Ephemeral bool `json:"ephemeral"`
}
