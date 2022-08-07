package form

import (
	"fmt"
	"net/url"
)

type ValueEncoding[T any] interface {
	Encode(T) string
	Decode(string) (T, error)
	Equal(T, T) bool // Value comparison isn't strictly related to encoding. Does this belong here?
}

type Valuer interface {
	MarshalForm() []string
	UnmarshalForm([]string) error
}

type Fields map[string]Valuer

func Encode(fields Fields) url.Values {
	values := make(url.Values, len(fields))
	for name, field := range fields {
		value := field.MarshalForm()
		if len(value) > 0 {
			values[name] = value
		}
	}
	return values
}

func Decode(fields Fields, values url.Values) error {
	for _, field := range fields {
		if df, ok := field.(interface{ ApplyDefault() }); ok {
			df.ApplyDefault()
		}
	}
	for name, field := range fields {
		value, ok := values[name]
		if !ok {
			continue
		}
		if err := field.UnmarshalForm(value); err != nil {
			return fmt.Errorf("%s: %v", name, err)
		}
	}
	return nil
}

type Value[T any] struct {
	Encoding ValueEncoding[T]
	Value    *T
	Default  T
}

func (v Value[T]) MarshalForm() []string {
	// Don't encode the zero value.
	var zero T
	if v.Encoding.Equal(*v.Value, zero) {
		return nil
	}
	return []string{v.Encoding.Encode(*v.Value)}
}

func (v Value[T]) UnmarshalForm(values []string) (err error) {
	if len(values) < 1 {
		return nil
	}
	*v.Value, err = v.Encoding.Decode(values[0])
	return err
}

func (v Value[T]) WithDefault(value T) Value[T] {
	v.Default = value
	return v
}

func (v Value[T]) ApplyDefault() {
	*v.Value = v.Default
}
