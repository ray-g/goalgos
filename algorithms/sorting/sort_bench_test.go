package sorting

import (
	"math/rand"
	"testing"
)

func makeIntSlice(size int) []int {
	return rand.Perm(size)
}

func benchmarkSort(b *testing.B, foo func(s Sortable), size int) {
	s := makeIntSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		foo(IntSlice(s))
	}

}

func BenchmarkBubbleSort1000(b *testing.B) { benchmarkSort(b, BubbleSort, 1000) }
