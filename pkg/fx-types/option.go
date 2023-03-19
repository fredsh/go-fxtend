package fxtypes

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Option[T any] struct {
	value T
	isSet bool
}

func NewValueOption[T any](value T) Option[T] {
	return Option[T]{value: value, isSet: true}
}

func NewNoneNone[T any]() Option[T] {
	return Option[T]{isSet: false}
}

func (o Option[T]) IsSome() bool {
	return o.isSet
}

func (o Option[T]) IsNone() bool {
	return !o.isSet
}

// Unwrap returns the value pointer if it is set or nil if it is not.
func (o Option[T]) Unwrap() *T {
	if !o.isSet {
		return nil
	}
	return &o.value
}

func (o Option[T]) Get() (T, error) {
	if !o.isSet {
		var defaultValue T
		return defaultValue, fmt.Errorf("Option is empty")
	}
	return o.value, nil
}

func (o Option[T]) OrElse(def T) T {
	if !o.isSet {
		return def
	}
	return o.value
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: value, isSet: true}
}

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.isSet {
		return json.Marshal(o.value)
	} else {
		return json.Marshal(nil)
	}
}

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		o.isSet = false
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	o.value = value
	o.isSet = true
	return nil
}
