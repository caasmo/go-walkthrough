package main

import (
    "fmt"
    "strings"
)

// https://golangbot.com/structs/

func main() {

	header("Struct Declaration")
    type Vertex struct {
        X int
        Y int
    }
    fmt.Println(Vertex{1, 2})

    // init and access field with a dot

	header("Struct Initialization, assigment")
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)

	header("Go follows the pointer to struct")
    // if p is pointer to struct, go allows p.X instead of (*p).X
    v1 := Vertex{1, 2}
    pv1 := &v1
    fmt.Println(pv1.X) // 1

	header("Struct Initialization with name")
    e := Vertex{
        X: 3,
        Y: 25,
    }
    fmt.Println(e)

	header("Struct Initialization without name")
    w := Vertex{1, 2}
    fmt.Println(w)

	header("Struct Initialization anonymous")
    g := struct {
        firstName string
        age       int
    }{
        firstName: "Andreah",
        age:       31,
    }

    fmt.Println(g)

	header("Zero field value")
    // Zero field value is the zero of the type of the value
    z1 := Vertex{}
    fmt.Println("zero value of unitializsed struct", z1.X) // 0
    // the same with var
    var z2 Vertex
    fmt.Println("zero value of unitializsed struct", z2.X) // 0

	header("Check for a empty *unitialized struct")
    a1 := Vertex{}
    if (Vertex{}) == a1  {
        fmt.Println("check if empty struct:", "Yes")
    }

	header("Promoted fields")
    type Address struct {
        city string
        state string
    }

    type Person struct {
        name string
        age  int
        Address
    }
    p := Person{
        name: "Naveen",
        age:  50,
        Address: Address{
            city:  "Chicago",
            state: "Illinois",
        },
    }

    fmt.Println("Promoted field: State:", p.state) //state is promoted field
}

func header(h string) {
    fmt.Println()
    fmt.Println(h)
    fmt.Println(strings.Repeat("-", len(h)))
}
