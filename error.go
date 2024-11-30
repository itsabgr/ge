package ge

import "fmt"

type err[T comparable] struct {
	value T
}

func (e err[T]) String() string {
	return "error: " + e.Error()
}

func (e err[T]) Error() string {
	return fmt.Sprint(e.value)
}

func (e err[T]) Value() any {
	return e.value
}

func New[T comparable](value T) error {
	if e, ok := any(value).(error); ok {
		return e
	}
	return err[T]{value: value}
}
