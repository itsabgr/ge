package ge

import (
	"errors"
)

func ErrorOf(err any) error {
	if e, ok := err.(interface{ Err() error }); ok {
		return e.Err()
	}
	return nil
}

type UnwrapError interface {
	Unwrap() error
}

type UnwrapErrors interface {
	Unwrap() []error
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func Join(errs ...error) error {
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errors.Join(errs...)
	}
}

func As[T any](err error) (T, bool) {
	var val T
	return val, errors.As(err, &val)
}
