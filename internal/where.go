package linq

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
