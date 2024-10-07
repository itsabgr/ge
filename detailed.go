package ge

import (
	"iter"
	"maps"
)

type D map[string]any

type Detailed struct {
	err     error
	details D
}

func (d Detailed) Error() string {
	return d.err.Error()
}

func (d Detailed) String() string {
	return "detailed error: " + d.Error()
}

func (d Detailed) Err() error {
	return d.err
}

func (d Detailed) Details() iter.Seq2[string, any] {
	return maps.All(d.details)
}

func (d Detailed) LookupDetails(key string) (any, bool) {
	v, f := d.details[key]
	return v, f
}

func WithDetails(err error, details D) error {
	if err == nil {
		return nil
	}
	if len(details) == 0 {
		return err
	}
	return Detailed{err, details}
}

func LookupDetails(err error, key string) (any, bool) {
	if e, ok := err.(interface{ LookupDetails(key string) (any, bool) }); ok {
		return e.LookupDetails(key)
	}
	return nil, false
}

func Details(err error) iter.Seq2[string, any] {
	if e, ok := err.(interface{ Details() iter.Seq2[string, any] }); ok {
		return e.Details()
	}
	return nil
}
