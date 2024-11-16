package routes

import (
	"encoding/json"
	"net/http"
	"sync"
)

var statusMutex sync.Mutex

func GetJobStatus(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobid")
	if jobID == "" {
		http.Error(w, "Missing jobid", http.StatusBadRequest)
		return
	}

	statusMutex.Lock()
	job, exists := jobStore[jobID]
	statusMutex.Unlock()

	if !exists {
		http.Error(w, "Job ID not found", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"job_id": jobID,
		"status": job.Status,
	}

	if job.Status == "failed" {
		response["error"] = job.Errors
	}

	json.NewEncoder(w).Encode(response)
}
