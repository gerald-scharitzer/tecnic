package reflect

import (
	"testing"

	"gotest.tools/v3/assert"
)

// This is just to test the iterator interface with arrays.
// Arrays can already be iterated over with a for range loop.
type arrayIterator[T any] struct {
	array *[]T
	index int
}

func newArrayIterator[T any](array *[]T) *arrayIterator[T] {
	return &arrayIterator[T]{array: array, index: 0}
}

// Implements Iterator
func (iter *arrayIterator[T]) Next() (*T, bool) {
	array := *iter.array
	if iter.index < len(array) {
		value := array[iter.index]
		iter.index++
		return &value, true
	}
	return nil, false
}

func TestArrayIterator(t *testing.T) {
	array := []int{1, 2, 3, 4}

	// test with for range
	iter := newArrayIterator(&array)
	for _, want := range array {
		value, ok := iter.Next()
		assert.Assert(t, ok)
		assert.Equal(t, *value, want)
	}
	_, ok := iter.Next()
	assert.Assert(t, !ok)

	// test with for while
	iter = newArrayIterator(&array)
	for value, ok := iter.Next(); ok; value, ok = iter.Next() {
		assert.Equal(t, *value, array[iter.index-1])
	}
}
