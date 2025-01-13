package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    var ops atomic.Uint64
    var wg sync.WaitGroup

    // Start 50 goroutines to increment the counter.
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func() {
            for c := 0; c < 1000; c++ {
                ops.Add(1)
                time.Sleep(time.Millisecond) // Simulate work
            }
            wg.Done()
        }()
    }

    // Start a goroutine to monitor the counter's value.
    go func() {
        for {
            currentOps := ops.Load()
            fmt.Println("Current ops:", currentOps)
            if currentOps >= 50000 {
                break
            }
            time.Sleep(100 * time.Millisecond) // Adjust the frequency as needed
        }
    }()

    // Wait for all incrementing goroutines to finish.
    wg.Wait()

    // Give the monitoring goroutine time to print the final value.
    time.Sleep(200 * time.Millisecond)
    fmt.Println("Final ops:", ops.Load())
}

