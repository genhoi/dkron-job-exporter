package dkron

import "time"

// Execution type holds all of the details of a specific Execution.
type Execution struct {
	// Id is the Key for this execution
	Id string `json:"id,omitempty"`

	// Name of the job this executions refers to.
	JobName string `json:"job_name,omitempty"`

	// Start time of the execution.
	StartedAt time.Time `json:"started_at,omitempty"`

	// When the execution finished running.
	FinishedAt time.Time `json:"finished_at,omitempty"`

	// If this execution executed successfully.
	Success bool `json:"success"`

	// Partial output of the execution.
	Output string `json:"output,omitempty"`

	// Node name of the node that run this execution.
	NodeName string `json:"node_name,omitempty"`

	// Execution group to what this execution belongs to.
	Group int64 `json:"group,omitempty"`

	// Retry attempt of this execution.
	Attempt uint `json:"attempt,omitempty"`
}
