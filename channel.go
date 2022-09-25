package main

import (
	"fmt"
	"time"
    "strings"

)

// https://golangbot.com/channels/

// - Channels can be thought of as pipes used by Goroutines to communicate
// - Each channel has a type associated with it
// - Receive read  from  to channel: data := <- ch
// - Send, write to channel: ch <- data
// - send and receive block by default, no lock
// - posible to convert bidirectional channel to a send or receive only channel,
// - not vice versa.
// - buffered channels: It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer is full.
// - buffered channels: receives only block when empty 
// - buffered channels: It is possible to read data from closed channels. they return zeo value and ok false  n, open = <-ch  or  for range loop
// - buffered channels: lenght is how many items the channel channels
// - waitgrop to wait all gorotines to finish

func main() {

	header("Declaration Of Channel")
	// [chan T] is the type of a channel of type T
	var a chan int
	fmt.Printf("Type of the channel: %T\n", a)

	header("Zero value")
	// Zero value of a channel is nil
	fmt.Println("zero value channel:", a)

	header("Declare channel with make")
	b := make(chan int)
	fmt.Printf("Type of b is %T\n", b)

	header("Done Channel")
	// - pass done channel to goroutine,
	// - let goroutine do long thing and inform
	// - inform of ready, but no return the work
	done := make(chan bool)
	fmt.Println("Create go goroutine")
	go hello(done)
	<-done
	fmt.Println("Main received done channel value")

	header("Result Channel")
	// - pass channel for work result
	// - let goroutine do work
	// - inform to caller of the result
	work1payload := 100
	work2payload := 200
	chForWork1 := make(chan int)
	chForWork2 := make(chan int)

	go calcWork1(work1payload, chForWork1)
	go calcWork2(work2payload, chForWork2)
	result1, result2 := <-chForWork1, <-chForWork2
	fmt.Println("Final output", result1, result2, result1+result2)

	header("Unidirectional Channels")
    // the  type is for a send only channel is chan<- int 
    // only send channel
    ch := make(chan<- int)
	fmt.Printf("Type of the channel: %T\n", ch)

    // only receive channel
    ch2 := make(<-chan int)
	fmt.Printf("Type of the channel: %T\n", ch2)

	header("Receive OK Value")
    ch3 := make(chan int)
    go producer(ch3)
    for {
        v, ok := <-ch3
        if ok == false {
            fmt.Println("ok is now:", ok)
            break
        }
        fmt.Println("Received ok value", v, ok)
    }

	header("for range loop until channel closed")
    ch4 := make(chan int)
    go producer(ch4)

    for v := range ch4 {
        fmt.Println("for range: Received in ok value", v)
    }

	header("Buffered channel has capacity")
    ch5 := make(chan int, 2)
	fmt.Printf("Type of the channel: %T\n", ch5)

	header("Buffered channel lenght")
    ch6 := make(chan string, 3)
    ch6 <- "naveen"
    ch6 <- "paul"
    fmt.Println("capacity is", cap(ch6))
    fmt.Println("length is", len(ch6))
    fmt.Println("read value", <-ch6)
    fmt.Println("new length is", len(ch6))

}

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func calcWork1(payload int, ch chan int) {

	time.Sleep(400 * time.Millisecond)
	ch <- payload + 42
}

func calcWork2(payload int, ch chan int) {

	time.Sleep(700 * time.Millisecond)
	ch <- payload + 42
}

func producer(chnl chan int) {
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)
}

func header(h string) {
    fmt.Println()
    fmt.Println(h)
    fmt.Println(strings.Repeat("-", len(h)))
}
