package fx

import (
	"context"
)

type MapFilterFunc[K comparable, V any] func(key K, value V) bool
type KeyValueMapperFunc[K1, K2 comparable, V1, V2 any] func(k K1, v V1) (K2, V2, error)
type KeyValueMapperFuncCtx[K1, K2 comparable, V1, V2 any] func(ctx context.Context, k K1, v V1) (K2, V2, error)

func MapGetValues[K comparable, V any](input map[K]V) []V {
	res := make([]V, 0, len(input))
	for _, v := range input {
		res = append(res, v)
	}
	return res
}

func MapGetKeys[K comparable, V any](input map[K]V) []K {
	res := make([]K, 0, len(input))
	for k := range input {
		res = append(res, k)
	}
	return res
}

func MapFilter[K comparable, V any](input map[K]V, shouldKeep MapFilterFunc[K, V]) map[K]V {
	res := make(map[K]V, len(input))
	for k, v := range input {
		if shouldKeep(k, v) {
			res[k] = v
		}
	}
	return res
}

func MapFilterMut[K comparable, V any](input map[K]V, shouldKeep MapFilterFunc[K, V]) {
	for k, v := range input {
		if !shouldKeep(k, v) {
			delete(input, k)
		}
	}
}

func MapApply[K1, K2 comparable, V1, V2 any](input map[K1]V1, mapper KeyValueMapperFunc[K1, K2, V1, V2]) (map[K2]V2, []error) {
	res := make(map[K2]V2, len(input))
	errs := []error{}

	for k1, v1 := range input {
		k2, v2, err := mapper(k1, v1)
		if err != nil {
			errs = append(errs, err)
		} else {
			res[k2] = v2
		}
	}
	return res, errs
}

func MapFilterApply[K1, K2 comparable, V1, V2 any](
	input map[K1]V1,
	shouldKeep MapFilterFunc[K1, V1],
	mapper KeyValueMapperFunc[K1, K2, V1, V2],
) (map[K2]V2, []error) {
	res := make(map[K2]V2, len(input))
	errs := []error{}

	for k1, v1 := range input {
		if !shouldKeep(k1, v1) {
			continue
		}
		k2, v2, err := mapper(k1, v1)
		if err != nil {
			errs = append(errs, err)
		} else {
			res[k2] = v2
		}
	}
	return res, errs
}
