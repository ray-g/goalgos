package number

// https://en.wikipedia.org/wiki/Fibonacci_number

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}

	x, y := 1, 1
	for i := 1; i < n; i++ {
		x, y = x+y, x
	}

	return x
}
