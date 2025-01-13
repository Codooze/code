package main

import (
	"fmt"
	"sync"
	"time"
)

func makeBreakfast(item string, duration time.Duration) {
    fmt.Printf("Starting to make %s\n", item)
    time.Sleep(duration) // Simulate cooking time
    fmt.Printf("Finished making %s\n", item)
}

// Without goroutines (sequential)
func breakfastSequential() {
    start := time.Now()
    
    makeBreakfast("coffee", 3*time.Second)
    makeBreakfast("toast", 2*time.Second)
    makeBreakfast("eggs", 5*time.Second)
    
    fmt.Printf("Total time: %v\n", time.Since(start))
}
 
// With goroutines (concurrent)
func breakfastConcurrent() {
    start := time.Now()
    
    // Start all tasks concurrently
    go makeBreakfast("coffee", 3*time.Second)
    go makeBreakfast("toast", 2*time.Second)
    go makeBreakfast("eggs", 5*time.Second)
    
    // Wait a bit to see results (in real code, we'd use WaitGroup)
    time.Sleep(6 * time.Second)
    
    fmt.Printf("Total time: %v\n", time.Since(start))
}


//  ------------------------------
// Simulated file download
func downloadFile(filename string, wg *sync.WaitGroup) {
    defer wg.Done() // Mark this task as done when finished
    
    fmt.Printf("Starting download: %s\n", filename)
    // Simulate different download times
    time.Sleep(time.Duration(2+len(filename)) * time.Second)
    fmt.Printf("Finished download: %s\n", filename)
}

func main() {
    // fmt.Println("Sequential breakfast:")
    // breakfastSequential()
    
    // fmt.Println("\nConcurrent breakfast:")
    // breakfastConcurrent()

	// -------------------------------------
	  files := []string{"photo.jpg", "document.pdf", "music.mp3"}
    
    // Create a WaitGroup to track all downloads
    var wg sync.WaitGroup
    wg.Add(len(files)) // Set counter to number of files
    
    fmt.Println("Starting downloads...")
    start := time.Now()
    
    // Start each download in its own goroutine
    for _, file := range files {
        go downloadFile(file, &wg)
    }
    
    // Wait for all downloads to complete
    wg.Wait()
    
    fmt.Printf("All downloads completed in %v\n", time.Since(start))
}
