package main

import (
	"fmt"
	"math"
)

type PointOnScreen struct {
	pt    Vertex
	label string
}

type Vertex struct {
	X int
	Y int
}

type Circle struct {
	x, y int
	r    float64
}

type Rectangle struct {
	width, height float64
}

// Go supports methods defined on struct types.
// access this function like {Circle}.area()

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c *Circle) circ() float64 {
	return 2 * math.Pi * c.r
}

func (c *Circle) diam() float64 {
	return 2 * c.r
}

func (c *Circle) perim() float64 {
	return c.diam()
}

func (r *Rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

// interfaces are named collections of method signatures

type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func MeasureAll(vals []geometry) {
	for _, val := range vals {
		measure(val)
	}
}

func main() {

	var (
		pzv = &Vertex{1, 2} // has type *Vertex
		// pointer to struct
		pv *Vertex
	)

	// alias
	p := fmt.Println

	p(Vertex{10, 20})
	p(Vertex{Y: 20})

	v := Vertex{1, 2}
	pv = &v
	pv.Y = 29

	pzv.Y = 1e2
	p(v, pzv)

	pos := PointOnScreen{*pzv, "My point"}
	p(pos)

	c := Circle{1, 2, 3}
	p("Diam", c.diam())
	p("Circ", c.circ())
	p("Area", c.area())

	gr := Rectangle{width: 3, height: 4}
	gc := Circle{r: 5}

	MeasureAll([]geometry{&gr, &gc})

	// @todo http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
}
