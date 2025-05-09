package immutable_test

import (
	"fmt"

	"github.com/benbjohnson/immutable"
)

// TODO(Izzette): add examples for the All functions of the SortedMap and SortedSet types.
// TODO(Izzette): add examples which prematurely break the iteration.

// Demonstrates how to use the [immutable.List.All] function to iterate over an [immutable.List] of integers.
//
// In Go 1.23 with GOEXPERIMENT=rangefunc, or in Go 1.24 and later, this will be a range-like operation:
//
//	for v := range seq {
//		fmt.Println("Yielded:", v)
//	}
func ExampleList_All() {
	// Create a new list with some values
	l := immutable.NewList(1, 2, 3, 4, 5)

	// Get a sequence from the list
	seq := l.All()

	// Define a yield function that increments the counter
	yield := func(v int) bool {
		fmt.Println("Yielded:", v)
		return true
	}

	seq(yield)
	// Output:
	// Yielded: 1
	// Yielded: 2
	// Yielded: 3
	// Yielded: 4
	// Yielded: 5
}

// Demonstrates how to use the [immutable.List.All] function to iterate over an [immutable.List] of integers with their indices.
//
// In Go 1.23 with GOEXPERIMENT=rangefunc, or in Go 1.24 and later, this will be a range-like operation:
//
//	for i, v := range seq {
//		fmt.Printf("Yielded: %d at index %d\n", v, i)
//	}
func ExampleList_AllWithIndex() {
	// Create a new list with some values
	l := immutable.NewList(1, 2, 3, 4, 5)

	// Create a sequence from the list
	seq := l.AllWithIndex()

	// Define a yield function that increments the counter
	yield := func(i int, v int) bool {
		fmt.Printf("Yielded: %d at index %d\n", v, i)
		return true
	}

	seq(yield)
	// Output:
	// Yielded: 1 at index 0
	// Yielded: 2 at index 1
	// Yielded: 3 at index 2
	// Yielded: 4 at index 3
	// Yielded: 5 at index 4
}

// Demonstrates how to use the [immutable.Map.All] function to iterate over an [immutable.Map] of integers to strings.
//
// In Go 1.23 with GOEXPERIMENT=rangefunc, or in Go 1.24 and later, this will be a range-like operation:
//
//	for k, v := range seq {
//		fmt.Printf("Yielded: %s at key %d\n", v, k)
//	}
func ExampleMap_All() {
	// Create a new map builder with some values
	mb := immutable.NewMapBuilder[int, string](nil)
	mb.Set(1, "one")
	mb.Set(2, "two")
	mb.Set(3, "three")

	// Create a map from the builder
	m := mb.Map()

	// Create a sequence from the map
	seq := m.All()

	// Define a yield function that increments the counter
	yield := func(k int, v string) bool {
		fmt.Printf("Yielded: %s at key %d\n", v, k)
		return true
	}

	seq(yield)
	// Output:
	// Yielded: one at key 1
	// Yielded: two at key 2
	// Yielded: three at key 3
}

// Demonstrates how to use the [immutable.Set.All] function to iterate over an [immutable.Set] of integers.
//
// In Go 1.23 with GOEXPERIMENT=rangefunc, or in Go 1.24 and later, this will be a range-like operation:
//
//	for v := range seq {
//		fmt.Println("Yielded:", v)
//	}
func ExampleSet_All() {
	// Create a new set with some values
	s := immutable.NewSet(nil, 1, 2, 3)

	// Create a Sequence from the set
	seq := s.All()

	// Define a yield function that increments the counter
	yield := func(v int) bool {
		fmt.Println("Yielded:", v)
		return true
	}

	seq(yield)
	// Output:
	// Yielded: 1
	// Yielded: 2
	// Yielded: 3
}
