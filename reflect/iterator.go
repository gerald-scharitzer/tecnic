package reflect

// Wrap a sequence of elements of type T in an iterator.
type Iterator[T any] interface {
	// Returns the next element and a boolean indicating if there was one.
	Next() (*T, bool)
}

// TODO
type Option[T any] interface {
	Value() *T
}

// TODO
type None struct{}

// TODO
func (x None) Value() *any {
	return nil
}

// TODO
type Some[T any] struct {
	value *T
}

// TODO
func (x Some[T]) Value() *T {
	return x.value
}
