package main

import (
	"fmt"
	"sync"
	"time"
)

func parallel_task(task int, done ...chan int) {
	fmt.Println("Parallel task #: ", task)
	time.Sleep(time.Second)
	if done != nil && len(done) != 0 {
		done[0] <- 0 // Signal completion
	}
}

func main() {
	// Create the wait group (basically a semaphore)
	var wait_group sync.WaitGroup

	// Start a Timer
	before := time.Now()

	// Create a loop to run tasks in order
	for i := 0; i < 10; i++ {
		wait_group.Add(1)
		i := i
		go func() {
			// Run the task and indicate completion once it's done
			defer wait_group.Done()
			parallel_task(i)
		}()
		// Require that we wait for the wait_group to continute tasks
		wait_group.Wait()
	}
	// Capture end time
	after := time.Now()

	// Print the result
	fmt.Println("Finished tasks in: ", after.Sub(before))
	fmt.Printf("\n")

	// Retry for parallel execution
	// Start a new Timer
	before = time.Now()

	// Create a new loop that runs tasks in parallel
	done := make(chan int)
	for i := 0; i < 1000; i++ {
		go parallel_task(i, done)
	}

	// Wait for completed execution of all threads before continuing
	<-done

	// Capture end time
	after = time.Now()

	// Print new result
	fmt.Println("Finished tasks in: ", after.Sub(before))
}
