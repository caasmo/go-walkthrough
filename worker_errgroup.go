package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	g := new(errgroup.Group)
	g.SetLimit(runtime.NumCPU())

	jobSlice, _ := getJobList()

	for _, p := range jobSlice {
		g.Go(worker(p))
	}
	if err := g.Wait(); err != nil {
		fmt.Println("ok")
	}
	fmt.Println("not ok")
}

func worker(jobItem string) func() error {
	return func() error {
		expensiveFunc(jobItem)
		return nil
	}
}

func expensiveFunc(job string) {
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("job done %s\n", job)
}

func getJobList() ([]string, error) {
	s := make([]string, 1e5)
	for i := range s {
		s[i] = "hello" + strconv.Itoa(i)
	}

	return s, nil
}
