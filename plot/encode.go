package plot

import (
	"github.com/itsabgr/ge"
	"github.com/itsabgr/ge/pkg/de"
	"maps"
)

type jsonErr struct {
	Text    string         `json:"text,omitempty" yaml:"text,omitempty" bson:"text,omitempty" xml:"text,omitempty"`
	Error   any            `json:"error,omitempty" yaml:"error,omitempty" bson:"error,omitempty" xml:"error,omitempty"`
	Wrapped any            `json:"wrapped,omitempty" yaml:"wrapped,omitempty" bson:"wrapped,omitempty" xml:"wrapped,omitempty"`
	Details map[string]any `json:"details,omitempty" yaml:"details,omitempty" bson:"details,omitempty" xml:"details,omitempty"`
}

func Encode(root error) any {

	if root == nil {
		return nil
	}

	if err, ok := root.(ge.UnwrapErrors); ok {
		errs := err.Unwrap()
		list := make([]any, len(errs))
		for i, err := range errs {
			list[i] = Encode(err)
		}
		switch len(list) {
		case 0:
			return nil
		case 1:
			return list[0]
		default:
			return list
		}
	}

	txt := root.Error()

	var details map[string]any

	if iter := de.DetailsOf(root); iter != nil {
		if collected := maps.Collect[string, any](iter); len(collected) > 0 {
			details = collected
		}
	}

	inner := Encode(ge.ErrorOf(root))

	var wrapped any

	if unwrapError, ok := root.(ge.UnwrapError); ok {
		wrapped = Encode(unwrapError.Unwrap())
	}

	if inner == nil && wrapped == nil && len(details) == 0 && txt != "" {
		return txt
	}

	if inner != nil {
		txt = ""
	}

	err := jsonErr{
		Text:    txt,
		Error:   inner,
		Wrapped: wrapped,
		Details: details,
	}

	return err

}
