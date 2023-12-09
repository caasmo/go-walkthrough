package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	// Create a new errgroup with context
	g, ctx := errgroup.WithContext(context.Background())

	// Set the limit to the number of CPUs
	g.SetLimit(runtime.NumCPU())

	// Add tasks to the errgroup
	for i := 1; i <= 100; i++ {
		i := i // Create a new 'i' to avoid the late binding problem
		g.Go(func() error {
			// Check if the context has been cancelled or if there's an error
			if err := ctx.Err(); err != nil {
				return err
			}

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
	sleepDuration := rand.Intn(taskID) + 1

	// If sleepDuration is 42, return an error
	if sleepDuration == 10 {
		return errors.New("sleepDuration is 42")
	}

	fmt.Printf("Processing task %d with duration %d\n", taskID, sleepDuration)
	time.Sleep(time.Duration(sleepDuration) * time.Second)
	fmt.Printf("task %d END", taskID)


	return nil
}
