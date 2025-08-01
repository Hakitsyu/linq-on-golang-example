package linq

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
