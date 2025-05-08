package immutable

import (
	"reflect"
	"testing"
)

func TestIndexedSeq(t *testing.T) {
	// Create a new list with some values
	l := NewList(1, 2, 3, 4, 5)

	// Create an IndexedIterator for the list
	it := l.Iterator()

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
	// Create a new list with some values
	l := NewList(1, 2, 3, 4, 5)

	// Create an IndexedIterator for the list
	it := l.Iterator()

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
	// Create a new Map with some key-value pairs
	mb := NewMapBuilder[string, int](nil)
	mb.Set("a", 1)
	mb.Set("b", 2)
	mb.Set("c", 3)
	m := mb.Map()

	// Create a KeyedIterator for the map
	it := m.Iterator()

	// Create a sequence using KeyedSeq
	seq := KeyedSeq[string, int](it)

	// Initialize a counter to track the number of yielded values
	count := 0
	found := make(map[string]int, m.Len())

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
	// Create a new SortedMap with some key-value pairs
	mb := NewSortedMapBuilder[string, int](nil)
	mb.Set("a", 1)
	mb.Set("b", 2)
	mb.Set("c", 3)
	m := mb.Map()

	// Create a KeyedIterator for the map
	it := m.Iterator()

	// Create a sequence using KeyedSeq
	seq := KeyedSeq[string, int](it)

	// Initialize a counter to track the number of yielded values
	count := 0
	found := make(map[string]int, m.Len())

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
	// Create a new set with some values
	s := NewSet(nil, 1, 2, 3)

	// Create an UnkeyedIterator for the set
	it := s.Iterator()

	// Create a sequence using UnkeyedSeq
	seq := UnkeyedSeq[int](it)

	// Initialize a counter to track the number of yielded values
	count := 0
	found := make(map[int]struct{}, s.Len())

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
	// Create a new sorted set with some values
	s := NewSortedSet(nil, 1, 2, 3)

	// Create an UnkeyedIterator for the set
	it := s.Iterator()

	// Create a sequence using UnkeyedSeq
	seq := UnkeyedSeq[int](it)

	// Initialize a counter to track the number of yielded values
	count := 0
	found := make(map[int]struct{}, s.Len())

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
