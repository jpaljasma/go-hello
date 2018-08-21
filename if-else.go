package main

import "fmt"

// alias function - print formatted, followed by a line feed
func pfnl(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Println()
}

func main() {

	var d uint16 = 7
	pnl := fmt.Println

	// https://en.wikipedia.org/wiki/Yoda_conditions
	if 0 == d%2 {
		pfnl("%d is odd", d)
	} else {
		pfnl("%d is even", d)
	}

	if 8%4 == 0 {
		pfnl("8 is divisible by 4")
	}

	// A statement can precede conditionals;
	// any variables declared in this statement are available in all branches.
	if num := 9; num < 0 {
		pnl(num, "is negative")
	} else if num < 10 {
		pnl(num, "has 1 digit")
	} else {
		pnl(num, "has multiple digits")
	}
}
