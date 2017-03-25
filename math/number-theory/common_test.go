package number

import "testing"

func TestIsOdd(t *testing.T) {
	b := true
	for n := -31; n < 32; n++ {
		if IsOdd(n) != b {
			t.Errorf("IsOdd(%d) returned: %v", n, !b)
		}
		b = !b
	}
}

func TestIsEven(t *testing.T) {
	b := true
	for n := -32; n < 32; n++ {
		if IsEven(n) != b {
			t.Errorf("IsEven(%d) returned: %v", n, !b)
		}
		b = !b
	}
}

func TestGCD(t *testing.T) {
	cases := [][]int{
		[]int{0, 5, 0},
		[]int{5, 0, 0},
		[]int{5, 3, 1},
		[]int{6, 4, 2},
	}

	for _, vars := range cases {
		v := GCD(vars[0], vars[1])
		if v != vars[2] {
			t.Errorf("GCD(%d, %d) != %d, returned: %d", vars[0], vars[1], vars[2], v)
		}
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(123456789, 987654321)
	}
}

func TestLCM(t *testing.T) {
	cases := [][]int{
		[]int{0, 5, 0},
		[]int{5, 0, 0},
		[]int{5, 3, 15},
		[]int{6, 4, 12},
	}

	for _, vars := range cases {
		v := LCM(vars[0], vars[1])
		if v != vars[2] {
			t.Errorf("LCM(%d, %d) != %d, returned: %d", vars[0], vars[1], vars[2], v)
		}
	}
}

func BenchmarkLCM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LCM(123456789, 987654321)
	}
}

func TestISqrt(t *testing.T) {
	cases := [][]int{
		[]int{0, 0},
		[]int{5, 2},
		[]int{4, 2},
		[]int{100, 10},
		[]int{99, 9},
		[]int{100000000, 10000},
		[]int{99999999, 9999},
	}

	for _, vars := range cases {
		v := ISqrt(vars[0])
		if v != vars[1] {
			t.Errorf("ISqrt(%d) != %d, returned: %d", vars[0], vars[1], v)
		}
	}
}

func BenchmarkISqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ISqrt(123456789)
	}
}
