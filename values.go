package form

import (
	"github.com/deitrix/go-form/constraints"
	"github.com/deitrix/go-form/encoding"
)

func Float[T constraints.Float](val *T) Value[T] {
	return Value[T]{
		Encoding: encoding.DefaultFloat[T](),
		Value:    val,
	}
}

func Int[T constraints.Signed](val *T) Value[T] {
	return Value[T]{
		Encoding: encoding.DefaultInt[T](),
		Value:    val,
	}
}

func String[T constraints.String](val *T) Value[T] {
	return Value[T]{
		Encoding: encoding.DefaultString[T](),
		Value:    val,
	}
}

func Uint[T constraints.Unsigned](val *T) Value[T] {
	return Value[T]{
		Encoding: encoding.DefaultUint[T](),
		Value:    val,
	}
}
