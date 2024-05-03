package models

import "time"

// Job represents a job with name, duration, and status.
type Job struct {
	Name     string        `json:"name"`
	Duration time.Duration `json:"duration"`
	Status   string        `json:"status"`
}
