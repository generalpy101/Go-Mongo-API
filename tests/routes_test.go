package routes_tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	routes "github.com/generalpy101/Go-Mongo-API/routers"
)

// Using httptest to test the routes

// Test the /test route
func TestTestRoute(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := routes.Router()

	// Serve the HTTP request to our recorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := "{\"message\": \"Server is up and running\"}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
