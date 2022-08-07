package encoding

import "github.com/deitrix/go-form/constraints"

func DefaultString[T constraints.String]() String[T] {
	return String[T]{}
}

type String[T constraints.String] struct{ Comparable[T] }

func (String[T]) Encode(val T) string {
	return string(val)
}

func (String[T]) Decode(s string) (T, error) {
	return T(s), nil
}
