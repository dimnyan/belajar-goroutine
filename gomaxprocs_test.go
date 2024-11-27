package belajar_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxprocs(t *testing.T) {

	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
	    group.Add(1)
		go func() {
            defer group.Done()
			time.Sleep(3 * time.Second)
        }()
	
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU cores:", totalCpu)

	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total threads:", totalThreads)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutines:", totalGoroutine)

	group.Wait()

}

func TestChangeThread(t *testing.T) {

	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
	    group.Add(1)
		go func() {
            defer group.Done()
			time.Sleep(3 * time.Second)
        }()
	
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU cores:", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total threads:", totalThreads)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutines:", totalGoroutine)

	group.Wait()

}