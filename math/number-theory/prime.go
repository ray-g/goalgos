package number

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

func SieveOfEratosthenes(n int) []int {
	if n < 2 {
		return nil
	}
	nums := make([]bool, n)
	for i := 2; i <= ISqrt(n); i++ {
		if nums[i] == false {
			for j := i * i; j < n; j += i {
				nums[j] = true
			}
		}
	}
	primes := []int{2}
	for i := 3; i < n; i += 2 {
		if nums[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}

func LargestPrimeBelow(n int) (p int) {
	if n < 2 {
		return 0
	}

	p = 0

	if IsPrime(n) {
		p = n
		return
	}

	d := n % 6
	i := n - d

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
			break
		}

		if IsPrime(i - 5) {
			p = i - 5
			break
		}
	}
	return
}

func LargestPrimeFactor(n int) (p int) {
	if n < 2 {
		return 0
	}

	p = 0

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
			p = LargestPrimeFactor(n / x)
			break
		}

		x = i + 5
		if IsPrime(x) && n%x == 0 {
			p = LargestPrimeFactor(n / x)
			break
		}
	}
	return
}
