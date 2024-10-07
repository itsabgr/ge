package ge

type Result[T any] struct {
	value T
	error error
}

func (r Result[T]) Must(wrap error) T {
	if r.error == nil {
		return r.value
	}
	var err error
	if wrap == nil {
		err = r.error
	} else {
		err = Wrap(r.error, wrap)
	}
	panic(err)
}

func (r Result[T]) Err() error {
	return r.error
}

func Err[T any](err error) Result[T] {
	return Result[T]{error: err}
}

func OK[T any](val T) Result[T] {
	return Result[T]{value: val}
}
