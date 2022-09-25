package main

import (
    "fmt"
    "strings"
)

func main() {
    // A pointer holds the memory address of a value.

	header("Declaration Of Pointer")
    var p *int
	fmt.Printf("Type of the pointer: %T\n", p)

	header("Zero Value")
    // zero value of a poiter is nil
    fmt.Println("zero valueof pointer is ", p) // <nil>

	header("& Operator")
    // operator &  generates a pointer to its operand.
    i := 42
    var pi *int
    pi = &i
    fmt.Println("Generating pointer", pi) // 0xc0000ba008 

	header("Dereferencing Operator *")
    // dereferencing: operator * gives underlying value of a pointer
    fmt.Println("dereferencing", *pi) // 42

	header("Set underlying value of pointer")
    *pi = 21
    fmt.Println(*pi) // 21

	header("Null Pointer Inside Interface")

    fi := func(ptr *int) interface{} {
        return ptr
    }
    fmt.Println("Is nil?: ", fi(nil) == nil)
	fmt.Printf("underlying Type is : %T\n", fi(nil))
	fmt.Printf("underlying value is : %v\n", fi(nil))

	header("Pointer as function argument")
    a := 5
    fWithPointer := func(a *int) {
	    fmt.Printf("Value is: %v\n", *a)
    }
    fWithPointer(&a)

	header("Pointer as function return type")
    f := func() *int {  // specify return type as pointer
        v := 101
        // return the address
        return &v
    }

    n := f()
    fmt.Println("func returns pointer",  n)  // 0x40e020
    fmt.Println("func returns value",  *n)  // 101

	header("New generates apointer")
    ptri := new(int)
    *ptri = 67
    fmt.Println("Pinter is ", ptri)  // 0x40e020
    fmt.Println("Value is ",*ptri) // 67

	header("No need a pointer to function")
    f2 := func() {
        fmt.Println("a function")
    }

	fmt.Printf("Type of the function: %T\n", f2)
    // pf := &f NO!!!!
    pf := f2
    fmt.Println("Value of a Type function is the address of the function", pf) // the value is adress of fucntion
    pf()
}

func header(h string) {
    fmt.Println()
    fmt.Println(h)
    fmt.Println(strings.Repeat("-", len(h)))
}
