package math

import "errors"

// GCD Greatest Common Divisor
// Using Euclid's algorithm
// https://en.wikipedia.org/wiki/Greatest_common_divisor
func GCD(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("Must be non-zero number")
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a, nil
}

// LCM Least Common Multiple
// Using a*b/GCD(a,b)
// https://en.wikipedia.org/wiki/Least_common_multiple
func LCM(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("Must be non-zero number")
	}

	gcd, _ := GCD(a, b)

	return a * b / gcd, nil
}
