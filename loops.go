package main

import "fmt"

func p(s ...interface{}) {
	fmt.Print(s, " ")
}

/*
	for is Go’s only looping construct.
*/

func main() {
	pnl := fmt.Println
	// p := fmt.Print

	var i int32

	pnl(`LOOPS`)

	i = 0
	for i < 3 {
		p(i)
		i++
	}
	pnl()

	for j := 0; j < 10; j++ {
		p(j)
	}
	pnl()

	for k := 0; k < 16; k += 3 {

		// breaking the loop
		if 0 == k%12 && k > 0 {
			p("Breaking at", k)
			break
		} else
		// skipping in the loop
		if 0 == k%2 {
			p("Skipping", k)
			continue
		}

		p(k)
	}

	pnl()

	i = 2

	// empty for == while
	for {
		p(i)

		if i >= 256 {
			break
		}
		i *= 2
	}

	pnl()

	for m := 9; m >= 0; m-- {
		p(m)
	}

	pnl()
}
