package fxopt

import fxtypes "github.com/fredsh/go-fxtend/pkg/fx-types"

// Some creates a new Option with a value
func Some[T any](value T) fxtypes.Option[T] {
	return fxtypes.NewValueOption(value)
}

// None creates a new empty Option
func None[T any]() fxtypes.Option[T] {
	return fxtypes.NewNoneNone[T]()
}

// Map applies a function to the value of an Option and returns a new Option with the result
func Map[T, U any](opt fxtypes.Option[T], fn func(T) U) fxtypes.Option[U] {
	if opt.IsNone() {
		return fxtypes.NewNoneNone[U]()
	}
	return fxtypes.Some(fn(*opt.Unwrap()))
}

// FlatMap applies a function to the value of an Option and returns a new Option with the result
func FlatMap[T, U any](opt fxtypes.Option[T], fn func(T) fxtypes.Option[U]) fxtypes.Option[U] {
	if opt.IsNone() {
		return fxtypes.NewNoneNone[U]()
	}
	return fn(*opt.Unwrap())
}
