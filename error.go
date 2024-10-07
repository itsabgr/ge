package ge

import "fmt"

type Error[T comparable] struct {
	value T
}

func (e Error[T]) String() string {
	return "error: " + e.Error()
}

func (e Error[T]) Error() string {
	return fmt.Sprint(e.value)
}

func (e Error[T]) Value() any {
	return e.value
}

func New[T comparable](value T) error {
	if e, ok := any(value).(error); ok {
		return e
	}
	return Error[T]{value: value}
}
