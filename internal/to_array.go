package linq

func ToArray[K interface{}](source Enumerable[K]) []K {
	enumerator := source.GetEnumerator()

	result := make([]K, 0)
	for enumerator.MoveNext() {
		result = append(result, enumerator.Current())
	}

	return result
}
