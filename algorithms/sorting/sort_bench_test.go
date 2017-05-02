package sorting

import (
	"math/rand"
	"sort"
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

func benchmarkLibSort(b *testing.B, size int) {
	s := makeIntSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Sort(IntSlice(s))
	}
}

func BenchmarkSort1000(b *testing.B)   { benchmarkLibSort(b, 1000) }
func BenchmarkSort10000(b *testing.B)  { benchmarkLibSort(b, 10000) }
func BenchmarkSort100000(b *testing.B) { benchmarkLibSort(b, 100000) }

func BenchmarkBubbleSort1000(b *testing.B) { benchmarkSort(b, BubbleSort, 1000) }

// func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkSort(b, BubbleSort, 10000) }
// func BenchmarkBubbleSort100000(b *testing.B) { benchmarkSort(b, BubbleSort, 100000) }

func BenchmarkQuickSort1000(b *testing.B)   { benchmarkSort(b, QuickSort, 1000) }
func BenchmarkQuickSort10000(b *testing.B)  { benchmarkSort(b, QuickSort, 10000) }
func BenchmarkQuickSort100000(b *testing.B) { benchmarkSort(b, QuickSort, 100000) }

func BenchmarkHeapSort1000(b *testing.B)   { benchmarkSort(b, HeapSort, 1000) }
func BenchmarkHeapSort10000(b *testing.B)  { benchmarkSort(b, HeapSort, 10000) }
func BenchmarkHeapSort100000(b *testing.B) { benchmarkSort(b, HeapSort, 100000) }
