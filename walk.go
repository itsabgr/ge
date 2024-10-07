package ge

import "iter"

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

	if e, ok := root.(interface{ Unwrap() []error }); ok {
		for _, errs := range e.Unwrap() {
			walk(errs, depth, fn)
		}
	}
	if e, ok := root.(interface{ Unwrap() error }); ok {
		err := e.Unwrap()
		walk(err, depth, fn)
	}

}
