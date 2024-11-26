package plot

import (
	"github.com/itsabgr/ge"
	"iter"
)

func walk(err error) iter.Seq2[int, error] {
	return func(yield func(int, error) bool) {
		walkFn(err, 0, func(depth int, err error) bool {
			return yield(depth, err)
		})
	}
}

func errOf(errorable any) error {
	if e, ok := errorable.(interface{ Err() error }); ok {
		return e.Err()
	}
	return nil
}

func walkFn(root error, depth int, fn func(depth int, err error) bool) {

	if root == nil {
		return
	}

	if !fn(depth, root) {
		return
	}

	depth += 1

	switch e := errOf(root).(type) {
	case nil:
	case ge.UnwrapError:
		err := e.Unwrap()
		walkFn(err, depth, fn)
	case ge.UnwrapErrors:
		for _, err := range e.Unwrap() {
			walkFn(err, depth, fn)
		}
	default:
		fn(depth, e)
	}

	switch e := root.(type) {
	case ge.UnwrapError:
		err := e.Unwrap()
		walkFn(err, depth, fn)
	case ge.UnwrapErrors:
		for _, err := range e.Unwrap() {
			walkFn(err, depth, fn)
		}
	}

}
