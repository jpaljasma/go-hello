# Introduction to Go programming language

## Hello World

A simple introduction to Go. Creates a web server in port 8080. Does some header handling.

`go run go-hello.go`

Then open http://localhost:8080 to see the output

## 1. Values

Outputting values of different types, working with Slices, aliasing functions.

`go run values.go`

## 2. Variables

`go run variables.go`

## 3. Constants

`go run constants.go`

## 4. Loops

`go run loops.go`

## 5. If-Else

`go run if-else.go`

There is no [ternary if](https://en.wikipedia.org/wiki/%3F:) `?:` in go

## 6. Switch

`go run switch.go`

## 7. Arrays, Slices and Maps

Closer look at the **Arrays**, **Slices** and **Maps** in Golang, + work on some internals like reflecting types,
slice `cap` property & helper methods like `make` and `append`.
We also do some fun like calculating prime numbers, and some stats like mean and median.

`go run arrays.go`

### References

[Go Maps in Action](https://blog.golang.org/go-maps-in-action)

[Doubly Linked List](https://en.wikipedia.org/wiki/Doubly_linked_list)

[Package list](https://golang.org/pkg/container/list/)

## 8. Range

`go run range.go`

## 9. Functions

Regular functions, Recursive, Variadic, Closures. Fibonacci implementations (recursive & linear).
 
`go run functions.go`

## 10. Panic & Recover

`go run panic.go`

## 11. Pointers

`go run pointers.go`

## 12. Structs

`go run structs.go`

---

## Pre-commit

### Format all go files
`gofmt -w *.go`

## References

[Go By Example](https://www.gobyexample.com)

[A Tour of Go](https://tour.golang.org/welcome/1)

[Golang Tutorial Series](https://golangbot.com/learn-golang-series/)
