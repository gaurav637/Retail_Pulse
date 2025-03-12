package enums

//JOB ENUMS-------------------> 
type JobStatus string

const (
	JobStatusOngoing   JobStatus = "ongoing"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
)
