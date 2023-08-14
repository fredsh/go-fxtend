package fx

import (
	"bytes"
	"encoding/json"
	"errors"
)

type Maybe[T any] struct {
	value T
	isSet bool
}

func NewSome[T any](value T) Maybe[T] {
	return Maybe[T]{value: value, isSet: true}
}

func NewNone[T any]() Maybe[T] {
	return Maybe[T]{isSet: false}
}

func (m Maybe[T]) IsSome() bool {
	return m.isSet
}

func (m Maybe[T]) IsNone() bool {
	return !m.isSet
}

// Unwrap returns the value pointer if it is set or panic if none.
func (m Maybe[T]) Unwrap() T {
	if !m.isSet {
		panic("attempted to Unwrap None value of Maybe")
	}
	return m.value
}

func (o Maybe[T]) Get() (T, error) {
	if !o.isSet {
		var defaultValue T
		return defaultValue, errors.New("no value set in Maybe")
	}
	return o.value, nil
}

func (m Maybe[T]) OrElse(def T) T {
	if !m.isSet {
		return def
	}
	return m.value
}

func (m Maybe[T]) MarshalJSON() ([]byte, error) {
	if m.isSet {
		return json.Marshal(m.value)
	} else {
		return json.Marshal(nil)
	}
}

func (m *Maybe[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		m.isSet = false
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	m.value = value
	m.isSet = true
	return nil
}

func MaybeMap[T, U any](m Maybe[T], fn func(T) U) Maybe[U] {
	if m.IsNone() {
		return NewNone[U]()
	}
	return NewSome(fn(m.Unwrap()))
}

func MaybeFlatMap[T, U any](opt Maybe[T], fn func(T) Maybe[U]) Maybe[U] {
	if opt.IsNone() {
		return NewNone[U]()
	}
	return fn(opt.Unwrap())
}
