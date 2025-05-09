// immutableiter is a package that provides interfaces for iterators used by [github.com/benbjohnson/immutable] collections.
// It allows converting iterators to functions that can be used with the Go 1.23 iterable function API.
// It also has types for retro-compatibility with [iter.Seq] and [iter.Seq2] from Go 1.23 and [cmp.Ordered] from Go 1.21.
package immutableiter

// Iterator is an interface for iterating over immutable collections.
type Iterator[T any] interface {
	// First positions the iterator on the first index.
	// If source is empty then no change is made.
	First()

	// Done returns true if there are no more elements in the iterator.
	Done() bool
}

// IndexedIterator is an iterator that returns the index and value of the underlying collection.
//
// TODO(Izzette): could this use a better name?
type IndexedIterator[T any] interface {
	Iterator[T]

	// Next returns the current index and its value & moves the iterator forward.
	// Returns an index of -1 if the there are no more elements to return.
	Next() (int, T)
}

// KeyedIterator is an iterator that returns the key and value of the underlying collection.
//
// TODO(Izzette): could this use a better name?
type KeyedIterator[K, V any] interface {
	Iterator[V]

	// Next returns the next key/value pair.
	// Returns a nil key when no elements remain.
	Next() (K, V, bool)
}

// UnkeyedIterator is an iterator that returns the value of the underlying collection.
//
// TODO(Izzette): could this use a better name?
type UnkeyedIterator[T any] interface {
	Iterator[T]

	// Next moves the iterator to the next value.
	Next() (T, bool)
}

// IndexedSeq takes an IndexedIterator for immutable collections (such as those returned by [github.com/benbjohnson/immutable.List.Iterator]) and returns a function satisfying the [iter.Seq]
// type available starting in Go 1.23.
// This can be used to perform range-like operations on the [iter.Seq] and inter-operate with other libraries using the Go 1.23 iterable function API.
//
// TODO(Izzette): could this use a better name?
func IndexedSeq[T any](it IndexedIterator[T]) Seq[T] {
	seq2 := IndexedSeqWithIndex(it)
	return func(yield func(T) bool) {
		seq2(func(_ int, v T) bool {
			return yield(v)
		})
	}
}

// IndexedSeqWithIndex takes an IndexedIterator for immutable collections (such as those returned by [github.com/benbjohnson/immutable.List.Iterator]) and returns a function satisfying the
// [iter.Seq2] type available starting in Go 1.23.
// This can be used to perform range-like operations on the [iter.Seq2] and inter-operate with other libraries using the Go 1.23 iterable function API.
func IndexedSeqWithIndex[T any](it IndexedIterator[T]) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for !it.Done() {
			index, v := it.Next()
			assert(index >= 0, "index should never be negative after checking Done()")

			if !yield(index, v) {
				return
			}
		}
	}
}

// KeyedSeq takes a KeyedIterator for immutable collections (such as those [github.com/benbjohnson/immutable.Map.Iterator] and [github.com/benbjohnson/immutable.OrderedMap.Iterator]) and returns
// a function satisfying the [iter.Seq2] type available starting in Go 1.23.
// This can be used to perform range-like operations on the [iter.Seq2] and inter-operate with other libraries using the Go 1.23 iterable function API.
//
// TODO(Izzette): could this use a better name?
func KeyedSeq[K, V any](it KeyedIterator[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for !it.Done() {
			k, v, ok := it.Next()
			assert(ok, "keyed iterator should always return a key/value pair after checking Done()")

			if !yield(k, v) {
				return
			}
		}
	}
}

// UnkeyedSeq takes an UnkeyedIterator for immutable collections (such as those returned by [github.com/benbjohnson/immutable.Set.Iterator]) and
// [github.com/benbjohnson/immutable.SortedSet.Iterator]) and returns a function satisfying the [iter.Seq] type available starting in Go 1.23.
// This can be used to perform range-like operations on the [iter.Seq] and inter-operate with other libraries using the Go 1.23 iterable function API.
//
// TODO(Izzette): could this use a better name?
func UnkeyedSeq[T any](it UnkeyedIterator[T]) Seq[T] {
	return func(yield func(T) bool) {
		for !it.Done() {
			v, ok := it.Next()
			assert(ok, "unkeyed iterator should always return a value after checking Done()")

			if !yield(v) {
				return
			}
		}
	}
}

// assert is a helper function that panics with the given message if the condition is false.
func assert(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}
