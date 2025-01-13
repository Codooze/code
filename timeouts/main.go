package main

import (
    "fmt"
    "runtime"
    "time"
)

// printGoroutineCount prints the current number of goroutines
func printGoroutineCount(scenario string) {
    fmt.Printf("[%s] Number of goroutines: %d\n", scenario, runtime.NumGoroutine())
}

// leakyOperation demonstrates a goroutine leak with unbuffered channel
func leakyOperation() {
    fmt.Println("\n=== Starting Leaky Operation ===")
    printGoroutineCount("Before leak")

    ch := make(chan string) // Unbuffered channel
    go func() {
        time.Sleep(2 * time.Second)
        // This send will block forever if no one receives
        ch <- "result"
        fmt.Println("This line will never be printed in leaky scenario")
    }()

    // Quick timeout, abandoning the goroutine
    select {
    case <-ch:
        fmt.Println("Received result")
    case <-time.After(100 * time.Millisecond):
        fmt.Println("Timeout")
    }

    printGoroutineCount("After leak")
}

// nonLeakyOperation demonstrates proper goroutine cleanup with buffered channel
func nonLeakyOperation() {
    fmt.Println("\n=== Starting Non-Leaky Operation ===")
    printGoroutineCount("Before operation")

    ch := make(chan string, 1) // Buffered channel
    go func() {
        time.Sleep(2 * time.Second)
        // This send will succeed even if no one receives
        ch <- "result"
        fmt.Println("Goroutine completed in non-leaky scenario")
    }()

    // Quick timeout, but goroutine will still complete
    select {
    case <-ch:
        fmt.Println("Received result")
    case <-time.After(100 * time.Millisecond):
        fmt.Println("Timeout")
    }

    printGoroutineCount("After operation")
}

func main() {
    // Print initial number of goroutines (usually 1 for main)
    printGoroutineCount("Start")

    // Test leaky scenario
    leakyOperation()

    // Test non-leaky scenario
    nonLeakyOperation() // a leak can still happen in buffered channels if we try to send more messages than the size of the channel

    // Wait a bit to see final goroutine count
    time.Sleep(3 * time.Second)
    printGoroutineCount("Final count")
}
