package main

import (
	"fmt"

	linq "github.com/Hakitsyu/linq-on-golang-example/internal"
)

func main() {
	enumerable := linq.NewArrayEnumerable([]int{0, 1, 2, 3, 4, 5})
	enumerator := enumerable.GetEnumerator()

	for enumerator.MoveNext() {
		fmt.Printf("%v; ", enumerator.Current())
	}

	println("")

	enumerableMultipliedBy10 := linq.Select(enumerable, func(value int) int {
		return value * 10
	})
	enumeratorMultipliedBy10 := enumerableMultipliedBy10.GetEnumerator()
	for enumeratorMultipliedBy10.MoveNext() {
		fmt.Printf("%v; ", enumeratorMultipliedBy10.Current())
	}

	println("")

	enumerableGreaterThan30 := linq.Where(enumerableMultipliedBy10, func(value int) bool {
		return value > 30
	})
	enumeratorGreaterThan30 := enumerableGreaterThan30.GetEnumerator()
	for enumeratorGreaterThan30.MoveNext() {
		fmt.Printf("%v; ", enumeratorGreaterThan30.Current())
	}

	println("")
}
