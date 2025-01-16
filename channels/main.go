package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	/*
	   messages := make(chan string)
	   mess2 := make(chan string,2)

	   go func() { messages <- "ping" }()
	   mess2 <- "Buffered"
	   mess2 <- "channel"

	   msg := <-messages
	   fmt.Println(msg)
	   fmt.Println(<-mess2)
	   fmt.Println(<-mess2)


	   done := make(chan bool, 1)
	   go worker(done)

	   <-done
	*/
	//  ---------------

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// * Example select
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	//*End
}
