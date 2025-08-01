package linq

type FirstFn[K interface{}] func(item K) bool

type FirstEnumerator[K interface{}] struct {
	Enumerator[K]
	source Enumerator[K]
	fn     FirstFn[K]
}

func (enumerator *FirstEnumerator[K]) Reset() {
	enumerator.source.Reset()
}

func (enumerator *FirstEnumerator[K]) MoveNext() bool {
	for enumerator.source.MoveNext() {
		if enumerator.fn(enumerator.Current()) {
			return true
		}
	}

	return false
}

func (enumerator *FirstEnumerator[K]) Current() K {
	return enumerator.source.Current()
}

type FirstEnumerable[K interface{}] struct {
	Enumerable[K]
	source Enumerable[K]
	fn     FirstFn[K]
}

func (enumerable *FirstEnumerable[K]) GetEnumerator() Enumerator[K] {
	return &FirstEnumerator[K]{
		source: enumerable.source.GetEnumerator(),
		fn:     enumerable.fn,
	}
}

func First[K interface{}](source Enumerable[K], fn FirstFn[K]) Enumerable[K] {
	return &FirstEnumerable[K]{
		source: source,
		fn:     fn,
	}
}
