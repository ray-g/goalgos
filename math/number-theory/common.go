package number

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
	if a == 0 || b == 0 {
		return 0
	}

	gcd := GCD(a, b)

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
