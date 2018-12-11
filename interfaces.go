package main

import (
	"fmt"
	"sort"
)

type Content interface {
	Render() string
}

type ContentPage struct {
}

type SpecialsPage struct {
}

func (c ContentPage) Render() string {
	return "Content Page render"
}

func (s SpecialsPage) Render() string {
	return "Special Offers Page render"
}

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

func main() {
	animals := []Content{ContentPage{}, SpecialsPage{}}
	fpl := fmt.Println

	si := Sequence{9, 2, 3, 4, 0}

	fpl(si.Less(0, 1)) // false
	si.Swap(0, 1)
	fpl(si.Less(0, 1)) // true
	fpl(si)            // calls .String() internally
	// same
	fpl(si.String())

	for _, page := range animals {
		fpl(page.Render())
	}

	// Todo: https://golang.org/doc/effective_go.html#interfaces
}
