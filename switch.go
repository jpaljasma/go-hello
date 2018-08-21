package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"
)

func main() {
	fp := fmt.Print
	fpl := fmt.Println

	i := 2
	fp("Write ", i, " as")

	switch i {
	case 1:
		fpl(" one")
	case 2:
		fpl(" two")
	case 3:
		fpl(" three")
	default:
		fpl(" other")
	}

	s := "red herring"

	// last 3 characters of a string
	switch s[len(s)-3:] {
	case "ing":
		fpl("Ends with ING")
	default:
		fpl("Something else")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fpl("It's the weekend")
	default:
		fpl("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fpl("It's before noon")
	default:
		fpl("It's afternoon")
	}

	fpl()

	// type check
	detectType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fpl("To be or not to be. I'm a Boolean")
		case int:
			fpl("I'm an Int")
		case float64:
			fpl("I'm Floating")
		case nil:
			fpl("Looks like a <nil>, walks like a <nil>")
		default:
			//fmt.Printf("Don't recognize the type %T\n", t)
			fpl(fmt.Sprintf(`Don't recognize the type "%T"`, t))
		}
	}

	detectType(true)
	detectType(1)
	detectType(nil)
	detectType(1.234)
	detectType(math.Pi)
	detectType("hey")

	fpl()

	// You can also declare a variable and it will have the type of each case

	boom := func(v interface{}) string {
		switch u := v.(type) {
		case bool:
			return strconv.FormatBool(!u) // u has a type boolean
		case int:
			return strconv.Itoa(u * 2) // u has a type int
		case string:
			mid := len(u) / 2
			return u[mid:] + u[:mid]
		}
		return "unknown"
	}

	fpl(boom(true))
	fpl(boom(21))
	fpl(boom("bitrab"))
	fpl(boom("discovery"))

	fpl()

	// checking multiple string options + fallthrough
	switch runtime.GOOS {
	case "darwin",
		"windows":
		fpl("This is either a OSX or Windows operating system")
		fallthrough
	default:
		fpl(fmt.Sprintf("Detected OS name: %s", runtime.GOOS))
	}

	// A switch with no value means "switch true", making it a cleaner version of an if-else chain,
	// as in this example from Effective Go

	var b byte = 0xd
	fpl(b)
	fpl(unhex(b))
	fpl(unhex('B'))

	fpl()

	// noop case
	fpl(pluralize("foo", 1))
	fpl(pluralize("bar", 2))
	fpl(pluralize("man", 3))
	fpl(pluralize("car", 0))
}

func pluralize(s string, n int) string {
	ending := ""

	switch n {
	case 1: // in PHP, this line will do fallthrough. in Go, this is noop()
	default:
		ending = "s"
	}

	return strconv.Itoa(n) + " " + s + ending
}

func unhex(c byte) byte {
	// cleaner version of if then else
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
