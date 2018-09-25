package main

import "fmt"

type Page struct {
	id    int
	title string
}

func main() {

	i, j := 10, 4181

	var p *int
	p = &i          // point to i
	fmt.Println(*p) // read i through the pointer p

	*p = 33 // set i value through pointer p

	fmt.Println(i) // print new i value (which is 33)

	p = &j       // now move pointer to j
	*p = *p / 37 // divide j through the pointer

	fmt.Println(*p)

	p = nil // kill the pointer

	fmt.Println(i, j)

	stringsSlice := []string{"go", "version", "go1.10.3 darwin/amd64"}

	p2 := &stringsSlice

	*p2 = append(*p2, "intel") // append an element to stringSlice through the pointer
	fmt.Println(*p2)

	var n = 0
	p = &n

	// pointer and loop
	for n = 0; n < 10; n++ {
		fmt.Print(*p, " ")
	}
	fmt.Println()

	g := Page{title: "Title", id: 121}

	pg := &g // point to a Page type
	// set pointer to an id element of a page
	px := &pg.id

	// page id set to 199
	*px = 199
	pg.title = "Programming in GO"

	// print the same results
	fmt.Println(g)
	fmt.Println(*pg)

	pg = nil
	px = nil
}
