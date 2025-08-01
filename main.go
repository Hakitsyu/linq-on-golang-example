package main

import (
	"fmt"

	linq "github.com/Hakitsyu/linq-on-golang-example/internal"
)

func main() {
	enumerable := linq.NewArrayEnumerable([]int{0, 1, 2, 3, 4, 5})

	for _, item := range linq.ToArray(enumerable) {
		fmt.Printf("%v; ", item)
	}

	println("")

	enumerableMultipliedBy10 := linq.Select(enumerable, func(value int) int {
		return value * 10
	})

	for _, itemMultipliedBy10 := range linq.ToArray(enumerableMultipliedBy10) {
		fmt.Printf("%v; ", itemMultipliedBy10)
	}

	println("")

	enumerableGreaterThan30 := linq.Where(enumerableMultipliedBy10, func(value int) bool {
		return value > 30
	})
	for _, itemGreaterThan30 := range linq.ToArray(enumerableGreaterThan30) {
		fmt.Printf("%v; ", itemGreaterThan30)
	}

	println("")
}
