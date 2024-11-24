package belajar_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func OnlyIn(channel chan<- string) {
    channel <- "Hello, World!"
}

func OnlyOut(channel <- chan string){
	fmt.Println(<-channel)
}

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}