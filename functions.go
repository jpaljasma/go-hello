package main

import (
	"fmt"
	"math"
	"math/big"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func mean(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func nameTheReturnType(x float64) (n float64) {
	n = math.Pow(x, 3.0)
	return // no need to call 'return n'
}

// function may return multiple values
func returningMultipleValues() (int, int) {
	return 5, 6
}

// There is a special form available for the last parameter in a Go function:
// By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters.
func variadicSum(args ...float64) float64 {
	total := 0.0

	for _, n := range args {
		total += n
	}

	return total
}

func nextOddGenerator() func() uint {
	j := uint(1)

	return func() uint {
		ret := j
		j += 2
		return ret
	}
}

func factorial(x uint) uint {

	if 0 == x {
		return 1
	}

	return x * factorial(x-1)
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// recursive fibonacci
func fib(n uint) uint {

	if 0 == n || 1 == n {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// linear fibonacci implementation O(n)
func fibLinear(n uint) uint {
	fn := make(map[uint]uint)

	for i := uint(0); i <= n; i++ {
		var f uint

		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}

	return fn[n]
}

func fibBig(n int) *big.Int {

	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 0; i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return a
}

func main() {

	p := fmt.Println

	res := plus(1, 2)
	p("1 + 2 =", res)

	p("1 + 2 + 3 =", plusPlus(1, 2, 3))

	xs := []float64{98, 93, 77, 82, 83}
	p("Mean", mean(xs))
	p("Sum", variadicSum(xs...)) // following slice with ... allows us passing slice to a variadic function
	p("Pow3", nameTheReturnType(3))

	five, six := returningMultipleValues()
	p(five, six)

	// closure

	add := func(x, y int) int {
		return x + y
	}
	p("Closure", add(1, 2))

	// closure with local scope
	x := 0
	increment := func() int {
		x++
		return x
	}

	p("Inc", increment())
	p("Inc", increment())

	// closure can also return a function
	nextEven := makeEvenGenerator()
	nextOdd := nextOddGenerator()

	p(nextEven())
	p(nextEven())
	p(nextEven())

	p(nextOdd())
	p(nextOdd())
	p(nextOdd())

	p("Factorial 4 =", factorial(4))

	p(fib(14))
	p(fibLinear(14))
	p(fibBig(14))

	// panic & recover

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

	fmt.Println("Continuing")

}
