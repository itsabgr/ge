package ge

import (
	"errors"
)

type UnwrapError interface {
	Unwrap() error
}

type UnwrapErrors interface {
	Unwrap() []error
}

func Unwrap(err error) error {
	u, ok := err.(UnwrapError)
	if !ok {
		return nil
	}
	return u.Unwrap()
}

func UnwrapAll(err error) []error {
	u, ok := err.(UnwrapErrors)
	if !ok {
		return nil
	}
	return u.Unwrap()
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func Join(errs ...error) (err error) {
	err = errors.Join(errs...)
	if errs = UnwrapAll(err); len(errs) == 1 {
		return errs[0]
	}
	return err
}

func As[T any](err error) (T, bool) {
	var val T
	return val, errors.As(err, &val)
}

func ErrOf(errorable any) error {
	if e, ok := errorable.(interface{ Err() error }); ok {
		return e.Err()
	}
	return nil
}
