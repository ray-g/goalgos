package number

import "errors"

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
