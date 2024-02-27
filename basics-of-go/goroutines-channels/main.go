package main

import (
	"fmt"
	"time"
)

func printMessages(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond)
	}
}

func pushMessages(message string, channel chan string) {
	for i := 0; i < 10; i++ {
		m := fmt.Sprintf("%v - %v", i, message)
		channel <- m
	}
}

func main() {
	// goroutines
	go printMessages("thread 1")
	go printMessages("thread 2")
	printMessages("main thread")

	// goroutines with buffered channel
	channel := make(chan string, 20)
	go pushMessages("main thread", channel)
	go pushMessages("thread 1", channel)

	for i := 0; i < 20; i++ {
		fmt.Println(<-channel)
	}
}
