package fxtypes

// Result is a type representing either a success value or an error
type Result[T any] struct {
	value T
	err   error
}

// NewSuccessResult creates a new Result with a success value
func NewSuccessResult[T any](value T) Result[T] {
	return Result[T]{value: value}
}

// NewErrorResult creates a new Result with an error
func NewErrorResult[T any](err error) Result[T] {
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
func (r Result[T]) Unwrap() *T {
	if r.err != nil {
		return nil
	}
	return &r.value
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
