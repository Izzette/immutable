package immutable

import (
	"reflect"
	"sort"
	"testing"

	"github.com/benbjohnson/immutable/immutableiter"
)

// TODO(Izzette): add tests for the All function of the Map, SortedMap, Set, and SortedSet types and their corresponding builders.

func TestList_All(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		seqTestEmpty(t, func() immutableiter.Seq[int] {
			return NewList[int]().All()
		})
	})
	t.Run("One", func(t *testing.T) {
		seqTestOne(t, func() immutableiter.Seq[int] {
			return NewList[int](1).All()
		}, 1)
	})
	t.Run("OneBreak", func(t *testing.T) {
		seqTestOneBreak(t, func() immutableiter.Seq[int] {
			return NewList[int](1).All()
		}, 1)
	})
	t.Run("AllOrdered", func(t *testing.T) {
		seqTestAllOrdered(t, func() immutableiter.Seq[int] {
			return NewList[int](1, 2, 3).All()
		}, []int{1, 2, 3})
	})
	t.Run("AllOrderedBreak", func(t *testing.T) {
		seqTestAllOrderedBreak(t, func() immutableiter.Seq[int] {
			return NewList[int](1, 2, 3, 4, 5).All()
		}, []int{1, 2, 3})
	})
}

func TestListBuilder_All(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		seqTestEmpty(t, func() immutableiter.Seq[int] {
			return NewListBuilder[int]().All()
		})
	})
	t.Run("One", func(t *testing.T) {
		seqTestOne(t, func() immutableiter.Seq[int] {
			lb := NewListBuilder[int]()
			lb.Append(1)
			return lb.All()
		}, 1)
	})
	t.Run("OneBreak", func(t *testing.T) {
		seqTestOneBreak(t, func() immutableiter.Seq[int] {
			lb := NewListBuilder[int]()
			lb.Append(1)
			return lb.All()
		}, 1)
	})
	t.Run("AllOrdered", func(t *testing.T) {
		seqTestAllOrdered(t, func() immutableiter.Seq[int] {
			lb := NewListBuilder[int]()
			lb.Append(1)
			lb.Append(2)
			lb.Append(3)
			return lb.All()
		}, []int{1, 2, 3})
	})
	t.Run("AllOrderedBreak", func(t *testing.T) {
		seqTestAllOrderedBreak(t, func() immutableiter.Seq[int] {
			lb := NewListBuilder[int]()
			for i := 1; i <= 5; i++ {
				lb.Append(i)
			}
			return lb.All()
		}, []int{1, 2, 3})
	})
}

// seqTestEmpty checks if the sequence returned by allFunc is empty.
func seqTestEmpty[T any](t *testing.T, allFunc func() immutableiter.Seq[T]) {
	seq := allFunc()

	actual := make([]T, 0)
	yield := func(v T) bool {
		actual = append(actual, v)
		return true
	}
	seq(yield)

	if len(actual) != 0 {
		t.Errorf("Expected empty sequence, got %v", actual)
	}
}

// seqTestOne checks if the sequence returned by allFunc contains exactly one
// element and that it matches the expected value.
func seqTestOne[T any](t *testing.T, allFunc func() immutableiter.Seq[T], expected T) {
	seq := allFunc()

	actual := make([]T, 0)
	yield := func(v T) bool {
		actual = append(actual, v)
		return true
	}
	seq(yield)

	if len(actual) != 1 {
		t.Errorf("Expected 1 element, got %d", len(actual))
	}
	if !reflect.DeepEqual(actual, []T{expected}) {
		t.Errorf("Expected %v, got %v", expected, actual[0])
	}
}

// seqTestOneBreak checks if the sequence returned by allFunc contains exactly
// one element and that it matches the expected value, but stops yielding after
// the first element.
func seqTestOneBreak[T any](t *testing.T, allFunc func() immutableiter.Seq[T], expected T) {
	seq := allFunc()

	actual := make([]T, 0)
	yield := func(v T) bool {
		actual = append(actual, v)
		// Do not continue yielding, but this is the last element so it should
		// not impact the result.
		return false
	}
	seq(yield)

	if len(actual) != 1 {
		t.Errorf("Expected 1 element, got %d", len(actual))
	}
	if !reflect.DeepEqual(actual, []T{expected}) {
		t.Errorf("Expected %v, got %v", expected, actual[0])
	}
}

// seqTestAllOrdered checks if the sequence returned by allFunc contains all
// elements in the expected order.
// It compares the actual sequence with the expected sequence.
func seqTestAllOrdered[T any](
	t *testing.T,
	allFunc func() immutableiter.Seq[T],
	expected []T,
) {
	seq := allFunc()

	actual := make([]T, 0, len(expected))
	yield := func(v T) bool {
		actual = append(actual, v)
		return true
	}
	seq(yield)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// seqTestAllUnordered checks if the sequence returned by allFunc contains all
// elements in any order.
// It compares the actual sequence with the expected sequence after sorting
// both slices.
func seqTestAllUnordered[T immutableiter.Ordered](
	t *testing.T,
	allFunc func() immutableiter.Seq[T],
	expected []T,
) {
	seq := allFunc()

	actual := make([]T, 0, len(expected))
	yield := func(v T) bool {
		actual = append(actual, v)
		return true
	}
	seq(yield)

	// Sort both slices for comparison
	sort.Slice(actual, func(i, j int) bool {
		return actual[i] < actual[j]
	})
	sort.Slice(expected, func(i, j int) bool {
		return expected[i] < expected[j]
	})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// seqTestAllOrderedBreak checks if the sequence returned by allFunc contains
// all elements in the expected order, but stops yielding after the nth element
// where n is the length of the expected slice.
func seqTestAllOrderedBreak[T any](
	t *testing.T,
	allFunc func() immutableiter.Seq[T],
	expected []T,
) {
	seq := allFunc()

	actual := make([]T, 0, len(expected))
	yield := func(v T) bool {
		actual = append(actual, v)
		return len(actual) < len(expected)
	}
	seq(yield)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// seqTestAllSet checks if the sequence returned by allFunc contains all
// elements in any order, but ensures that no element is returned more than
// once.
// It stops yielding after the nth element where n is the sample parameter.
func seqTestAllSetBreak[T comparable](
	t *testing.T,
	allFunc func() immutableiter.Seq[T],
	sample int,
	expected []T,
) {
	seq := allFunc()

	actual := make([]T, 0, len(expected))
	yield := func(v T) bool {
		actual = append(actual, v)
		return len(actual) < sample
	}
	seq(yield)

	// Convert expected slice to a map for comparison
	expectedMap := make(map[T]int, len(expected))
	for _, v := range expected {
		expectedMap[v] = 0
	}
	// Check if all yielded values are in the expected map no more than once
	for _, v := range actual {
		if count, ok := expectedMap[v]; !ok {
			t.Errorf("Unexpected value %v in yielded values", v)
		} else if count > 0 {
			t.Errorf("Duplicate value %v in yielded values", v)
		} else {
			expectedMap[v]++
		}
	}
}
