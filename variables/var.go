package main

import (
	"fmt"
)

func main() {

	var a int // Declaration of a integer varibale
    fmt.Println(a) // Value for uninitialized integer is 0

	var b string // Declaration of a string varibale
    fmt.Println(b) // Value for uninitialized string is ""

	var c byte // Declaration of a integer varibale
    fmt.Println(c) // Value for uninitialized byte is 0

	var (
		d string = "Ken Thompson"
		e int = 1
		f float64 = 2.0
		g byte = 97) 
	// Declaration and Initilization of multiple data types
    fmt.Println(d, e, f, g ) // Output: Ken Thompson 1 2 97

    fmt.Printf("%c \n", g ) // Output: a 
	// Character representation of the byte.

    var h, i int = 1, 2 // You can declare multiple variables at once.
    fmt.Println(h, i) // Output: 1 2

    var j = true // Initialization of a boolean 
    fmt.Println(j) // Output: true

    k := "Rob Pike" // The := syntax is shorthand for declaring and initializing a variable
    fmt.Println(k) // Output: Rob Pike

}