package main

import "fmt"

type Enumerator[K interface{}] interface {
	Reset()
	MoveNext() bool
	Current() K
}

type Enumerable[K interface{}] interface {
	GetEnumerator() Enumerator[K]
}

type ArrayEnumerator[K interface{}] struct {
	Enumerator[K]
	source   []K
	position int
}

func NewArrayEnumerator[K interface{}](source []K) *ArrayEnumerator[K] {
	return &ArrayEnumerator[K]{
		position: -1,
		source:   source,
	}
}

func (enumerator *ArrayEnumerator[K]) Reset() {
	enumerator.position = -1
}

func (enumerator *ArrayEnumerator[K]) MoveNext() bool {
	nextPosition := enumerator.position + 1

	if nextPosition > (len(enumerator.source) - 1) {
		return false
	}

	enumerator.position = nextPosition
	return true
}

func (enumerator *ArrayEnumerator[K]) Current() K {
	return enumerator.source[enumerator.position]
}

type ArrayEnumerable[K interface{}] struct {
	Enumerable[K]
	source []K
}

func (enumerable *ArrayEnumerable[K]) GetEnumerator() Enumerator[K] {
	return NewArrayEnumerator(enumerable.source)
}

func NewArrayEnumerable[K interface{}](source []K) *ArrayEnumerable[K] {
	return &ArrayEnumerable[K]{
		source: source,
	}
}

type SelectFn[K interface{}, R interface{}] func(item K) R

type SelectEnumerator[K interface{}, R interface{}] struct {
	Enumerator[K]
	source Enumerator[K]
	fn     SelectFn[K, R]
}

func (enumerator *SelectEnumerator[K, R]) Reset() {
	enumerator.source.Reset()
}

func (enumerator *SelectEnumerator[K, R]) MoveNext() bool {
	return enumerator.source.MoveNext()
}

func (enumerator *SelectEnumerator[K, R]) Current() R {
	return enumerator.fn(enumerator.source.Current())
}

type SelectEnumerable[K interface{}, R interface{}] struct {
	Enumerable[R]
	source Enumerable[K]
	fn     SelectFn[K, R]
}

func (enumerable *SelectEnumerable[K, R]) GetEnumerator() Enumerator[R] {
	return &SelectEnumerator[K, R]{
		source: enumerable.source.GetEnumerator(),
		fn:     enumerable.fn,
	}
}

func Select[K interface{}, R interface{}](source Enumerable[K], fn SelectFn[K, R]) Enumerable[R] {
	return &SelectEnumerable[K, R]{
		source: source,
		fn:     fn,
	}
}

type WhereFn[K interface{}] func(item K) bool

type WhereEnumerator[K interface{}] struct {
	Enumerator[K]
	source Enumerator[K]
	fn     WhereFn[K]
}

func (enumerator *WhereEnumerator[K]) Reset() {
	enumerator.source.Reset()
}

func (enumerator *WhereEnumerator[K]) MoveNext() bool {
	for enumerator.source.MoveNext() {
		if enumerator.fn(enumerator.Current()) {
			return true
		}
	}

	return false
}

func (enumerator *WhereEnumerator[K]) Current() K {
	return enumerator.source.Current()
}

type WhereEnumerable[K interface{}] struct {
	Enumerable[K]
	source Enumerable[K]
	fn     WhereFn[K]
}

func (enumerable *WhereEnumerable[K]) GetEnumerator() Enumerator[K] {
	return &WhereEnumerator[K]{
		source: enumerable.source.GetEnumerator(),
		fn:     enumerable.fn,
	}
}

func Where[K interface{}](source Enumerable[K], fn WhereFn[K]) Enumerable[K] {
	return &WhereEnumerable[K]{
		source: source,
		fn:     fn,
	}
}

func main() {
	enumerable := NewArrayEnumerable([]int{0, 1, 2, 3, 4, 5})
	enumerator := enumerable.GetEnumerator()

	for enumerator.MoveNext() {
		fmt.Printf("%v; ", enumerator.Current())
	}

	println("")

	enumerableMultipliedBy10 := Select(enumerable, func(value int) int {
		return value * 10
	})
	enumeratorMultipliedBy10 := enumerableMultipliedBy10.GetEnumerator()
	for enumeratorMultipliedBy10.MoveNext() {
		fmt.Printf("%v; ", enumeratorMultipliedBy10.Current())
	}

	println("")

	enumerableGreaterThan30 := Where(enumerableMultipliedBy10, func(value int) bool {
		return value > 30
	})
	enumeratorGreaterThan30 := enumerableGreaterThan30.GetEnumerator()
	for enumeratorGreaterThan30.MoveNext() {
		fmt.Printf("%v; ", enumeratorGreaterThan30.Current())
	}

	println("")
}
