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

func benchmarkSortInt(b *testing.B, foo func(s []int), size int) {
	s := makeIntSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		foo(s)
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

// func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkSort(b, BubbleSort, 10000) }
// func BenchmarkBubbleSort100000(b *testing.B) { benchmarkSort(b, BubbleSort, 100000) }
func BenchmarkBubbleSort1000(b *testing.B) { benchmarkSort(b, BubbleSort, 1000) }

func BenchmarkQuickSort1000(b *testing.B)   { benchmarkSort(b, QuickSort, 1000) }
func BenchmarkQuickSort10000(b *testing.B)  { benchmarkSort(b, QuickSort, 10000) }
func BenchmarkQuickSort100000(b *testing.B) { benchmarkSort(b, QuickSort, 100000) }

func BenchmarkHeapSort1000(b *testing.B)   { benchmarkSort(b, HeapSort, 1000) }
func BenchmarkHeapSort10000(b *testing.B)  { benchmarkSort(b, HeapSort, 10000) }
func BenchmarkHeapSort100000(b *testing.B) { benchmarkSort(b, HeapSort, 100000) }

// func BenchmarkSelectionSort10000(b *testing.B) { benchmarkSort(b, SelectionSort, 10000) }
// func BenchmarkSelectionSort100000(b *testing.B) { benchmarkSort(b, SelectionSort, 100000) }
func BenchmarkSelectionSort1000(b *testing.B) { benchmarkSort(b, SelectionSort, 1000) }

//func BenchmarkInsertionSort100000(b *testing.B) { benchmarkSort(b, InsertionSort, 100000) }
func BenchmarkInsertionSort1000(b *testing.B)  { benchmarkSort(b, InsertionSort, 1000) }
func BenchmarkInsertionSort10000(b *testing.B) { benchmarkSort(b, InsertionSort, 10000) }

func BenchmarkShellSort1000(b *testing.B)   { benchmarkSort(b, ShellSort, 1000) }
func BenchmarkShellSort10000(b *testing.B)  { benchmarkSort(b, ShellSort, 10000) }
func BenchmarkShellSort100000(b *testing.B) { benchmarkSort(b, ShellSort, 100000) }

func BenchmarkMergeSortDown1000(b *testing.B)   { benchmarkSortInt(b, MergeSortDown, 1000) }
func BenchmarkMergeSortDown10000(b *testing.B)  { benchmarkSortInt(b, MergeSortDown, 10000) }
func BenchmarkMergeSortDown100000(b *testing.B) { benchmarkSortInt(b, MergeSortDown, 100000) }

func BenchmarkMergeSortUp1000(b *testing.B)   { benchmarkSortInt(b, MergeSortUp, 1000) }
func BenchmarkMergeSortUp10000(b *testing.B)  { benchmarkSortInt(b, MergeSortUp, 10000) }
func BenchmarkMergeSortUp100000(b *testing.B) { benchmarkSortInt(b, MergeSortUp, 100000) }
