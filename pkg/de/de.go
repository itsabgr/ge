package de

import (
	"fmt"
	"iter"
	"maps"
	"strings"
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
	return "error: " + d.err.Error() + " " + d.details.String()
}

func (d Error) Err() error {
	return d.err
}

func (d Error) Details() iter.Seq2[string, any] {
	return maps.All(d.details)
}

func (d D) String() string {
	builder := strings.Builder{}
	for k, v := range d {
		builder.WriteString(k)
		builder.WriteString(":")
		builder.WriteString(fmt.Sprint(v))
		builder.WriteByte(' ')
	}
	str := builder.String()
	if len(str) == 0 {
		return "[]"
	}
	return "[" + str[:len(str)-1] + "]"
}
