package encoding

import (
	"strconv"

	"github.com/deitrix/go-form/constraints"
)

func DefaultInt[T constraints.Signed]() Int[T] {
	return Int[T]{
		Base: 10,
	}
}

type Int[T constraints.Signed] struct {
	Comparable[T]
	Base int
}

func (Int[T]) Encode(val T) string {
	return strconv.FormatInt(int64(val), 10)
}

func (Int[T]) Decode(s string) (T, error) {
	val, err := strconv.ParseInt(s, 10, 64)
	return T(val), err
}
