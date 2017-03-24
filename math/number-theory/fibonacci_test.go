package number

import "testing"

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

func benchmarkFibonacci(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

func BenchmarkFibonacci10(b *testing.B)     { benchmarkFibonacci(10, b) }
func BenchmarkFibonacci100(b *testing.B)    { benchmarkFibonacci(100, b) }
func BenchmarkFibonacci1000(b *testing.B)   { benchmarkFibonacci(1000, b) }
func BenchmarkFibonacci10000(b *testing.B)  { benchmarkFibonacci(10000, b) }
func BenchmarkFibonacci100000(b *testing.B) { benchmarkFibonacci(100000, b) }
