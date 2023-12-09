package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"time"
)

func main() {
	// Create a new errgroup with context
	g, ctx := errgroup.WithContext(context.Background())

	// Set the limit to 4
	g.SetLimit(4)

	// Add tasks to the errgroup
	for i := 0; i < 10; i++ {
		i := i // Create a new 'i' to avoid the late binding problem
		g.Go(func() error {
			return processTask(ctx, i)
		})
	}

	// Wait for all tasks to be processed
	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("All tasks processed successfully")
	}
}

// processTask is a function that simulates processing a task
func processTask(ctx context.Context, taskID int) error {
	// Generate a random sleep duration between 1 and taskID seconds
	rand.Seed(time.Now().UnixNano())
	if taskID == 0 {
		return errors.New("taskID cannot be 0")
	}
	sleepDuration := rand.Intn(taskID) + 1

	// If sleepDuration is 42, return an error
	if sleepDuration == 42 {
		return errors.New("sleepDuration is 42")
	}

	time.Sleep(time.Duration(sleepDuration) * time.Second)

	// Process task
	fmt.Printf("Processing task %d\n", taskID)

	return nil
}
