package immutable

// Iterator is an interface for iterating over immutable collections.
type Iterator[T any] interface {
	// First positions the iterator on the first index.
	// If source is empty then no change is made.
	First()

	// Done returns true if there are no more elements in the iterator.
	Done() bool
}

// IndexedIterator is an iterator that returns the index and value of the
// underlying collection.
type IndexedIterator[T any] interface {
	Iterator[T]

	// Next returns the current index and its value & moves the iterator forward.
	// Returns an index of -1 if the there are no more elements to return.
	Next() (int, T)
}

// KeyedIterator is an iterator that returns the key and value of the underlying
// collection.
type KeyedIterator[K, V any] interface {
	Iterator[V]

	// Next returns the next key/value pair.
	// Returns a nil key when no elements remain.
	Next() (K, V, bool)
}

// UnkeyedIterator is an iterator that returns the value of the underlying
// collection.
type UnkeyedIterator[T any] interface {
	Iterator[T]

	// Next moves the iterator to the next value.
	Next() (T, bool)
}

// IndexedSeq takes an IndexedIterator for immutable collections (such as those
// returned by [List.Iterator]) and returns a function satisfying the [iter.Seq]
// type available starting in Go 1.23.
// This can be used to perform range-like operations on the [iter.Seq] and
// inter-operate with other libraries using the Go 1.23 iterable function API.
func IndexedSeq[T any](it IndexedIterator[T]) func(func(T) bool) {
	seq2 := IndexedSeqWithIndex(it)
	return func(yield func(T) bool) {
		seq2(func(_ int, v T) bool {
			return yield(v)
		})
	}
}

// IndexedSeqWithIndex takes an IndexedIterator for immutable collections (such
// as those returned by [List.Iterator]) and returns a function satisfying the
// [iter.Seq2] type available starting in Go 1.23.
// This can be used to perform range-like operations on the [iter.Seq2] and
// inter-operate with other libraries using the Go 1.23 iterable function API.
func IndexedSeqWithIndex[T any](it IndexedIterator[T]) func(func(int, T) bool) {
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

// KeyedSeq takes a KeyedIterator for immutable collections (such as those
// [Map.Iterator] and [OrderedMap.Iterator]) and returns a function satisfying
// the [iter.Seq2] type available starting in Go 1.23.
// This can be used to perform range-like operations on the [iter.Seq2] and
// inter-operate with other libraries using the Go 1.23 iterable function API.
func KeyedSeq[K, V any](it KeyedIterator[K, V]) func(func(K, V) bool) {
	return func(yield func(K, V) bool) {
		for !it.Done() {
			k, v, ok := it.Next()
			assert(
				ok,
				"keyed iterator should always return a key/value pair after checking "+
					"Done()",
			)

			if !yield(k, v) {
				return
			}
		}
	}
}

// UnkeyedSeq takes an UnkeyedIterator for immutable collections (such as those
// returned by [List.Iterator]) and returns a function satisfying the [iter.Seq]
// type available starting in Go 1.23.
// This can be used to perform range-like operations on the [iter.Seq] and
// inter-operate with other libraries using the Go 1.23 iterable function API.
func UnkeyedSeq[T any](it UnkeyedIterator[T]) func(func(T) bool) {
	return func(yield func(T) bool) {
		for !it.Done() {
			v, ok := it.Next()
			assert(
				ok, "unkeyed iterator should always return a value after checking Done()")

			if !yield(v) {
				return
			}
		}
	}
}
