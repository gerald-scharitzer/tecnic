package reflect

import (
	"reflect"
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

// This is just to test the iterator interface with maps.
// Maps can already be iterated over with a for range loop.
type mapIterator[K comparable, V any] struct {
	iterator *reflect.MapIter
}

// This is just to test the iterator interface with channels.
// Channels can already be iterated over with a for range loop.
type channelIterator[T any] struct {
	// The channel must be readable and the iterator must not write to the channel.
	// TODO test with both read-only and read-write channels
	channel *<-chan T
}

func newArrayIterator[T any](array *[]T) *arrayIterator[T] {
	return &arrayIterator[T]{array: array, index: 0}
}

func newStringIterator(str *string) *stringIterator {
	reader := strings.NewReader(*str)
	return &stringIterator{reader: reader}
}

func newMapIterator[K comparable, V any](mapPointer *map[K]V) *mapIterator[K, V] {
	mapValue := reflect.ValueOf(*mapPointer)
	iterator := mapValue.MapRange()
	return &mapIterator[K, V]{iterator: iterator}
}

func newChannelIterator[T any](channel *chan T) *channelIterator[T] {
	receiver := (<-chan T)(*channel)
	return &channelIterator[T]{channel: &receiver}
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
func (iter *stringIterator) Next() (rune, bool) { // FIXME interface is not [T any] (*T, bool)
	char, _, err := (*iter.reader).ReadRune()
	if err == nil {
		return char, true
	}
	return 0, false
}

// Implements Iterator
func (iter *mapIterator[K, V]) Next() (*K, *V, bool) { // FIXME interface is not [T any] (*T, bool)
	iterator := iter.iterator
	if iterator.Next() {
		key := iterator.Key().Interface().(K)
		value := iterator.Value().Interface().(V)
		return &key, &value, true
	}
	return nil, nil, false
}

// Implements Iterator
func (iter *channelIterator[T]) Next() (*T, bool) {
	value, ok := <-*iter.channel
	return &value, ok
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

func TestMapIterator(t *testing.T) {
	m := map[int]string{1: "one", 2: "two", 3: "three"}

	// test with for range
	iter := newMapIterator(&m)
	for key, want := range m {
		k, v, ok := iter.Next()
		assert.Assert(t, ok)
		assert.Equal(t, *k, key) // FIXME there is no guarantee that the order will be the same
		assert.Equal(t, *v, want)
	}
	_, _, ok := iter.Next()
	assert.Assert(t, !ok)

	// test with for while
	iter = newMapIterator(&m)
	for k, v, ok := iter.Next(); ok; k, v, ok = iter.Next() {
		assert.Equal(t, m[*k], *v)
	}
}

func TestChannelIterator(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 0
	ch <- 1
	ch <- 2
	close(ch)

	// test with for clause
	iter := newChannelIterator(&ch)
	for want := 0; want <= 2; want++ {
		value, ok := iter.Next()
		assert.Assert(t, ok)
		assert.Equal(t, *value, want)
	}
	_, ok := iter.Next()
	assert.Assert(t, !ok)

	// re-make channel, because it was closed
	ch = make(chan int, 3)
	ch <- 0
	ch <- 1
	ch <- 2
	close(ch)

	// test with for while
	iter = newChannelIterator(&ch)
	want := 0
	for value, ok := iter.Next(); ok; value, ok = iter.Next() {
		assert.Assert(t, ok)
		assert.Equal(t, *value, want)
		want++
	}
}
