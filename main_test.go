package hdtest

import (
	"math/rand"
	"testing"
)

func generateRandomSig() (sig [4]uint64) {
	for x := range sig {
		sig[x] = uint64(rand.Uint64())
	}
	return sig
}

func BenchmarkHammingDist(b *testing.B) {

	sigA := generateRandomSig()
	sigB := generateRandomSig()

	b.Run("256-bit Slice Hamming distance", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64(sigA[:], sigB[:])
		}
	})

	b.Run("256-bit Array Hamming distance", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist4Array64(&sigA, &sigB)
		}
	})
}
