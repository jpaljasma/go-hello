package main

import (
	"fmt"
	"strings"
)

// function alias
var out = fmt.Println

func main() {
	var vTest = false
	stringsSlice := []string{"go", "version", "go1.10.3 darwin/amd64"}

	out("go" + " version " + stringsSlice[2])
	out(strings.Join(stringsSlice[:], ", "))

	// have compiler guess the length of the the spaceDogs array
	spaceDogs := [...]string{"Лайка", "Белка", "Стрелка", "Пушинка", "Пчёлка", "Мушка", "Чернушка"}
	out(spaceDogs)
	fmt.Printf("%s was a Soviet space dog\n", spaceDogs[0])
	out()

	// First 3 space dogs sliced
	firstThreeDogs := spaceDogs[:3]
	out(firstThreeDogs)
	out(spaceDogs[3:]) //	slice from index 3
	out()

	// Slicing DOES NOT COPY the slice's data.
	spaceDogs[1] = `Бобик`
	out(firstThreeDogs) // 2nd element is Бобик
	firstThreeDogs[1] = `Чернушка`
	out(firstThreeDogs)
	out(spaceDogs) // 2nd element is now Чернушка

	out()

	out("1 + 1 =", 1+1)
	fmt.Printf("7.0 / 3.0 = %0.6f\n", 7.0/3.0)
	out("true && false", true && false)
	out("true || false =", true || false)
	out("!true =", !true)
	out("false == vTest", false == vTest)
}
