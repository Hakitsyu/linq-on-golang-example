package linq

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
