package main

import (
    "context"
    "sync"
    "testing"
	"log"
)

var wg sync.WaitGroup
var Servers []*MockServer

func TestShutdownServers(t *testing.T) {

    // create a mock server for testing purposes
    mockServer := &MockServer{}

    // add the mock server to the Servers slice
    Servers = append(Servers, mockServer)

    // add the server to the WaitGroup
    wg.Add(1)

    // call the ShutdownServers function
    ShutdownServers()

    // check that the server was shut down correctly
    if !mockServer.IsShutdown {
        t.Error("Server was not shut down")
    }

    // check that the Servers slice was reset to nil
    if Servers != nil {
        t.Error("Servers slice was not reset to nil")
    }

    // check that the WaitGroup is empty
    wg.Wait()
}

// MockServer is a mock implementation of the http.Server type
type MockServer struct {
    IsShutdown bool
}

func (s *MockServer) Shutdown(ctx context.Context) error {
    s.IsShutdown = true
    return nil
}

func ShutdownServers() {
    defer wg.Wait()
    for _, server := range Servers {
        // gracefully shut down the server
        if err := server.Shutdown(context.Background()); err != nil {
            log.Fatal("Server shutdown failed:", err)
        }
        wg.Done() // signal that the WaitGroup is done waiting
    }
    Servers = nil
    return
}