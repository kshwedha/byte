package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kshwedha/core/src/api/handlers"
)

func TestUpdateUserHandler(t *testing.T) {
	// Create a new request with a sample request body
	req, err := http.NewRequest("GET", "/api/users/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}", handlers.UpdateModuleHandler)

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedResponse := "Updating user with ID: 123"
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rr.Body.String())
	}
}
