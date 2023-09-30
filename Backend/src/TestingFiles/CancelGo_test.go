package main

import (
    "context"
    "testing"
    "time"
)

func TestCancelGoroutine(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    done := make(chan bool)

    // Start a goroutine that does some work
    go func(ctx context.Context) {
        // Simulate some work
        time.Sleep(2 * time.Second)

        // Check if context is cancelled
        select {
        case <-ctx.Done():
            t.Logf("Goroutine cancelled: %v", ctx.Err())
            done <- true
        default:
            t.Error("Goroutine did not cancel")
            done <- false
        }
    }(ctx)

    // Cancel the goroutine
    cancel()

    // Wait for goroutine to finish
    <-done
}
