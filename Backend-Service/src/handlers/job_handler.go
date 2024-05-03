package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ABHI2598/Backend-Service/src/models"
	"github.com/ABHI2598/Backend-Service/src/scheduler"
)

// JobHandler handles HTTP requests related to jobs.
type JobHandler struct {
	Scheduler *scheduler.SJFScheduler
}

func (jh *JobHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs := jh.Scheduler.GetJobs()
	respondWithJSON(w, http.StatusOK, jobs)
}

func (jh *JobHandler) SubmitJob(w http.ResponseWriter, r *http.Request) {
	var newJob models.Job
	if err := json.NewDecoder(r.Body).Decode(&newJob); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Add the submitted job to the scheduler
	jh.Scheduler.AddJob(newJob)

	respondWithJSON(w, http.StatusCreated, newJob)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
