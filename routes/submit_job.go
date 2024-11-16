package routes

import (
	"encoding/json"
	"net/http"
	"retail-pulse/models"
	"retail-pulse/utils"
	"sync"
)

var jobStore = make(map[string]*models.Job)
var jobMutex sync.Mutex

func SubmitJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload models.SubmitJobRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.Count != len(payload.Visits) {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	jobID := utils.GenerateJobID()
	job := models.NewJob(jobID, payload)

	jobMutex.Lock()
	jobStore[jobID] = job
	jobMutex.Unlock()

	go job.ProcessJob()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"job_id": jobID})
}
