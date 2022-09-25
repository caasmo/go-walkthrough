package main

// https://golangbot.com/arrays-and-slices/
import "fmt"

func main() {

	header("Declaration Of Array")
    var m [4]int
    m[0] = 1
    m1 := m[0]
    fmt.Println(m1)

	header("Slice literal")
    letters := []string{"a", "b", "c", "d"}
    fmt.Println(letters)

    // slice create make
	header("Slice create with `make`")
    var s []byte
    s = make([]byte, 5, 5)
    // s == []byte{0, 0, 0, 0, 0}
    fmt.Println(s)

    // slicing
    c := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
    c1 := c[1:4] // []byte{'o', 'l', 'a'}, sharing the same storage as b
    fmt.Println(c1)

    // copy
    t := make([]byte, 5, 5)
    t1 := make([]byte, len(t)*2, (cap(t)+1)*2)
    copy(t1, t)
    fmt.Println(t1)

    //
    // append elements
    //
    a := make([]int, 1)
    a = append(a, 1, 2, 3)
    // a == []int{0, 1, 2, 3}
    fmt.Println(a)

    //
    // append two slices
    //
    b1 := []string{"John", "Paul"}
    b2 := []string{"George", "Ringo", "Pete"}
    // a == []string{"John", "Paul", "George", "Ringo", "Pete"}
    b1 = append(b1, b2...)
    fmt.Println(b1)

    // Example gist cryptowatch
    srcSlice := []int{1, 2, 3}
    anotherSlice := srcSlice[:3]
    for i := 0; i < 3; i++ {
        srcSlice[i] += 10
    }

    fmt.Println(srcSlice)
    fmt.Println(anotherSlice)
}

func header(h string) {
    fmt.Println()
    fmt.Println(h)
    fmt.Println(strings.Repeat("-", len(h)))
}
