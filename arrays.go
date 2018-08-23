package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"sort"
	"time"
)

// In Go, an array is a numbered sequence of elements of a specific length.
func main() {

	// Empty array with length of 5
	// By default an array is zero-valued, which for ints means 0s.
	var a [5]int

	p := fmt.Println
	lpf := log.Printf

	// type is array
	p("typ: ", reflect.TypeOf(a).Kind())

	p("emp: ", a)
	p("len: ", len(a))

	a[4] = 100
	p("set:", a)
	p("get:", a[4])

	// declare and initialize an array
	b := [5]int{1, 2, 3, 4, 5}
	p("dcl:", b)

	// 2d array
	twoD := [2][3]int{{0, 1, 2}, {1, 2, 3}}
	p("2d:", twoD)
	p()

	// 3d array 3x3x3
	var threeD [3][3][3]int

	// direct set
	threeD[0][0][1] = 1
	threeD[1][2][0] = 2
	threeD[2][0][2] = 3

	p("3d:", threeD)

	c := 1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				c *= 2
				threeD[i][j][k] = c
			}
		}
	}

	p("3d:", threeD)

	// SLICES
	// Go arrays are fixed in size, but thanks to the builtin append method, we get dynamic behavior.
	p()

	x := []int{0, 1, 2, 3, 4}
	x = append(x, 5)

	// type is slice
	p("typ: ", reflect.TypeOf(x).Kind())

	p(x, "len:", len(x), "cap:", cap(x))

	// makes an empty slice of 4 elements with 0 as a value
	x = make([]int, 4)
	p(x, "len:", len(x), "cap:", cap(x))

	// create an empty slice with 0 length and capacity of 4
	x = make([]int, 0, 4)
	// empty slice, length = 0 cap = 4
	p(x, "len:", len(x), "cap:", cap(x))

	// append 5 elements, length = 5 cap = 8 (doubled)
	x = append(x, 5, 6, 7, 8, 9)
	p(x, "len:", len(x), "cap:", cap(x))
	p()

	// calculating primes
	nPrimeLoop := 10000

	start := time.Now()
	for i := 1; i <= nPrimeLoop; i++ {
		if IsPrime(i) {
			//fmt.Printf("%v ", i)
		}
	}
	elapsed := time.Since(start)
	lpf("Primes (linear) took %s", elapsed)

	start = time.Now()
	for i := 1; i <= nPrimeLoop; i++ {
		if IsPrimeSqrt(i) {
			//fmt.Printf("%v ", i)
		}
	}
	p()
	elapsed = time.Since(start)
	lpf("Primes (sqrt) took %s", elapsed)

	start = time.Now()
	primes := SieveOfEratosthenes(nPrimeLoop)
	elapsed = time.Since(start)
	lpf("Primes (SieveOfEratosthenes) took %s", elapsed)

	p(len(primes), "primes found")

	// underscore (_) means yea, it's a variable but I dont need it so ignore it
	for _, v := range primes {
		fmt.Print(v, " ")
	}

	p()

	a2 := []float64{1.99, 1, 2, 7, 3, -2, 4, 5, 6.789}
	a3 := []float64{5, 11, 12, 4, 8, 21}
	a4 := []float64{9, 13, 9, 11, 9, 13, 11, 9, 10, 8, 11}
	a5 := []float64{math.Phi}

	p("array sum: ", arraySum(a2))
	p("array mean: ", arrayAvg(a2))
	p("array geometric mean: ", geometricMean(a3))

	p("median", a3, arrayMedian(a3))
	p("median", a4, arrayMedian(a4))
	p("median", a5, arrayMedian(a5))

	// Maps are like associative arrays in PHP but strong typed
	z := map[string]float64{
		"fish":    23.985,
		"chips":   11.0,
		"burgers": 8.95,
	}

	for key, value := range z {
		p(fmt.Sprintf("%s costs $%.2f", key, value))
	}

	// Iteration order: keys are not guaranteed to be sorted
	// we would have to create a separate sorted slice of keys
	var keys = make([]string, len(z), len(z))
	i := 0
	for k := range z {
		keys[i] = k
		i++
	}
	p("orig:  ", keys)
	sort.Strings(keys)
	p("sorted:", keys)
	// <T>sAreSorted()
	p("Keys are sorted: ", sort.StringsAreSorted(keys))

	cartTotal := 0.0
	for i := range keys {
		cartTotal += z[keys[i]]
		p(fmt.Sprintf("%s costs $%.2f", keys[i], z[keys[i]]))
	}

	p(fmt.Sprintf("Cart total: $%.2f", cartTotal))
	p()

	// deleting a key from map
	zz := z
	zz["route"] = 66

	delete(zz, "fish")
	p(zz)

	// test a key without retrieving a value
	_, ok := zz["fish"]
	if !ok {
		p("Fish key does not exist")
	}

	// if requested key does not exist, we get that value type's zero value
	jz := zz["baz"]

	p("no key:", 0 == jz)

	// len() returns number of items in map
	p("len:", len(zz))

	// Exploiting zero values
	// @see https://blog.golang.org/go-maps-in-action#TOC_4.
	// @see https://en.wikipedia.org/wiki/Doubly_linked_list
	// Doubly Linked List Node
	type Node struct {
		Next 	*Node	// A reference to the next node
		Prev	*Node	// A reference to the previous node
		Data 	interface{}		// Data or a reference to data
	}

	type DoublyLinkedList struct {
		First	*Node	// points to first node of list
		Last 	*Node	// points to last node of list
	}

	var insertAfter = func(list *DoublyLinkedList, node *Node, newNode *Node) {
		newNode.Prev = node
		if nil == node.Next {
			newNode.Next = nil
			list.Last = newNode
		} else {
			newNode.Next = node.Next
			node.Next.Prev = newNode
		}
		node.Next = newNode
	}

	var insertBefore = func(list *DoublyLinkedList, node *Node, newNode *Node) {
		newNode.Next = node
		if nil == node.Prev {
			newNode.Prev = nil
			list.First = newNode
		} else {
			newNode.Prev = node.Prev
			node.Prev.Next = newNode
		}
		node.Prev = newNode
	}

	var insertBeginning = func(list *DoublyLinkedList, newNode *Node) {
		if nil == list.First {
			list.First = newNode
			list.Last  = newNode
			newNode.Prev = nil
			newNode.Next = nil
		} else {
			insertBefore(list, list.First, newNode)
		}
	}

	var insertEnd = func(list *DoublyLinkedList, newNode *Node) {
		if nil == list.Last {
			insertBeginning(list, newNode)
		} else {
			insertAfter(list, list.Last, newNode)
		}
	}

	// Removal of a node requires special handling if the node to be removed is the firstNode or lastNode:
	var removeNode = func(list *DoublyLinkedList, node *Node) {
		if nil == node.Prev {
			list.First = node.Next
		} else {
			node.Prev.Next = node.Next
		}

		if nil == node.Next {
			list.Last = node.Prev
		} else {
			node.Next.Prev = node.Prev
		}
	}

	var list = DoublyLinkedList{}

	insertEnd(&list, &Node{ Data: "Agent" })
	insertEnd(&list, &Node{ Data: "007" })
	insertEnd(&list, &Node{ Data: "Bond" })

	// we will use remove method
	nodeToRemove := Node{ Data: "James" }
	insertBeginning(&list, &nodeToRemove)

	// TODO: https://www.geeksforgeeks.org/detect-and-remove-loop-in-a-linked-list/
	//insertEnd(&list, &nodeToRemove)

	p("First: ", list.First.Data)
	p("Last:  ", list.Last.Data)

	p()
	p("Forward")

	visited := make(map[*Node]bool)

	// Traversing forward
	node := list.First
	for nil != node {
		if visited[node] {
			p("Cycle detected:", node)
			break
		}
		visited[node] = true
		// do something with node.Data
		p("Node: ", node.Data)
		node = node.Next
	}

	p()
	p("Backwards")
	// Traversing backward
	node = list.Last
	for nil != node {
		// do something with node.Data
		p("Node: ", node.Data)
		node = node.Prev
	}

	p()
	p("Removal")
	// removal
	p("First before: ", list.First.Data)
	removeNode(&list, &nodeToRemove)
	p("First after : ", list.First.Data)

	// TODO: Go has linked list implemented https://golang.org/pkg/container/list/
}

