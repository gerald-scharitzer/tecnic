package reflect

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

// This is just to test the iterator interface with arrays.
// Arrays can already be iterated over with a for range loop.
type arrayIterator[T any] struct {
	array *[]T
	index int
}

// This is just to test the iterator interface with strings.
// Strings can already be iterated over with a for range loop.
type stringIterator struct {
	reader *strings.Reader
}

func newArrayIterator[T any](array *[]T) *arrayIterator[T] {
	return &arrayIterator[T]{array: array, index: 0}
}

func newStringIterator(str *string) *stringIterator {
	reader := strings.NewReader(*str)
	return &stringIterator{reader: reader}
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

// Implements Iterator
func (iter *stringIterator) Next() (rune, bool) {
	char, _, err := (*iter.reader).ReadRune()
	if err == nil {
		return char, true
	}
	return 0, false
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

func TestStringIterator(t *testing.T) {
	str := "hello"

	// test with for range
	iter := newStringIterator(&str)
	for _, want := range str {
		value, ok := iter.Next()
		assert.Assert(t, ok)
		assert.Equal(t, value, want)
	}
	_, ok := iter.Next()
	assert.Assert(t, !ok)

	// test with for while
	iter = newStringIterator(&str)
	x := 0
	for value, ok := iter.Next(); ok; value, ok = iter.Next() {
		assert.Equal(t, value, rune(str[x])) // this works only because the string is ASCII
		x++
	}
}
