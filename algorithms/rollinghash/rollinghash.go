// Package rollinghash implements several versions of rolling hash
// A rollinghash can be updated byte by byte,
// by specifying which byte enters the window.
// https://en.wikipedia.org/wiki/Rolling_hash
package rollinghash

import (
	"hash"
)

// Roller defines the rolling hash interface
type Roller interface {
	// Next updates the hash of a window from the entering byte.
	Next(b byte)
}

// Hash extends hash.Hash by adding the method Next(b byte).
type Hash interface {
	hash.Hash
	Roller
}

// Hash32 extends hash.Hash32 by adding the method Next(b byte).
type Hash32 interface {
	hash.Hash32
	Roller
}

// Hash64 extends hash.Hash64 by adding the method Next(b byte).
type Hash64 interface {
	hash.Hash64
	Roller
}
