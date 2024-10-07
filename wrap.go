package ge

type Wrapped struct {
	wrapper  error
	original error
}

func (w Wrapped) Error() string {
	return w.wrapper.Error()
}

func (w Wrapped) Err() error {
	return w.wrapper
}

func (w Wrapped) String() string {
	return "error: " + w.wrapper.Error() + ": " + w.original.Error()
}

func (w Wrapped) Unwrap() error {
	return w.original
}

func Wrap(newWrapperError, originalError error) error {
	if originalError == nil {
		return nil
	}
	return Wrapped{newWrapperError, originalError}
}
