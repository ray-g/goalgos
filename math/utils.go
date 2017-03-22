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

func IsPrime(n int) bool {
	if n == 2 {
		return true
	}

	if n < 2 || n&1 == 0 {
		return false
	}

	for i := 3; i <= ISqrt(n); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func LargestPrimeBelow(n int) (p int, e error) {
	if n < 2 {
		return 0, errors.New("Number must larger than 2")
	}

	p = 0
	e = nil

	if IsPrime(n) {
		p = n
		return
	}

	d := n % 6
	i := n - d

	if d == 5 && IsPrime(i+5) {
		p = i + 5
		return
	}

	if d >= 1 && IsPrime(i+1) {
		p = i + 1
		return
	}

	if n == 4 {
		p = 3
		return
	}

	for ; i > 0; i -= 6 {
		if IsPrime(i - 1) {
			p = i - 1
			return
		}

		if IsPrime(i - 5) {
			p = i - 5
			return
		}
	}
	return
}

func LargestPrimeFactor(n int) (p int, e error) {
	if n < 2 {
		return 0, errors.New("Number must larger than 2")
	}

	p = 1
	e = nil

	if IsPrime(n) {
		p = n
		return
	}

	if n%2 == 0 {
		return LargestPrimeFactor(n / 2)
	}

	if n%3 == 0 {
		return LargestPrimeFactor(n / 3)
	}

	if n%5 == 0 {
		return LargestPrimeFactor(n / 5)
	}

	for i := 6; i < ISqrt(n); i += 6 {
		x := i + 1
		if IsPrime(x) && n%x == 0 {
			return LargestPrimeFactor(n / x)
		}

		x = i + 5
		if IsPrime(x) && n%x == 0 {
			return LargestPrimeFactor(n / x)
		}
	}
	return
}
