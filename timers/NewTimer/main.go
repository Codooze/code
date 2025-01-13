package main

import (
	"fmt"
	"time"
)

func main() {
	// Using Timer
	timer := time.NewTimer(5 * time.Second)

	go func() {
    fmt.Println("⭐started⭐")
		<-timer.C
		fmt.Println("Timer expired! ⏰")
	}()
  fmt.Println("sleeeping...zzzzz")
	time.Sleep(4 * time.Second) //sleeps 

	if timer.Stop() {
		fmt.Println("Timer stopped before expiration!")
	} else {
		fmt.Println("Timer already expired.")
	}

	// Using Sleep (blocking)
	fmt.Println("Sleeping for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("Sleep finished!")
}

