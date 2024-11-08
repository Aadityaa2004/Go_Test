package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p1  = &Vertex{1, 2} // has type *Vertex
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)


	// ========================================
	// Defer
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// ========================================
	// Pointers
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// ========================================
	// Arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// ========================================
	// Structs
	fmt.Println(v1, p1, v2, v3)

	// ========================================
	// Slices
	var s []int = primes[1:4]
	fmt.Println(s)

	// ========================================
	// Slices are like references to arrays
	var s12 []int
	fmt.Printf("len=%d cap=%d %v\n", len(s12), cap(s12), s12)

	// append works on nil slices.
	s12 = append(s, 0)
	fmt.Printf("len=%d cap=%d %v\n", len(s12), cap(s12), s12)

}

