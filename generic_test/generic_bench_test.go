package generic_test

import "testing"

func filterInPlace[T any](collection []T, test func(T) bool) []T {
	var position, size = 0, len(collection)
	for i := 0; i < size; i++ {
		if test(collection[i]) {
			collection[position] = collection[i]
			position++
		}
	}

	return collection[:position]
}

const (
	CollectionSize = 1000
	ElementSize    = 3
)

func BenchmarkFilter_Typed_Copying(b *testing.B) {
	c := generateStringCollection(CollectionSize, ElementSize)

	b.Run("Equals to AAA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			filterStrings(c, func(s string) bool { return s == "AAA" })
		}
	})
}

func BenchmarkFilter_Generics_Copying(b *testing.B) {
	c := generateStringCollection(CollectionSize, ElementSize)

	b.Run("Equals to AAA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			filter(c, func(s string) bool { return s == "AAA" })
		}
	})
}

func BenchmarkFilter_Generics_Inplace(b *testing.B) {
	c := generateStringCollection(CollectionSize, ElementSize)

	b.Run("Equals to AAA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			filterInPlace(c, func(s string) bool { return s == "AAA" })
		}
	})
}
