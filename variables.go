package main

import (
	"fmt"
)

func main() {

	fmt.Println("VARIABLES")

	// var declares one or more variables
	var a = "initial"
	fmt.Println(a)

	// appending to a string variable
	a += " (modified)"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variable
	var d = false
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable
	f := "shorthand"
	// you could also write this `var f string = "shorthand"`
	fmt.Println(f)

	var g float64 = 1.23456789e3
	fmt.Println(g)

	// this copies the value
	h := g
	h -= 100
	fmt.Println(h, g)

}
