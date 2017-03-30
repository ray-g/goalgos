package number

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53}
	for _, p := range primes {
		if !IsPrime(p) {
			t.Errorf("Number %d shows not a prime", p)
		}
	}

	notprimes := []int{6, 4, 10, 21, 22, 25, 100, 51, 81, 121, 999997}
	for _, np := range notprimes {
		if IsPrime(np) {
			t.Errorf("Number %d shows as a prime", np)
		}
	}
}

func TestSieveOfEratosthenes(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53}
	results := SieveOfEratosthenes(55)
	for i, v := range results {
		if primes[i] != v {
			t.Errorf("%dth results %d not same as %d", i, v, primes[i])
		}
	}

	results = SieveOfEratosthenes(1)
	if results != nil {
		t.Error()
	}

	results = SieveOfEratosthenes(2)
	if len(results) != 1 {
		t.Error()
	}
}

func ExampleSieveOfEratosthenes() {
	primes := SieveOfEratosthenes(55)
	fmt.Println(primes)

	// Output:
	// [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53]
}

func TestLargestPrimeBelow(t *testing.T) {
	cases := [][]int{
		[]int{0, 0},
		[]int{1, 0},
		[]int{2, 2},
		[]int{3, 3},
		[]int{4, 3},
		[]int{5, 5},
		[]int{100, 97},
		[]int{99999999, 99999989},
		[]int{100000000, 99999989},
		[]int{10000000, 9999991},
		[]int{1000000, 999983},
		[]int{100000, 99991},
		[]int{10000, 9973},
		[]int{1000, 997},
	}

	for _, vars := range cases {
		v := LargestPrimeBelow(vars[0])
		if v != vars[1] {
			t.Errorf("LargestPrimeBelow(%d) != %d, returned: %d", vars[0], vars[1], v)
		}
	}
}

func TestLargestPrimeFactor(t *testing.T) {
	cases := [][]int{
		[]int{0, 0},
		[]int{1, 0},
		[]int{2, 2},
		[]int{3, 3},
		[]int{4, 2},
		[]int{5, 5},
		[]int{25, 5},
		[]int{9, 3},
		[]int{11 * 13 * 97, 97},
	}

	for _, vars := range cases {
		v := LargestPrimeFactor(vars[0])
		if v != vars[1] {
			t.Errorf("LargestPrimeFactor(%d) != %d, returned: %d", vars[0], vars[1], v)
		}
	}
}
