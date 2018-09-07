package main

import "fmt"

/*

range iterates over elements in a variety of data structures.

*/
func main() {

	p := fmt.Println

	nums := []int{2, 3, 4}
	var sum int

	// we do not need index, so we omit it with underscore "_"
	for _, v := range nums {
		sum += v
	}

	p("sum:", sum)

	// in case we need index
	for i, v := range nums {
		if 3 == v {
			p(fmt.Sprintf("%d found at index %d", v, i))
		}
	}

	// we can also simply get all keys
	var kz []int
	for k := range nums {
		kz = append(kz, k)
	}
	p("keys:", kz)

	// range on strings iterates over Unicode code points.
	for i, c := range "🍌🍆🍒🍑💩¯\\_(ツ)_/¯" {
		fmt.Println(i, c)
	}

}
