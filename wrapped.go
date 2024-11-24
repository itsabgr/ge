package ge

type wrapped struct {
	wrapper  error
	original error
}

func (w wrapped) Error() string {
	return w.wrapper.Error()
}

func (w wrapped) Err() error {
	return w.wrapper
}

func (w wrapped) String() string {
	return "error: " + w.wrapper.Error() + ": " + w.original.Error()
}

func (w wrapped) Unwrap() error {
	return w.original
}

func Wrap(newWrapperError, originalError error) error {
	if originalError == nil {
		return nil
	}
	return wrapped{newWrapperError, originalError}
}
