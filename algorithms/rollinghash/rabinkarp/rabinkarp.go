// Package rabinkarp implements the Rabin-Karp rolling hash algorithm
package rabinkarp

import (
	"container/ring"

	"github.com/ray-g/gobackup/chunker/rollinghash"
)

// Size of hash value in bytes
const Size = 8

// digest represents the partial evaluation of a checksum.
type digest struct {
	prime uint64
	value uint64

	// window is ring buffer
	windowSize int
	buffer     *ring.Ring

	// Lookup tables
	out [256]uint64
}

// Reset resets the Hash to its initial state.
func (d *digest) Reset() {
	d.value = 0
	d.buffer = ring.New(d.windowSize)

	primeNthPower := uint64(1)

	for i := 0; i < d.windowSize; i++ {
		d.buffer.Value = byte(0)
		d.buffer = d.buffer.Next()

		if rollinghash.MulOverflows64(primeNthPower, d.prime) {
			panic("multiplication overflow uint64 detected")
		}
		primeNthPower *= d.prime
	}

	for i := uint64(0); i < 256; i++ {
		if rollinghash.MulOverflows64(primeNthPower, i) {
			panic("multiplication overflow uint64 detected")
		}
		d.out[i] = primeNthPower * i
	}
}

func NewWithPrimeAndSize(prime, windowSize int) rollinghash.Hash64 {
	p := rollinghash.LargestPrimeBelow(prime)
	d := &digest{
		prime:      uint64(p),
		windowSize: windowSize,
	}
	d.Reset()
	return d
}

// 53 Power 8 won't overflow uint64
func New() rollinghash.Hash64 {
	return NewWithPrimeAndSize(53, 8)
}

// Size returns the number of bytes Sum will return.
func (d *digest) Size() int { return Size }

// BlockSize returns the hash's underlying block size.
func (d *digest) BlockSize() int { return 1 }

func (d *digest) AddByte(in byte) {
	out := d.buffer.Value.(byte)
	d.buffer.Value = in
	d.buffer = d.buffer.Next()

	// Update rolling hash
	d.value *= d.prime
	d.value = (d.value + uint64(in)) - d.out[out]
}

func (d *digest) AddBytes(in []byte) {
	for _, b := range in {
		d.AddByte(b)
	}
}

// Write (via the embedded io.Writer interface) adds more data to the
// running hash. It never returns an error.
func (d *digest) Write(data []byte) (int, error) {
	d.AddBytes(data)
	return len(data), nil
}

func (d *digest) Sum64() uint64 {
	return d.value
}

// appends four bytes if the hash will fit, else 8 bytes
func (d *digest) Sum(b []byte) []byte {
	s := d.Sum64()
	return append(b, byte(s>>56), byte(s>>48), byte(s>>40), byte(s>>32), byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
}

// Next updates the checksum of the window from the leaving byte and the
// entering byte.
func (d *digest) Next(b byte) {
	d.AddByte(b)
}
