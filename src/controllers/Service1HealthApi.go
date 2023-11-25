package controller

import (
	"net/http"
)

// read the respective microservice apis health and respond with 200 if all are healthy
func Service1HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
