package main

import (
    "fmt"
    "strings"
)

//https://golangbot.com/interfaces-part-1/
//https://golangbot.com/interfaces-part-2/

func main() {

	// interface is a set of methods signatures

	header("Interface Declaration")

	name := MyString("hello")

	var v StrFinder
	v = name
	// Internal representation tuple(type,value)
	fmt.Printf("Type = %T, value = %v\n", v, v)


	header("empty interface")
	s := "Hello World"
	var i interface{}
	i = s
	fmt.Printf("Type = %T, value = %v\n", i, i)

	header("Type assertion")
	var i2 interface{} = 56
	x := i2.(int)
	fmt.Println("type assertion for interface", x)

	header("check type of interface")
	var i3 interface{} = 56
	v3, ok3 := i3.(int)
	fmt.Println("is int?", v3, ok3)

	var i4 interface{} = "hello"
	v4, ok4 := i4.(int)
	fmt.Println("is int?", v4, ok4)
	fmt.Println(v4, ok4)

	header("type switch")
	var i5 interface{} = "Hello"
	switch i5.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i5.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i5.(int))
	default:
		fmt.Printf("Unknown type\n")
	}

	header("Pointer")
	// https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
	// convert, pass type to an interface
	// An assignment to a variable of interface type is valid if the value
	// being assigned implements the interface it is assigned to. It implements
	// it if its method set is a superset of the interface. The method set of
	// pointer types includes methods with both pointer and non-pointer
	// receiver. The method set of non-pointer types only includes methods with
	// non-pointer receiver.

	// Pointer types has method set with both pointer and non-pointer -> it can be dereferenced
	// NOn-Pointer types has method set with oly non-pointer -> it can be

	var d2 Describer
	a := Address{"Washington", "USA"}
	//d2 = a // compile error
	d2 = &a // pointer works
	fmt.Println("interface pointer receiver", d2)

	header("Zero value")
	var d3 Describer
	fmt.Println("Zero value", d3)

	header("interface value is a tuple")
	fmt.Println("type to value:", GetInterface(nil) == nil)
}

type MyString string
func (ms MyString) FindStr() string {
    return string(ms)
}

type StrFinder interface {
    FindStr() string
}


type Describer interface {
	Describe()
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func GetInterface(ptr *int) interface{} {
	return ptr
}

func header(h string) {
    fmt.Println()
    fmt.Println(h)
    fmt.Println(strings.Repeat("-", len(h)))
}
