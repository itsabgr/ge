package ge

type Result[T any] struct {
	value T
	error error
}

func (r Result[T]) Must(wraps ...error) T {
	if r.error == nil {
		return r.value
	}
	err := Join(wraps...)
	if err == nil {
		panic(r.error)
	}
	panic(Wrap(err, r.error))
}

func (r Result[T]) Err() error {
	return r.error
}

func (r Result[T]) Unwrap() (T, error) {
	return r.value, r.error
}

func Err[T any](err error) Result[T] {
	return Result[T]{error: err}
}

func OK[T any](val T) Result[T] {
	return Result[T]{value: val}
}

func Res[T any](val T, err error) Result[T] {
	return Result[T]{val, err}
}
