package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const s = `Constants`

// you can combine constants
const sd = s + ` (appended)`

const floatConst float64 = 1.342553235568970455640
const AppHead = "head"
const AppTail = "tail"

func main() {

	// a shorthand func
	p := fmt.Println
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// coin flip
	var n int32
	var i int

	// coin flip 10 times
	for i = 0; i < 10; i++ {
		n = r.Int31n(2)
	}

	fmt.Print(n, " ")

	if 1 == n {
		p(AppHead)
	} else {
		p(AppTail)
	}

	p(s)
	p(sd)

	const a = 500000000
	const d = 3e20 / a
	p(d)

	// A numeric constant has no type until it’s given one, such as by an explicit cast.
	p(int64(d))

	// use strconv.Itoa to print integer to string
	p(strconv.Itoa(d) + ` ` + sd)

	// formatting float constants
	p(floatConst + 1.234)
	p(strconv.FormatFloat(floatConst, 'f', 8, 64) + " formatted const")

	p()

	p(math.Cos(math.Pi))
	p(math.Sin(a))
	p(math.Sqrt(math.Pow(math.Pi, 2)))

	p()
	fmt.Printf("Now you have %g problems\n", math.Sqrt(math.Phi))

}
