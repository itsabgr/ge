package ge

import (
	"iter"
)

func Walk(err error) iter.Seq2[int, error] {
	return func(yield func(int, error) bool) {
		walk(err, 0, func(depth int, err error) bool {
			return yield(depth, err)
		})
	}
}

func walk(root error, depth int, fn func(depth int, err error) bool) {

	if root == nil {
		return
	}

	if !fn(depth, root) {
		return
	}

	depth += 1

	switch e := ErrorOf(root).(type) {
	case nil:
	case UnwrapError:
		err := e.Unwrap()
		walk(err, depth, fn)
	case UnwrapErrors:
		for _, errs := range e.Unwrap() {
			walk(errs, depth, fn)
		}
	default:
		fn(depth, e)
	}

	switch e := root.(type) {
	case UnwrapError:
		err := e.Unwrap()
		walk(err, depth, fn)
	case UnwrapErrors:
		for _, errs := range e.Unwrap() {
			walk(errs, depth, fn)
		}
	}

}
