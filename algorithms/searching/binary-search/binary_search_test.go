package binarysearch

import (
	"testing"
)

func makeSortedArr(size int) []int {
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = 2 * i
	}

	return arr
}

func TestBinarySearch(t *testing.T) {
	size := 10000
	arr := makeSortedArr(size)

	for i := 0; i < size; i++ {
		ret := search(arr, 2*i)

		if ret != i {
			t.Errorf("expects: arr[%d] = %d, got arr[%d] = %d", i, 2*i, ret, arr[ret])
		}
	}

	ret := search(arr, 3)
	if ret != -1 {
		t.Errorf("Should not got this arr[%d] = %d", ret, arr[ret])
	}
}

func benchmarkBinarySearch(b *testing.B, size int) {
	arr := makeSortedArr(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search(arr, 0)
	}
}

func BenchmarkBinarySearch1000(b *testing.B)      { benchmarkBinarySearch(b, 1000) }
func BenchmarkBinarySearch10000(b *testing.B)     { benchmarkBinarySearch(b, 10000) }
func BenchmarkBinarySearch100000(b *testing.B)    { benchmarkBinarySearch(b, 100000) }
func BenchmarkBinarySearch1000000(b *testing.B)   { benchmarkBinarySearch(b, 1000000) }
func BenchmarkBinarySearch10000000(b *testing.B)  { benchmarkBinarySearch(b, 10000000) }
func BenchmarkBinarySearch100000000(b *testing.B) { benchmarkBinarySearch(b, 100000000) }
