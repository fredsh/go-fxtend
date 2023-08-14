package fx

func SliceConcat[T any](inputs ...[]T) []T {
	totalLength := 0
	for _, input := range inputs {
		totalLength += len(input)
	}

	result := make([]T, totalLength)
	index := 0

	for _, input := range inputs {
		copy(result[index:], input)
		index += len(input)
	}

	return result
}
