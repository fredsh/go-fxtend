package fxres

import (
	"context"

	fxtypes "github.com/fredsh/go-fxtend/pkg/fx-types"
)

func Success[T any](success T) fxtypes.Result[T] {
	return fxtypes.NewSuccessResult(success)
}

func Failure[T any](err error) fxtypes.Result[T] {
	return fxtypes.NewErrorResult[T](err)
}

// Map applies the provided function to the value of the result and returns a new Result
// with the result of the function. If the original Result had an error, Map returns a new
// Result with the same error.
func Map[T any, U any](r fxtypes.Result[T], f func(T) U) fxtypes.Result[U] {
	if r.IsError() {
		return fxtypes.NewErrorResult[U](r.AsError())
	}
	return fxtypes.NewSuccessResult(f(*r.Unwrap()))
}

// FlatMapErr applies the given function to the value inside the Result,
// returning a new Result with the flattened output.
// If the Result is an error, the function is not applied and the error is propagated.
func FlatMapErr[T any, U any](success T, err error, f func(T) (U, error)) (def U, newErr error) {
	if err != nil {
		return def, err
	}
	return f(success)
}

// FlatMap applies the given function to the value inside the Result,
// returning a new Result with the flattened output.
// If the Result is an error, the function is not applied and the error is propagated.
func FlatMap[T any, U any](r fxtypes.Result[T], f func(T) fxtypes.Result[U]) fxtypes.Result[U] {
	if r.IsError() {
		return fxtypes.NewErrorResult[U](r.AsError())
	}
	return f(*r.Unwrap())
}

// Map applies the provided function to the value of the result and returns a new Result
// with the result of the function. If the original Result had an error, Map returns a new
// Result with the same error.
func MapCtx[T any, U any](ctx context.Context, r fxtypes.Result[T], f func(context.Context, T) U) fxtypes.Result[U] {
	if ctx.Err() != nil {
		return fxtypes.NewErrorResult[U](ctx.Err())
	}
	if r.IsError() {
		return fxtypes.NewErrorResult[U](r.AsError())
	}
	return fxtypes.NewSuccessResult(f(ctx, *r.Unwrap()))
}

// FlatMap applies the given function to the value inside the Result,
// returning a new Result with the flattened output.
// If the Result is an error, the function is not applied and the error is propagated.
func FlatMapCtx[T any, U any](ctx context.Context, r fxtypes.Result[T], f func(context.Context, T) fxtypes.Result[U]) fxtypes.Result[U] {
	if ctx.Err() != nil {
		return fxtypes.NewErrorResult[U](ctx.Err())
	}
	if r.IsError() {
		return fxtypes.NewErrorResult[U](r.AsError())
	}
	return f(ctx, *r.Unwrap())
}

// FlatMapErrCtx applies the given function to the value inside the Result,
// returning a new Result with the flattened output.
// If the Result is an error, the function is not applied and the error is propagated.
func FlatMapErrCtx[T any, U any](ctx context.Context, success T, err error, f func(context.Context, T) (U, error)) (def U, newErr error) {
	if err != nil {
		return def, err
	}
	return f(ctx, success)
}
