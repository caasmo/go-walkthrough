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

// TaskResult is a struct that contains the sleep duration, the result, and the error
type TaskResult struct {
	TaskId int
	SleepDuration int
	Result        int
	Error         error
}

func main() {
	// Create a new errgroup with context
	g, ctx := errgroup.WithContext(context.Background())

	// Set the limit to the number of CPUs
	g.SetLimit(runtime.NumCPU())

	// Create a channel to store the results
	results := make(chan TaskResult)
	defer close(results)
	// Process the results
    go func() {
        for result := range results {
            fmt.Printf("Task Id: %d, Duration: %d, Result: %d, err %q\n", result.TaskId, result.SleepDuration, result.Result, result.Error)
        }
    }()

	// Add tasks to the errgroup
	for i := 1; i <= 100; i++ {
		i := i // Create a new 'i' to avoid the late binding problem
		g.Go(func() error {
			// Check if the context has been cancelled or if there's an error
			//if err := ctx.Err(); err != nil {
			//	return err
			//}

			result := processTask(ctx, i)
			results <- result

			return nil
		})
	}

    g.Wait()
    fmt.Println("All tasks processed successfully")
}

// processTask is a function that simulates processing a task
func processTask(ctx context.Context, taskID int) TaskResult {
	// Generate a random sleep duration between 1 and taskID seconds
	rand.Seed(time.Now().UnixNano())
	sleepDuration := rand.Intn(10) + 1

    res := TaskResult{
        TaskId: taskID,
		SleepDuration: sleepDuration,
    }

	// If sleepDuration is 42, return an error
	if sleepDuration == 10 {
        res.Error = errors.New("Task Duration is 10")
        return res
	}

	time.Sleep(time.Duration(sleepDuration) * time.Second)

    res.Result = taskID * 2
    return res
}
