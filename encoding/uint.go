package encoding

import (
	"strconv"

	"github.com/deitrix/go-form/constraints"
)

func DefaultUint[T constraints.Unsigned]() Uint[T] {
	return Uint[T]{
		Base: 10,
	}
}

type Uint[T constraints.Unsigned] struct {
	Comparable[T]
	Base int
}

func (u Uint[T]) Encode(val T) string {
	return strconv.FormatUint(uint64(val), u.Base)
}

func (u Uint[T]) Decode(s string) (T, error) {
	val, err := strconv.ParseUint(s, u.Base, 64)
	return T(val), err
}
