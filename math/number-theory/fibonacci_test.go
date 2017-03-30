package number

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	cases := [][]int{
		[]int{0, 0},
		[]int{1, 1},
		[]int{2, 2},
		[]int{3, 3},
		[]int{4, 5},
		[]int{50, 20365011074},
	}

	for _, vars := range cases {
		v := Fibonacci(vars[0])
		if v != vars[1] {
			t.Errorf("Fibonacci(%d) != %d, returned: %d", vars[0], vars[1], v)
		}
	}
}

func benchmarkFibonacci(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

func BenchmarkFibonacci10(b *testing.B)     { benchmarkFibonacci(b, 10) }
func BenchmarkFibonacci100(b *testing.B)    { benchmarkFibonacci(b, 100) }
func BenchmarkFibonacci1000(b *testing.B)   { benchmarkFibonacci(b, 1000) }
func BenchmarkFibonacci10000(b *testing.B)  { benchmarkFibonacci(b, 10000) }
func BenchmarkFibonacci100000(b *testing.B) { benchmarkFibonacci(b, 100000) }

func ExampleFibonacci() {
	fib := Fibonacci(10)
	fmt.Println(fib)
	// Output:
	// 89
}
