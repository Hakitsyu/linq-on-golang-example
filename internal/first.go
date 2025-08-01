package linq

import "errors"

type FirstFn[K interface{}] func(item K) bool

var (
	ErrNoMatchFound = errors.New("no match found")
)

func First[K interface{}](source Enumerable[K]) (K, error) {
	enumerator := source.GetEnumerator()

	if enumerator.MoveNext() {
		return enumerator.Current(), nil
	}

	var zero K
	return zero, ErrNoMatchFound
}
