package fxcollection

import (
	"fmt"

	fxtypes "github.com/fredsh/go-fxtend/pkg/fx-types"
)

// ToMapWithOverride turns a slice of element of type V and turn it into a map
// the key is determined by executing keySelector on each item.
// If duplicate keys are found, later items overwrite earlier ones.
func ToMapWithOverride[K comparable, V any](input []V, keySelector func(item V) K) map[K]V {
	result := make(map[K]V, len(input))
	for _, v := range input {
		itemKey := keySelector(v)
		result[itemKey] = v
	}
	return result
}

// ToMap turns a slice of element of type V and turn it into a map
// the key is determined by executing keySelector on each item.
// If duplicate keys are found, it returns an error.
func ToMap[K comparable, V any](input []V, keySelector func(item V) K) (map[K]V, error) {
	result := make(map[K]V, len(input))
	for _, v := range input {
		itemKey := keySelector(v)
		if _, ok := result[itemKey]; ok {
			return nil, fmt.Errorf("duplicate key found: %v", itemKey)
		}
		result[itemKey] = v
	}
	return result, nil
}

// ToMapX turns a slice of element of type V and turn it into a map
// the key is determined by executing keySelector on each item.
// If duplicate keys are found, it returns a Result object.
func ToMapX[K comparable, V any](input []V, keySelector func(item V) K) fxtypes.Result[map[K]V] {
	result := make(map[K]V, len(input))
	for _, v := range input {
		itemKey := keySelector(v)
		if _, ok := result[itemKey]; ok {
			return fxtypes.NewErrorResult[map[K]V](fmt.Errorf("duplicate key found: %v", itemKey))
		}
		result[itemKey] = v
	}
	return fxtypes.NewSuccessResult(result)
}

// PrepareToMapWithOverride returns a function that can be used to convert a slice to a map
// using the specified key selector function.
// If duplicate keys are found, later items overwrite earlier ones.
func PrepareToMapWithOverride[K comparable, V any](keySelector func(V) K) func([]V) map[K]V {
	return func(input []V) map[K]V {
		return ToMapWithOverride(input, keySelector)
	}
}

// PrepareToMap returns a function that can be used to convert a slice to a map
// using the specified key selector function.
// If duplicate keys are found, it returns an error.
func PrepareToMap[K comparable, V any](keySelector func(V) K) func([]V) (map[K]V, error) {
	return func(input []V) (map[K]V, error) {
		return ToMap(input, keySelector)
	}
}
