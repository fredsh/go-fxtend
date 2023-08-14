package fx

import (
	fxerror "github.com/fredsh/go-fxtend/pkg/fx-error"
)

// MapReverse takes a map with keys of type T and values of type U and returns a new map
// with the keys and values swapped. If there are duplicate values, the function will
// return an error.
func MapReverse[T comparable, U comparable](m map[T]U) (map[U]T, error) {
	result := make(map[U]T, len(m))
	for k, v := range m {
		if _, ok := result[v]; ok {
			return nil, fxerror.NewDuplicateValueError(v)
		}
		result[v] = k
	}
	return result, nil
}

// MapReverseOverride takes a map with keys of type T and values of type U and returns
// a new map with the keys and values swapped. If there are duplicate values, the
// function will override the previous value.
func MapReverseOverride[T comparable, U comparable](m map[T]U) map[U]T {
	result := make(map[U]T, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// MapReverseX takes a map with keys of type T and values of type U and returns
// a new map with the keys and values swapped. If there are duplicate values, the
// function will return a Result object containing an error.
func MapReverseX[T comparable, U comparable](m map[T]U) Result[map[U]T] {
	result := make(map[U]T, len(m))
	for k, v := range m {
		if _, ok := result[v]; ok {
			return NewFailure[map[U]T](fxerror.NewDuplicateValueError(v))
		}
		result[v] = k
	}
	return NewSuccess(result)
}
