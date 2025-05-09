package immutableiter

import (
	"reflect"
	"testing"
)

// TODO(Izzette): test some non-happy cases where the assertations in iter.go will fail.

func TestIndexedSeq(t *testing.T) {
	// Create an IndexedIterator
	it := &mockIndexedIterator[int]{values: []int{1, 2, 3, 4, 5}}

	// Create a sequence using IndexedSeq
	seq := IndexedSeq[int](it)

	// Initialize a counter to track the number of yielded values
	count := 0

	// Define a yield function that increments the counter
	yield := func(v int) bool {
		count++
		return true
	}

	// Call the sequence with the yield function
	seq(yield)

	// Check if the count matches the expected number of elements in the list
	if count != 5 {
		t.Errorf("Expected 5 elements, got %d", count)
	}
}

func TestIndexedSeqWithIndex(t *testing.T) {
	// Create an IndexedIterator
	it := &mockIndexedIterator[int]{values: []int{1, 2, 3, 4, 5}}

	// Create a sequence using IndexedSeqWithIndex
	seq := IndexedSeqWithIndex[int](it)

	// Initialize a counter to track the number of yielded values
	count := 0

	// Define a yield function that increments the counter
	yield := func(i int, v int) bool {
		count++
		return true
	}

	// Call the sequence with the yield function
	seq(yield)

	// Check if the count matches the expected number of elements in the list
	if count != 5 {
		t.Errorf("Expected 5 elements, got %d", count)
	}
}

func TestKeyedSeq(t *testing.T) {
	// Create a KeyedIterator
	it := &mockKeyedIterator[string, int]{
		Pairs: []kvPair[string, int]{
			{"a", 1},
			{"b", 2},
			{"c", 3},
		},
	}

	// Create a sequence using KeyedSeq
	seq := KeyedSeq[string, int](it)

	// Initialize a counter to track the number of yielded values
	count := 0
	found := make(map[string]int, len(it.Pairs))

	// Define a yield function that increments the counter
	yield := func(k string, v int) bool {
		count++
		// Store the key-value pair in the found map
		found[k] = v

		return true
	}
	// Call the sequence with the yield function
	seq(yield)

	// Check if the count matches the expected number of elements in the map
	if count != 3 {
		t.Errorf("Expected 3 calls of yeild, got %d", count)
	}

	// Check if the keys and values are correct
	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Expected %v, got %v", expected, found)
	}
}

func TestKeyedSeqSortedMap(t *testing.T) {
	// Create a KeyedIterator
	it := &mockKeyedIterator[string, int]{
		Pairs: []kvPair[string, int]{
			{"a", 1},
			{"b", 2},
			{"c", 3},
		},
	}

	// Create a sequence using KeyedSeq
	seq := KeyedSeq[string, int](it)

	// Initialize a counter to track the number of yielded values
	count := 0
	found := make(map[string]int, len(it.Pairs))

	// Define a yield function that increments the counter
	yield := func(k string, v int) bool {
		count++
		// Store the key-value pair in the found map
		found[k] = v

		return true
	}
	// Call the sequence with the yield function
	seq(yield)

	// Check if the count matches the expected number of elements in the map
	if count != 3 {
		t.Errorf("Expected 3 calls of yeild, got %d", count)
	}

	// Check if the keys and values are correct
	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Expected %v, got %v", expected, found)
	}
}

func TestUnkeyedSeq(t *testing.T) {
	// Create an UnkeyedIterator
	it := &mockUnkeyedIterator[int]{Values: []int{1, 2, 3}}

	// Create a sequence using UnkeyedSeq
	seq := UnkeyedSeq[int](it)

	// Initialize a counter to track the number of yielded values
	count := 0
	found := make(map[int]struct{}, len(it.Values))

	// Define a yield function that increments the counter
	yield := func(v int) bool {
		count++
		// Store the value in the found map
		found[v] = struct{}{}

		return true
	}

	// Call the sequence with the yield function
	seq(yield)

	// Check if the count matches the expected number of elements in the set
	if count != 3 {
		t.Errorf("Expected 3 calls of yeild, got %d", count)
	}

	// Check if the keys and values are correct
	expected := map[int]struct{}{1: {}, 2: {}, 3: {}}
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Expected %v, got %v", expected, found)
	}
}

