---
jupytext:
  text_representation:
    extension: .md
    format_name: myst
    format_version: 0.13
    jupytext_version: 1.14.4
kernelspec:
  display_name: Go
  language: go
  name: gophernotes
---

# Channels

+++

Channels can be thought of as pipes used by Goroutines to communicate

+++

Notes from [Go by Example](https://golangbot.com/channels/).

```{code-cell}
:tags: [hide-output, hide-input]

import (
    "fmt"
    "time"
    "strings"
)
```

# The type of a channel

+++

Each channel has a type associated with it.

`chan T` is the type of a channel of type `T`

```{code-cell}
:tags: [hide-output]

var a chan int // This is the declaration
fmt.Printf("Type of the channel: %T\n", a)
```

# Zero value of a channel

The zero value of a channel is nil.

```{code-cell}
:tags: [hide-output]

fmt.Println("zero value of channel a:", a)
```

# Declare a Channel with `make`

With the short hand declaration

```{code-cell}
b := make(chan int)
fmt.Printf("Type of b is %T\n", b)
```

# Read from a Channel (Syntax)

The arrow points outwards from a and hence we are reading from channel a:

```go
data := <- a // read from channel a  
```

+++

# Send to a Channel (Syntax)

The arrow points towards a and hence we are writing to channel a:

```go
a <- data // write to channel a  
```

+++

# Send and receive are blocking by default

When data is sent to a channel, the control is blocked in the send statement
until some other Goroutine reads from that channel. Similarly, when data is
read from a channel, the read is blocked until some Goroutine writes data to
that channel.

+++

# Unidirectional Channels

the type for a send only channel is `chan<- int`

```{code-cell}
var ch chan<- int
fmt.Printf("Type of the channel: %T\n", ch)
```

And the type for a receive only channel is `<-chan int`

```{code-cell}
var ch2 <-chan int
fmt.Printf("Type of the channel: %T\n", ch2)
```

 It is possible to convert a bidirectional channel to a send only or receive
 only channel but not the vice versa.

+++

# Channel close: `ok` value

Receivers can use an additional variable `v, ok := <- ch` while receiving data
from the channel to check whether the channel has been closed. 

If `ok` is false it means that we are reading from a closed channel. The value
read from a closed channel will be the zero value of the channel's type:

```{code-cell}
ch3 := make(chan int) 
func producer(ch chan int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}

go producer(ch3)
for {
    v, ok := <-ch3
    if ok == false {
        fmt.Println("ok is now:", ok)
        break
    }
    fmt.Println("Received ok value", v, ok)
}
```

# Channel close: `for range` loop

You can iterate a channel to receive values from it:
until it is closed.

```{code-cell}
ch4 := make(chan int)

go producer(ch4)
for v := range ch4 {
    fmt.Println("for range: Received in ok value", v)
}
```

 Only the sender should close a channel, never the receiver. 
 Sending on a closed channel will cause a panic.

# Buffered Channel

Channels can be buffered. Provide the buffer length as the second argument to
make to initialize a buffered channel:

```
ch := make(chan int, 100)
```

# Deadlock

Sends to a buffered channel block only when the buffer is full

```{code-cell}
ch := make(chan string, 2)
ch <- "naveen"
ch <- "paul"
ch <- "steve"
fmt.Println(<-ch)
fmt.Println(<-ch)
```

Receives block when the buffer is empty.

# Pattern: completion (done) channel

- pass the completion (done) channel to goroutine in the signature
- let goroutine do long thing and inform
- the channel in the main routine informs of completion while blocking.

```{code-cell}
func busyCodeFunc(done chan bool) {
	time.Sleep(2000 * time.Millisecond)
	done <- true 
}

done := make(chan bool)
go busyCodeFunc(done)
fmt.Println("Waiting for busyCodeFunc complation")
<-done // not use or store that data in any variable. This is perfectly legal.
fmt.Println("Yo")
```

# Pattern: Result channel

- pass the job and a result channel to the goroutine
- let goroutine do heavy work
- inform to caller of the result

```{code-cell}
func calcWork1(payload int, ch chan int) {

	time.Sleep(400 * time.Millisecond)
	ch <- payload + 42
}

func calcWork2(payload int, ch chan int) {

	time.Sleep(700 * time.Millisecond)
	ch <- payload + 42
}

work1payload := 100
work2payload := 200
chForWork1 := make(chan int)
chForWork2 := make(chan int)

go calcWork1(work1payload, chForWork1)
go calcWork2(work2payload, chForWork2)
result1, result2 := <-chForWork1, <-chForWork2 // wait for the two
fmt.Println("Final output", result1, result2, result1+result2)
```

```{code-cell}

```
