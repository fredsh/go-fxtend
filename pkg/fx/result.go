package fx

import (
	"context"
)

// Result is a type representing either a success value or an error
type Result[T any] struct {
	value T
	err   error
}

// NewSuccess creates a new Result with a success value
func NewSuccess[T any](value T) Result[T] {
	return Result[T]{value: value}
}

// NewFailure creates a new Result with an error
func NewFailure[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// IsSuccess returns true if the result is a success
func (r Result[T]) IsSuccess() bool {
	return r.err == nil
}

// IsError returns true if the result is an error
func (r Result[T]) IsError() bool {
	return r.err != nil
}

// Unwrap returns the success value or panics if the result is an error
func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic("attempted to Unwrap Failure value of Result")
	}
	return r.value
}

// UnwrapOr returns the success value or a default value if the result is an error
func (r Result[T]) UnwrapOr(def T) T {
	if r.err != nil {
		return def
	}
	return r.value
}

// UnwrapOrElse returns the success value or the result of a function if the result is an error
func (r Result[T]) UnwrapOrElse(fn func(error) T) T {
	if r.err != nil {
		return fn(r.err)
	}
	return r.value
}

// UnwrapErr returns the success value and the error separately
func (r Result[T]) UnwrapErr() (T, error) {
	return r.value, r.err
}

// AsError returns the error wrapped by the result, or nil if the result is a success
func (r Result[T]) AsError() error {
	return r.err
}

// Map applies the provided function to the value of the result and returns a new Result
// with the result of the function. If the original Result had an error, Map returns a new
// Result with the same error.
func Map[T any, U any](r Result[T], f func(T) U) Result[U] {
	if r.IsError() {
		return NewFailure[U](r.AsError())
	}
	return NewSuccess(f(r.Unwrap()))
}

// FlatMapErr applies the given function to the value inside the Result,
// returning a new Result with the flattened output.
// If the Result is an error, the function is not applied and the error is propagated.
func FlatMapErr[T any, U any](r Result[T], f func(T) (U, error)) Result[U] {
	if r.IsError() {
		return NewFailure[U](r.AsError())
	}
	res, err := f(r.Unwrap())
	return Result[U]{
		value: res,
		err:   err,
	}
}

// FlatMap applies the given function to the value inside the Result,
// returning a new Result with the flattened output.
// If the Result is an error, the function is not applied and the error is propagated.
func FlatMap[T any, U any](r Result[T], f func(T) Result[U]) Result[U] {
	if r.IsError() {
		return NewFailure[U](r.AsError())
	}
	return f(r.Unwrap())
}

// Map applies the provided function to the value of the result and returns a new Result
// with the result of the function. If the original Result had an error, Map returns a new
// Result with the same error.
func MapCtx[T any, U any](ctx context.Context, r Result[T], f func(context.Context, T) U) Result[U] {
	if ctx.Err() != nil {
		return NewFailure[U](ctx.Err())
	}
	if r.IsError() {
		return NewFailure[U](r.AsError())
	}
	return NewSuccess(f(ctx, r.Unwrap()))
}

// FlatMap applies the given function to the value inside the Result,
// returning a new Result with the flattened output.
// If the Result is an error, the function is not applied and the error is propagated.
func FlatMapCtx[T any, U any](ctx context.Context, r Result[T], f func(context.Context, T) Result[U]) Result[U] {
	if ctx.Err() != nil {
		return NewFailure[U](ctx.Err())
	}
	if r.IsError() {
		return NewFailure[U](r.AsError())
	}
	return f(ctx, r.Unwrap())
}

// FlatMapErrCtx applies the given function to the value inside the Result,
// returning a new Result with the flattened output.
// If the Result is an error, the function is not applied and the error is propagated.
func FlatMapErrCtx[T any, U any](ctx context.Context, r Result[T], f func(context.Context, T) (U, error)) Result[U] {
	if ctx.Err() != nil {
		return NewFailure[U](ctx.Err())
	}
	if r.IsError() {
		return NewFailure[U](r.AsError())
	}
	res, err := f(ctx, r.Unwrap())
	return Result[U]{
		value: res,
		err:   err,
	}
}
