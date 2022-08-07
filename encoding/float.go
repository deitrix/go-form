package encoding

import (
	"strconv"

	"github.com/deitrix/go-form/constraints"
)

func DefaultFloat[T constraints.Float]() Float[T] {
	return Float[T]{
		Fmt:  'f',
		Prec: -1,
	}
}

type Float[T constraints.Float] struct {
	Comparable[T]
	Fmt  byte
	Prec int
}

func (f Float[T]) Encode(val T) string {
	return strconv.FormatFloat(float64(val), f.Fmt, f.Prec, 64)
}

func (Float[T]) Decode(s string) (T, error) {
	val, err := strconv.ParseFloat(s, 64)
	return T(val), err
}
