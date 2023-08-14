package fx

type KeySelector[T any, K comparable] func(T) K

// SliceGroupBy turns input into a map using result of key selector as index.
// Items with the same key are grouped together as slice under the same key
func SliceGroupBy[T any, K comparable](input []T, keySelector KeySelector[T, K]) map[K][]T {
	res := map[K][]T{}
	for _, v := range input {
		key := keySelector(v)
		if values, found := res[key]; found {
			res[key] = append(values, v)
		} else {
			res[key] = []T{v}
		}
	}
	return res
}
