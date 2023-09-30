package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
	"fmt"
    "github.com/gorilla/mux"
)

func TestGetHandler(t *testing.T) {
    // Create a new router
    router := mux.NewRouter()

    // Define a handler for GET requests to the root path
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, World!")
    }).Methods("GET")

    // Create a new test request to the root path
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new mock HTTP server
    rr := httptest.NewRecorder()

    // Call the router's ServeHTTP method to process the request
    router.ServeHTTP(rr, req)

    // Check the status code returned by the server
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the body of the response
    expected := "Hello, World!"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
