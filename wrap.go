package ge

type Wrapped struct {
	wrapper  error
	original error
}

func (w Wrapped) Error() string {
	return w.wrapper.Error()
}

func (w Wrapped) String() string {
	return "wrapped error: " + w.Error()
}

func (w Wrapped) Unwrap() error {
	return w.original
}

func Wrap(originalError error, newWrapperError error) error {
	if originalError == nil {
		return nil
	}
	return Wrapped{newWrapperError, originalError}
}
