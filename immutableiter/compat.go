package immutableiter

type (
	// Seq has the signature of the [iter.Seq] type available starting in Go 1.23.
	Seq[T any] func(func(T) bool)

	// Seq2 has the signature of the [iter.Seq2] type available starting in Go 1.23.
	Seq2[K, V any] func(func(K, V) bool)
)

// Ordered has the signature of the [cmp.Ordered] type available starting in Go 1.21.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}
