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

func TestMax(t *testing.T) {
	if Max(1, 2) != 2 {
		t.Error()
	}

	if Max(1, 1) != 1 {
		t.Error()
	}

	if Max(1, 0) != 1 {
		t.Error()
	}
}

func TestMin(t *testing.T) {
	if Min(1, 2) != 1 {
		t.Error()
	}

	if Min(1, 1) != 1 {
		t.Error()
	}

	if Min(1, 0) != 0 {
		t.Error()
	}
}

func TestMulOverflows64(t *testing.T) {
	cases := []struct {
		a, b    uint64
		expects bool
	}{
		{0, 0, false},
		{1, 1, false},
		{0, 1, false},
		{0, 1, false},
		{1, 0, false},
		{1, 0, false},
		{1, 1<<64 - 1, false},
		{1<<32 - 1, 1<<32 - 1, false},
		{1 << 32, 1 << 32, true},
	}

	for _, tc := range cases {
		actual := MulOverflows64(tc.a, tc.b)
		if actual != tc.expects {
			t.Errorf("%v * %v, expects %v, actual %v.", tc.a, tc.b, tc.expects, actual)
		}
	}
}

func TestMulOverflows32(t *testing.T) {
	cases := []struct {
		a, b    uint32
		expects bool
	}{
		{0, 0, false},
		{1, 1, false},
		{0, 1, false},
		{0, 1, false},
		{1, 0, false},
		{1, 0, false},
		{1, 1<<32 - 1, false},
		{1<<16 - 1, 1<<16 - 1, false},
		{1 << 16, 1 << 16, true},
	}

	for _, tc := range cases {
		actual := MulOverflows32(tc.a, tc.b)
		if actual != tc.expects {
			t.Errorf("%v * %v, expects %v, actual %v.", tc.a, tc.b, tc.expects, actual)
		}
	}
}

func TestSignedMulOverflows64(t *testing.T) {
	cases := []struct {
		a, b    int64
		expects bool
	}{
		{0, 0, false},
		{1, 1, false},
		{0, 1, false},
		{0, 1, false},
		{1, 0, false},
		{1, 0, false},
		{-(1 << 63), 1<<63 - 1, true},
		{1<<63 - 1, -(1 << 63), true},
		{mostNegative64, 1<<63 - 1, true},
		{1<<63 - 1, mostNegative64, true},
		{1, 1<<63 - 1, false},
		{1<<31 - 1, 1<<31 - 1, false},
		{1 << 31, 1 << 31, false},
		{1<<32 - 1, 1<<32 - 1, true},
		{1 << 32, 1 << 32, true},
	}

	for _, tc := range cases {
		actual := SignedMulOverflows64(tc.a, tc.b)
		if actual != tc.expects {
			t.Errorf("%v * %v, expects %v, actual %v.", tc.a, tc.b, tc.expects, actual)
		}
	}
}

func TestSignedMulOverflows32(t *testing.T) {
	cases := []struct {
		a, b    int32
		expects bool
	}{
		{0, 0, false},
		{1, 1, false},
		{0, 1, false},
		{0, 1, false},
		{1, 0, false},
		{1, 0, false},
		{-(1 << 31), 1<<31 - 1, true},
		{1<<31 - 1, -(1 << 31), true},
		{mostNegative32, 1<<31 - 1, true},
		{1<<31 - 1, mostNegative32, true},
		{1, 1<<31 - 1, false},
		{1<<15 - 1, 1<<15 - 1, false},
		{1 << 15, 1 << 15, false},
		{1<<16 - 1, 1<<16 - 1, true},
		{1 << 16, 1 << 16, true},
	}

	for _, tc := range cases {
		actual := SignedMulOverflows32(tc.a, tc.b)
		if actual != tc.expects {
			t.Errorf("%v * %v, expects %v, actual %v.", tc.a, tc.b, tc.expects, actual)
		}
	}
}
