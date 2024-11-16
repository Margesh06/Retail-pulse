package main

import (
	"log"
	"net/http"
	"retail-pulse/routes"
)

func main() {
	http.HandleFunc("/api/submit/", routes.SubmitJob)
	http.HandleFunc("/api/status", routes.GetJobStatus)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
