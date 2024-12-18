package ge

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

type detailed struct {
	err     error
	details D
}

func Detail(err error, details D) error {
	if err == nil {
		return nil
	}
	if len(details) == 0 {
		return err
	}
	return detailed{err, details}
}

func Details(err error) iter.Seq2[string, any] {
	if e, ok := err.(Detailed); ok {
		return e.Details()
	}
	return nil
}

func (d detailed) Error() string {
	return d.err.Error()
}

func (d detailed) String() string {
	return "error: " + d.err.Error() + " " + d.details.String()
}

func (d detailed) Err() error {
	return d.err
}

func (d detailed) Details() iter.Seq2[string, any] {
	return maps.All(d.details)
}

func (d D) String() string {
	if len(d) == 0 {
		return "[]"
	}
	builder := strings.Builder{}
	for k, v := range d {
		builder.WriteString(k)
		builder.WriteString(":")
	        builder.WriteString(fmt.Sprint(v))
		builder.WriteByte(' ')
	}
	str := builder.String()
	return "[" + str[:len(str)-1] + "]"
}
