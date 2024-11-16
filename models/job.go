package models

import (
	"fmt"
	"math/rand"
	"net/http"
	"retail-pulse/utils"
	"sync"
	"time"
)

type SubmitJobRequest struct {
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

type Job struct {
	JobID   string     `json:"job_id"`
	Status  string     `json:"status"`
	Errors  []JobError `json:"errors"`
	Mutex   sync.Mutex `json:"-"`
	Payload SubmitJobRequest
}

type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

func NewJob(jobID string, payload SubmitJobRequest) *Job {
	return &Job{
		JobID:   jobID,
		Status:  "ongoing",
		Payload: payload,
	}
}

func (job *Job) ProcessJob() {
	for _, visit := range job.Payload.Visits {
		for _, url := range visit.ImageURLs {
			if err := processImage(url); err != nil {
				job.Mutex.Lock()
				job.Errors = append(job.Errors, JobError{
					StoreID: visit.StoreID,
					Error:   err.Error(),
				})
				job.Mutex.Unlock()
			}
		}
	}
	job.Mutex.Lock()
	if len(job.Errors) > 0 {
		job.Status = "failed"
	} else {
		job.Status = "completed"
	}
	job.Mutex.Unlock()
}

func processImage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	width, height := utils.GetImageDimensions(resp.Body)
	fmt.Printf("Image Perimeter: %d\n", 2*(width+height))

	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
	return nil
}
