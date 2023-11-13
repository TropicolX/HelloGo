package main

import (
	"fmt"
	"sync"
	"time"
)

func someTask(id int, data chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes

	for taskId := range data {
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker: %d executed Task %d\n", id, taskId)
	}
}

func main() {
	// Creating a channel
	channel := make(chan int)

	// Creating a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Creating 10.000 workers to execute the task
	for i := 0; i < 10000; i++ {
		wg.Add(1) // Increment the WaitGroup counter before starting each goroutine
		go someTask(i, channel, &wg)
	}

	// Filling channel with 100.000 numbers to be executed
	for i := 0; i < 100000; i++ {
		channel <- i
	}

	// Close the channel after filling it
	close(channel)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers have finished.")
}