func arrayMedian(a []float64) float64 {
	l := len(a)

	sort.Float64s(a)

	if 0 == l {
		return math.NaN()
	} else if 1 == l {
		// median value is the same as your single number
		return a[0]
	} else if 0 == l%2 {
		// for even numbers we take median of 2 middle numbers
		return arrayAvg(a[l/2-1 : l/2+1])
	} else {
		// for odd we take the midpoint
		return a[l/2]
	}
}

func geometricMean(a []float64) float64 {
	l := len(a)
	if 0 == l {
		return math.NaN()
	}
	// product of all numbers
	var p float64
	for _, n := range a {
		if 0 == p {
			p = n
		} else {
			p *= n
		}
	}

	// return geometric mean
	return math.Pow(p, 1/float64(l))
}

func arrayAvg(a []float64) float64 {
	l := len(a)
	return arraySum(a) / float64(l)
}

func arraySum(a []float64) float64 {
	s := 0.0

	for _, n := range a {
		s += n
	}
	return s
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func SieveOfEratosthenes(value int) []int {
	sqVal := int(math.Sqrt(float64(value)))
	f := make([]bool, value)

	//ret := make([]int, 0)
	// We will allocate return array cap:
	// You can Approximate pi(x) with x/(log x - 1)
	// performance improvement: 1 million - 7ms VS 4ms
	pix := int(float64(value) / (math.Log(float64(value)) - 1))
	ret := make([]int, 0, pix)

	for i := 2; i <= sqVal; i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if f[i] == false {
			ret = append(ret, i)
		}
	}

	return ret
}
