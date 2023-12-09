package main

import (
	"runtime"
	"sync"
)

type Job struct {
	id       int
	payload string
}

type Result struct {
	job         Job
	result int
    err error
}
var numWorkers = runtime.NumCPU()
var jobCh = make(chan string, numWorkers)
var resultCh = make(chan string, numWorkers)

func main() {
	jobSlice, _ := getAllJobItems()
	go queueJobs(jobSlice)
	errors := make(chan error)
	go receiveResults(done)
	var w sync.WaitGroup
	w.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, &w)
	}
	w.Wait()
	close(results)
	<-errors
}
func worker(i int, wg *sync.WaitGroup) {
	for jobItem := range jobCh {

        select {
        case <-errCh:
            fmt.Println("received value: ", v)
            wg.Done()
            return
        default:
            output, err := Result{jobItem, expensiveFunc(jobItem)}
            resultCh <- output
        }

	}
	wg.Done()
}
func queueJobs(jobSlice []string) {
	for _, p := range jobSlice {
		jobCh <- p
	}
	close(workCh)
}
func receiveResults(done chan bool) {
	for result := range results {

        if result.err != nil {
	        done <- err
            return
        }

		fmt.Printf("Job id %d, result %v\n", result.job.id, result)
	}
	done <- nil
}
