package ge

import (
	"errors"
)

var ErrAssertionFailed = errors.New("assertion failed")

func Assert(cond bool, errs ...error) {
	if cond {
		return
	}
	err := Join(errs...)
	if err == nil {
		panic(ErrAssertionFailed)
	}
	panic(errs)
}

func Throw(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func Must[T any](value T, err error) T {
	Throw(err)
	return value
}

func Try(fn func()) (recovered any) {
	defer func() {
		recovered = recover()
	}()
	fn()
	return nil
}

const UNREACHABLE = "UNREACHABLE"
