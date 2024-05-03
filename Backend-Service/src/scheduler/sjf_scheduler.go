package scheduler

import (
	"time"

	"github.com/ABHI2598/GOASSIGNMENT/Backend-Service/src/models"
)

// SJFScheduler implements the Shortest Job First (SJF) scheduling algorithm.
type SJFScheduler struct {
	Jobs []models.Job
}

// NewSJFScheduler creates a new SJFScheduler instance.
func NewSJFScheduler() *SJFScheduler {
	return &SJFScheduler{}
}

// AddJob adds a new job to the scheduler.
func (s *SJFScheduler) AddJob(job models.Job) {
	s.Jobs = append(s.Jobs, job)
}

// GetJobs returns the current list of jobs.
func (s *SJFScheduler) GetJobs() []models.Job {
	return s.Jobs
}

// Schedule starts scheduling jobs using the SJF algorithm.
func (s *SJFScheduler) Schedule() {
	for {
		// Sort jobs by duration
		for i := range s.Jobs {
			minIdx := i
			for j := i + 1; j < len(s.Jobs); j++ {
				if s.Jobs[j].Duration < s.Jobs[minIdx].Duration {
					minIdx = j
				}
			}
			s.Jobs[i], s.Jobs[minIdx] = s.Jobs[minIdx], s.Jobs[i]
		}

		// Execute jobs
		for i, job := range s.Jobs {
			s.Jobs[i].Status = "running"

			// Broadcast job status update
			// Logic to broadcast job status update
			time.Sleep(job.Duration * time.Second)

			s.Jobs[i].Status = "completed"

			// Broadcast job status update
			// Logic to broadcast job status update
		}
	}
}
