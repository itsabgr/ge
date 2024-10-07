package ge

import (
	"errors"
)

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

var ErrAssertionFailed = errors.New("assertion failed")

func Assert(cond bool, errs ...error) {
	if cond {
		return
	}
	switch len(errs) {
	case 0:
		panic(ErrAssertionFailed)
	case 1:
		panic(errs[0])
	default:
		panic(Join(errs...))
	}
}

func Try(fn func()) (recovered any) {
	defer func() {
		recovered = recover()
	}()
	fn()
	return nil
}
