package main

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
	"BackendPkg"
)


func TestListenPostReq(t *testing.T) {

    // create a mock server
    srv := httptest.NewServer(http.HandlerFunc(BackendPkg.NewUserResponse))
    defer srv.Close()

    // send a POST request to the server
    payload := `{"user": {"name": "Alice", "email": "alice@example.com"}}`
    req, err := http.NewRequest("POST", srv.URL+"/api/Signup", bytes.NewBufferString(payload))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    // assert that the response status code is OK
    if resp.StatusCode != http.StatusOK {
        t.Errorf("expected status code %v but got %v", http.StatusOK, resp.StatusCode)
    }
}
