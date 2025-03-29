package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

// work simulates some work being done
func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate some work
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Goroutine %d completed its work\n", id)
}

func main() {
	// Create a trace file
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Printf("Failed to create trace file: %v\n", err)
		return
	}
	defer f.Close()

	// Start tracing
	if err := trace.Start(f); err != nil {
		fmt.Printf("Failed to start trace: %v\n", err)
		return
	}
	defer trace.Stop()

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start 20 goroutines
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go work(i, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All goroutines completed!")
}
