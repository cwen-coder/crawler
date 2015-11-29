package main

import (
	"fmt"
	"time"
)

var chan1 chan int
var chanLength int = 18
var interval time.Duration = 1500 * time.Millisecond

func main() {
	chan1 = make(chan int, chanLength)
	go func() {
		for i := 0; i < chanLength; i++ {
			if i > 0 && i%3 == 0 {
				fmt.Println("Reset chan1 ....")
				chan1 = make(chan int, chanLength)
			}

			fmt.Printf("Send element %d...\n", i)
			chan1 <- i
			time.Sleep(interval)
		}

		fmt.Println("Close chan1...")
		close(chan1)
	}()
	receive()
}

func receive() {
	fmt.Println("Begin to receive elementd from chan1...")
	timer := time.After(30 * time.Second)
Loop:
	for {
		select {
		case e, ok := <-getChan():
			if !ok {
				fmt.Println("--chan1 closed.")
				break Loop
			}

			fmt.Printf("Received an element: %d\n", e)
			time.Sleep(interval)
		case <-timer:
			fmt.Println("Timeout!")
			break Loop
		}
	}

	fmt.Println("--end")
}

func getChan() chan int {
	return chan1
}
