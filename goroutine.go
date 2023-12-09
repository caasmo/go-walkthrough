package main

import (
    "fmt"
    "time"
    "sync"
)

// https://golangbot.com/goroutines/

// Goroutines can be thought of as lightweight threads.
// cheap few kb stack size 
// multiplexed to fewer os threads
// comunicate with channels
//
// goroutine call returns immediately.

func main() {

    go hello()

    time.Sleep(100 * time.Millisecond)
    fmt.Println("main function")

    // wait group fro coordinate finichs of goroutine
    no := 3
    var wg sync.WaitGroup
    for i := 0; i < no; i++ {
        wg.Add(1)
        go process(i, &wg)
    }
    wg.Wait()
    fmt.Println("All go routines finished executing")
}

func hello() {
    fmt.Println("goroutine: Hello world")
}

func process(i int, wg *sync.WaitGroup) {
    fmt.Println("started Goroutine ", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d ended\n", i)
    wg.Done()
}

