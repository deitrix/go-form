package encoding

type Comparable[T comparable] struct{}

func (c Comparable[T]) Equal(a, b T) bool {
	return a == b
}
