package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var numWorkers = runtime.NumCPU()
// the channel that describe the job. Could be a channel of structs
var jobCh = make(chan string, numWorkers)

func main() {

    // a description of the job with the 
	jobSlice, _ := getJobList()

    // job description list is put in a goroutine 
    // here we could have aldo the channel
	go queueWork(jobSlice)

	var w sync.WaitGroup
	w.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
        // the worker could also receive the channel, along with the job description
		go worker(i, &w)
	}
	w.Wait()
}

func worker(i int, wg *sync.WaitGroup) {
	for jobItem := range jobCh {
		runExpensiveFunc(jobItem)
	}
	wg.Done()
}

func queueWork(jobSlice []string) {
	for _, p := range jobSlice {
		jobCh <- p
	}
	close(jobCh)
}

func getJobList() ([]string, error) {
	s := make([]string, 100)
	for i := range s {
        s[i] = "job:" + strconv.Itoa(i)
	}

	return s, nil
}

func runExpensiveFunc(job string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Done with Job %s\n", job)
}
