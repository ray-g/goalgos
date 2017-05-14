package number

func IsOdd(n int) bool {
	return n&1 != 0
}

func IsEven(n int) bool {
	return n%2 == 0
}

// GCD Greatest Common Divisor
// Using Euclid's algorithm
// https://en.wikipedia.org/wiki/Greatest_common_divisor
func GCD(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM Least Common Multiple
// Using a*b/GCD(a,b)
// https://en.wikipedia.org/wiki/Least_common_multiple
func LCM(a, b int) int {
	gcd := GCD(a, b)

	if gcd == 0 {
		return 0
	}

	return a * b / gcd
}

// ISqrt Integer square root
// https://en.wikipedia.org/wiki/Integer_square_root
func ISqrt(n int) int {
	x := n
	y := (x + 1) >> 1
	for y < x {
		x = y
		y = (x + n/x) >> 1
	}
	return x
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func MulOverflows64(a, b uint64) bool {
	if a <= 1 || b <= 1 {
		return false
	}
	c := a * b
	return c/b != a
}

func MulOverflows32(a, b uint32) bool {
	if a <= 1 || b <= 1 {
		return false
	}
	c := a * b
	return c/b != a
}

const mostNegative64 = -(mostPositive64 + 1)
const mostPositive64 = 1<<63 - 1

func SignedMulOverflows64(a, b int64) bool {
	if a == 0 || b == 0 || a == 1 || b == 1 {
		return false
	}
	if a == mostNegative64 || b == mostNegative64 {
		return true
	}
	c := a * b
	return c/b != a
}

const mostNegative32 = -(mostPositive32 + 1)
const mostPositive32 = 1<<31 - 1

func SignedMulOverflows32(a, b int32) bool {
	if a == 0 || b == 0 || a == 1 || b == 1 {
		return false
	}
	if a == mostNegative32 || b == mostNegative32 {
		return true
	}
	c := a * b
	return c/b != a
}
