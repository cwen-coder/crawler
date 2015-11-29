package main

import (
	"fmt"
	"time"
)

func consumer(newC chan chan int) {
	var jobs chan int
	for {
		select {
		case c, ok := <-newC:
			if !ok {
				return
			}
			jobs = c
			fmt.Println("new job pipeline")
		case job := <-jobs:
			fmt.Println("executing job:", job)
		}
	}
}

func producer(newC chan chan int) {
	var i int
	var c chan int
	for {
		if i%3 == 0 {
			c = make(chan int)
			newC <- c
		}
		c <- i
		i = (i + 1) % 3
	}
}

func main() {
	newC := make(chan chan int)
	go producer(newC)
	go consumer(newC)
	time.Sleep(5 * time.Second)
}
