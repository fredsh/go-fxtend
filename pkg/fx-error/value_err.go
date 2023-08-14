package fxerror

import (
	"errors"
	"fmt"
)

var ErrDuplicateValue = errors.New("duplicate value encountered")

// ValueError represents an error that occurs when a duplicate value is encountered.
type ValueError struct {
	Value interface{} // The duplicate value that caused the error.
	err   error
}

func NewDuplicateValueError(value interface{}) *ValueError {
	return &ValueError{
		Value: value,
		err:   ErrDuplicateValue,
	}
}

// Error returns the error message for DuplicateValueError.
func (e *ValueError) Error() string {
	if e.err == nil {
		return fmt.Sprintf("non-categorized issue encountered with value: [%v]", e.Value)
	}
	return fmt.Sprintf("%s: [%v]", e.err.Error(), e.Value)
}

// Unwrap returns the wrapped error, which is always nil for DuplicateValueError.
func (e *ValueError) Unwrap() error {
	return e.err
}
