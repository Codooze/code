package main

import (
    "fmt"
    "time"
)
//If you don’t want to block but still want to handle the timer’s expiration, you can use a goroutine:
go
Copy code

func main() {
    timer := time.NewTimer(2 * time.Second)

    go func() {
        <-timer.C // Wait for the timer to expire in a separate goroutine
        fmt.Println("Timer expired in goroutine!")
    }()

    fmt.Println("Doing other work...")
    time.Sleep(3 * time.Second) // Let the program run long enough for the timer to expire
    fmt.Println("Program finished!")
}

