package number

import "testing"

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
