package main

import "fmt"

var z int                  // Package visibility
var h, j string = "h", "j" // Multi declaration and assignment

func main() {
	var a int // Declaration. The variable is initialized with the zero value of the designated type
	var b string
	var c bool
	var d float64

	fmt.Printf("a:\t%T\t%v\n", a, a) //%T prints out the type, %v prints out the value
	fmt.Printf("b:\t%T\t%v\n", b, b)
	fmt.Printf("c:\t%T\t%v\n", c, c)
	fmt.Printf("d:\t%T\t%v\n", d, d)

	e := "short hand" // Only works in a func. Shorthand declaration + initialization + assignmemt

	fmt.Printf("e:\t%T\t%v\n", e, e)
	fmt.Printf("h:\t%T\t%v\n", h, h)
	fmt.Printf("j:\t%T\t%v\n", j, j)
	fmt.Printf("z:\t%T\t%v\n", z, z)
}
