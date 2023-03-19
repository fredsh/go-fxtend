package fxcollection

import (
	"fmt"

	fxtypes "github.com/fredsh/go-fxtend/pkg/fx-types"
)

// ReverseMap takes a map with keys of type T and values of type U and returns a new map
// with the keys and values swapped. If there are duplicate values, the function will
// return an error.
func ReverseMap[T comparable, U comparable](m map[T]U) (map[U]T, error) {
	result := make(map[U]T, len(m))
	for k, v := range m {
		if _, ok := result[v]; ok {
			return nil, fmt.Errorf("duplicate value found: %v", v)
		}
		result[v] = k
	}
	return result, nil
}

// ReverseMapWithOverride takes a map with keys of type T and values of type U and returns
// a new map with the keys and values swapped. If there are duplicate values, the
// function will override the previous value.
func ReverseMapWithOverride[T comparable, U comparable](m map[T]U) map[U]T {
	result := make(map[U]T, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// ReverseMapX takes a map with keys of type T and values of type U and returns
// a new map with the keys and values swapped. If there are duplicate values, the
// function will return a Result object containing an error.
func ReverseMapX[T comparable, U comparable](m map[T]U) fxtypes.Result[map[U]T] {
	result := make(map[U]T, len(m))
	for k, v := range m {
		if _, ok := result[v]; ok {
			return fxtypes.NewErrorResult[map[U]T](fmt.Errorf("duplicate value found: %v", v))
		}
		result[v] = k
	}
	return fxtypes.NewSuccessResult(result)
}
