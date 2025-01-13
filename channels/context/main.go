package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("worker %d shutting down\n", id)
            return
        case job, ok := <-jobs:
            if !ok {
                return
            }
            // Simulate work
            time.Sleep(time.Millisecond * 100)
            results <- job * 2
        }
    }
}

func main() {
    // Create a context that will timeout after 2 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // Clean up

    jobs := make(chan int, 5)
    results := make(chan int, 5)

    // Start 3 workers
    for i := 1; i <= 3; i++ {
        go worker(ctx, i, jobs, results)
    }

    // Send some jobs
    go func() {
        for i := 1; i <= 10; i++ {
            select {
            case <-ctx.Done():
                fmt.Println("job sender shutting down")
                return
            case jobs <- i:
                fmt.Printf("sent job %d\n", i)
            }
        }
    }()

    // Collect results until context is done
    count := 0
    for {
        select {
        case <-ctx.Done():
            fmt.Println("main: context done, exiting")
            return
        case result := <-results:
            count++
            fmt.Printf("got result: %d\n", result)
        }
    }
}
/*
Let me show you how context works with multiple goroutines and channels.

This example demonstrates several important patterns:

1. Context is used to coordinate shutdown across multiple goroutines:
   - Each worker checks `ctx.Done()` to know when to quit
   - The job sender checks context before sending
   - The main routine checks context while collecting results

2. Context doesn't interfere with normal channel operations:
   - Workers can still receive jobs and send results
   - Multiple goroutines can share the same context
   - Channel operations continue normally until context cancellation

3. Context provides clean shutdown:
   - When timeout occurs, all goroutines get the signal via `ctx.Done()`
   - No goroutines are left hanging
   - Resources can be cleaned up properly

Some key benefits of using context:
- You don't need to create separate "done" channels
- Context can be passed through function calls
- You can have timeout, cancellation, or both
- Multiple goroutines can share the same shutdown signal
*/


