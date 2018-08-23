# Introduction to Go programming language

## Hello World

A simple introduction to Go. Creates a web server in port 8080. Does some header handling.

`go run go-hello.go`

Then open http://localhost:8080 to see the output

## Values

Outputting values of different types, working with Slices, aliasing functions.

`go run values.go`

## Variables

`go run variables.go`

## Constants

`go run constants.go`

## Loops

`go run loops.go`

## If-Else

`go run if-else.go`

There is no [ternary if](https://en.wikipedia.org/wiki/%3F:) `?:` in go

## Switch

`go run switch.go`

## Arrays, Slices and Maps

Closer look at the **Arrays**, **Slices** and **Maps** in Golang, + work on some internals like reflecting types,
slice `cap` property & helper methods like `make` and `append`.
We also do some fun like calculating prime numbers, and some stats like mean and median.

`go run arrays.go`

## Pre-commit

### Format all go files
`gofmt -w *.go`

## References

[Go By Example](https://www.gobyexample.com)

[A Tour of Go](https://tour.golang.org/welcome/1)
