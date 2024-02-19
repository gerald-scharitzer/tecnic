package reflect

// Returns the next element in the iterator and a boolean indicating if there was one.
type Iterator[T any] interface {
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
