package de

import (
	"iter"
	"maps"
)

type D map[string]any

type Detailed interface {
	Details() iter.Seq2[string, any]
}

type Error struct {
	err     error
	details D
}

func New(err error, details D) error {
	if err == nil {
		return nil
	}
	if len(details) == 0 {
		return err
	}
	return Error{err, details}
}

func DetailsOf(err error) iter.Seq2[string, any] {
	if e, ok := err.(Detailed); ok {
		return e.Details()
	}
	return nil
}

func (d Error) Error() string {
	return d.err.Error()
}

func (d Error) String() string {
	return "detailed error: " + d.Error()
}

func (d Error) Err() error {
	return d.err
}

func (d Error) Details() iter.Seq2[string, any] {
	return maps.All(d.details)
}
