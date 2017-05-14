package rabinkarp

import (
	"hash"
	"math/rand"
	"testing"

	"github.com/ray-g/gobackup/chunker/rollinghash"
)

func NewRollingHash() rollinghash.Hash64 {
	return New()
}

// This is a no-op to prove that we implement hash.Hash64
var _ = hash.Hash64(NewRollingHash())

func sum64ByWriteAndRoll(b []byte) uint64 {
	q := []byte(" ")
	q = append(q, b...)
	roll := NewRollingHash()
	roll.Write(q[:len(q)-1])
	roll.Next(q[len(q)-1])
	return roll.Sum64()
}

func sum64ByWriteOnly(b []byte) uint64 {
	roll := NewRollingHash()
	roll.Write(b)
	return roll.Sum64()
}

func RandomBytes() (res []byte) {
	n := rand.Intn(8192)
	res = make([]byte, n)
	rand.Read(res)
	return res
}

func TestRabinKarp(t *testing.T) {
	for i := 0; i < 1000; i++ {
		in := RandomBytes()
		if len(in) > 0 {
			wr := sum64ByWriteAndRoll(in)
			wo := sum64ByWriteOnly(in)

			if wo != wr {
				t.Errorf("Expected 0x%x, got 0x%x", wo, wr)
			}
		}
	}
}

func BenchmarkRollingKB(b *testing.B) {
	b.SetBytes(1024)
	window := make([]byte, 1024)
	for i := range window {
		window[i] = byte(i)
	}

	h := NewRollingHash()
	in := make([]byte, 0, h.Size())

	b.ResetTimer()
	h.Write(window)
	for i := 0; i < b.N; i++ {
		h.Next(byte(1024 + i))
		h.Sum(in)
	}
}

func BenchmarkRolling128B(b *testing.B) {
	b.SetBytes(1024)
	window := make([]byte, 128)
	for i := range window {
		window[i] = byte(i)
	}

	h := NewRollingHash()
	in := make([]byte, 0, h.Size())

	b.ResetTimer()
	h.Write(window)
	for i := 0; i < b.N; i++ {
		h.Next(byte(128 + i))
		h.Sum(in)
	}
}

func TestPanic(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Did not panic")
			}
		}()
		NewWithPrimeAndSize(65535, 64)
	}()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Did not panic")
			}
		}()
		NewWithPrimeAndSize(53, 10)
	}()
}

func TestSize(t *testing.T) {
	roll := NewRollingHash()
	if 1 != roll.BlockSize() {
		t.Error()
	}

	if 8 != roll.Size() {
		t.Error()
	}
}

func TestSum(t *testing.T) {
	roll := NewRollingHash()
	in := RandomBytes()
	roll.Write(in)
	roll.Sum(in)
}
