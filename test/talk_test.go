package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kshwedha/core/src/lib"
)

func TestRespond(t *testing.T) {
	// Create a new request with a sample request body
	reqBody := []byte(`{"message": "Hello"}`)
	req, err := http.NewRequest("POST", "/api/talk", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Call the Respond function
	lib.Respond(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedResponse := []byte(`{"result": "Hello, World!"}`)
	if !bytes.Equal(rr.Body.Bytes(), expectedResponse) {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, rr.Body.String())
	}
}
