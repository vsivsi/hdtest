package hdtest

import (
	"fmt"
	"math/rand"
	"testing"
)

func generateRandomSig() (sig [4]uint64) {
	for x := range sig {
		sig[x] = uint64(rand.Uint64())
	}
	return sig
}

func BenchmarkHD(b *testing.B) {

	sigA := generateRandomSig()
	sigB := generateRandomSig()

	b.Run("256-bit unrolled array HD", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist4Array64Unroll(&sigA, &sigB)
		}
	})

	b.Run("256-bit slice HD", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64(sigA[:], sigB[:])
		}
	})

	b.Run("256-bit unrolled slice HD", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64Unroll4Way(sigA[:], sigB[:])
		}
	})

	b.Run("256-bit unrolled 4 el slice HD", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist4Slice64Unroll(sigA[:], sigB[:])
		}
	})
}

func generateRandomSliceSig(n int) (sig []uint64) {
	sig = make([]uint64, n)
	for x := range sig {
		sig[x] = uint64(rand.Uint64())
	}
	return sig
}

func BenchmarkSliceHD(b *testing.B) {

	const max = 100

	sigA := generateRandomSliceSig(max)
	sigB := generateRandomSliceSig(max)

	testLens := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		12, 15, 18, 21, 24, 27, 30, 40, 50, 75, 100}

	for _, v := range testLens {
		A := sigA[:v]
		B := sigB[:v]

		fmt.Println()

		b.Run(fmt.Sprintf("%d-bit loop slice HD", v*64), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HammingDistSlice64(A, B)
			}
		})
		b.Run(fmt.Sprintf("%d-bit 3-way unrolled slice HD", v*64), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HammingDistSlice64Unroll3Way(A, B)
			}
		})
		b.Run(fmt.Sprintf("%d-bit 4-way unrolled slice HD", v*64), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HammingDistSlice64Unroll4Way(A, B)
			}
		})
		b.Run(fmt.Sprintf("%d-bit 6-way unrolled slice HD", v*64), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HammingDistSlice64Unroll6Way(A, B)
			}
		})
	}

}
