package belajar_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")
	time.Sleep(1*time.Second) // Wait for goroutine to finish
}


// go test -v -run=TestCreateGoroutine

func DisplayNumber(num int){
	fmt.Println("Number: ", num)
}

func TestManyGoroutine(t *testing.T) {
	for i := 1; i <= 100000; i++ {
        go DisplayNumber(i)
    }
    time.Sleep(5 * time.Second) // Wait for all goroutine to finish
}