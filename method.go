package main

import "fmt"

// https://golangbot.com/methods/
func main() {

    // method declaration  displaySalary()
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
	}
	emp1.displaySalary()

    //
    // Pointer Receivers vs Value Receivers
    //
    // change made inside a method with a pointer receiver is visible to the caller
    // whereas this is not the case in value receiver.
    //
    // Pointers receivers can also be used in places where it's expensive to copy a data structure.
    //
    // pointer receiver 
    e := Employee{
        name: "lipo",
        salary:  50,
    }

    fmt.Printf("Employee name before change: %s\n", e.name)
    e.changeName("Michael Andrew")
    fmt.Printf("Employee name after change: %s\n", e.name)

    //
    // method with a value receiver accepts value and pointer receiver
    //
    // It is legal to call a value method on anything which is a value or whose value can be dereferenced.
    e1 := Employee{
        name: "lipo",
        salary:  500,
    }

    p := &e1
    p.displaySalary() // is interpreted as (*p).displaySalary()

    //
    // method with a pointer receiver accepts value and pointer receiver
    //
    e2 := Employee{
        name: "lipo",
        salary:  500,
    }

    e2.changeName("Michael Andrew")
    fmt.Printf("Employee name after change with value: %s\n", e2.name)
    p2 := &e2
    p2.changeName("lipo")
    fmt.Printf("Employee name after change: %s\n", p2.name)

}

type Employee struct {
	name     string
	salary   int
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %d\n", e.name, e.salary)
}

// value receiver
func (e *Employee) changeName(newName string) {
    e.name = newName
}