func TestUnkeyedSeqSortedSet(t *testing.T) {
	// Create an UnkeyedIterator
	it := &mockUnkeyedIterator[int]{Values: []int{1, 2, 3}}

	// Create a sequence using UnkeyedSeq
	seq := UnkeyedSeq[int](it)

	// Initialize a counter to track the number of yielded values
	count := 0
	found := make(map[int]struct{}, len(it.Values))

	// Define a yield function that increments the counter
	yield := func(v int) bool {
		count++
		// Store the value in the found map
		found[v] = struct{}{}

		return true
	}

	// Call the sequence with the yield function
	seq(yield)

	// Check if the count matches the expected number of elements in the set
	if count != 3 {
		t.Errorf("Expected 3 calls of yeild, got %d", count)
	}

	// Check if the keys and values are correct
	expected := map[int]struct{}{1: {}, 2: {}, 3: {}}
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Expected %v, got %v", expected, found)
	}
}

// mockIndexedIterator is a mock implementation of [IndexedIterator] for testing purposes.
type mockIndexedIterator[T any] struct {
	// The current index of the iterator
	index int

	// The values to iterate over
	values []T
}

// First implements [IndexedIterator.First].
func (m *mockIndexedIterator[T]) First() {
	m.index = 0
}

// Done implements [IndexedIterator.Done].
func (m *mockIndexedIterator[T]) Done() bool {
	return m.index >= len(m.values)
}

// Next implements [IndexedIterator.Next].
func (m *mockIndexedIterator[T]) Next() (int, T) {
	if m.Done() {
		return -1, *new(T)
	}
	val := m.values[m.index]
	m.index++
	return m.index - 1, val
}

// kvPair is a key-value pair used in the mockKeyedIterator.
type kvPair[K, V any] struct {
	// The Key of the pair
	Key K

	// The value of the pair
	Value V
}

// mockKeyedIterator is a mock implementation of [KeyedIterator] for testing purposes.
type mockKeyedIterator[K, V any] struct {
	// The current Index of the iterator
	Index int

	// The key-value Pairs to iterate over
	Pairs []kvPair[K, V]
}

// First implements [KeyedIterator.First].
func (m *mockKeyedIterator[K, V]) First() {
	m.Index = 0
}

// Done implements [KeyedIterator.Done].
func (m *mockKeyedIterator[K, V]) Done() bool {
	return m.Index >= len(m.Pairs)
}

// Next implements [KeyedIterator.Next].
func (m *mockKeyedIterator[K, V]) Next() (K, V, bool) {
	if m.Done() {
		return *new(K), *new(V), false
	}
	pair := m.Pairs[m.Index]
	m.Index++
	return pair.Key, pair.Value, true
}

// mockUnkeyedIterator is a mock implementation of [UnkeyedIterator] for testing purposes.
type mockUnkeyedIterator[T any] struct {
	// The current Index of the iterator
	Index int

	// The Values to iterate over
	Values []T
}

// First implements [UnkeyedIterator.First].
func (m *mockUnkeyedIterator[T]) First() {
	m.Index = 0
}

// Done implements [UnkeyedIterator.Done].
func (m *mockUnkeyedIterator[T]) Done() bool {
	return m.Index >= len(m.Values)
}

// Next implements [UnkeyedIterator.Next].
func (m *mockUnkeyedIterator[T]) Next() (T, bool) {
	if m.Done() {
		return *new(T), false
	}
	val := m.Values[m.Index]
	m.Index++
	return val, true
}
