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

func BenchmarkArrayHD(b *testing.B) {

	const max = 100

	sigA := generateRandomSliceSig(max)
	sigB := generateRandomSliceSig(max)

	// testLens := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	// 12, 15, 18, 21, 24, 27, 30, 40, 50, 75, 100}

	var s1A, s1B [1]uint64
	copy(s1A[:], sigA)
	copy(s1B[:], sigB)

	b.Run(fmt.Sprintf("%d-bit loop array HD", len(s1A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist1Array64Unroll4Way(&s1A, &s1B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit loop slice HD", len(s1A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64Unroll4Way(s1A[:], s1B[:])
		}
	})

	var s2A, s2B [2]uint64
	copy(s2A[:], sigA)
	copy(s2B[:], sigB)

	b.Run(fmt.Sprintf("%d-bit loop array HD", len(s2A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist2Array64Unroll4Way(&s2A, &s2B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit loop slice HD", len(s2A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64Unroll4Way(s2A[:], s2B[:])
		}
	})

	var s3A, s3B [3]uint64
	copy(s3A[:], sigA)
	copy(s3B[:], sigB)

	b.Run(fmt.Sprintf("%d-bit loop array HD", len(s3A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist3Array64Unroll4Way(&s3A, &s3B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit loop slice HD", len(s3A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64Unroll4Way(s3A[:], s3B[:])
		}
	})

	var s4A, s4B [4]uint64
	copy(s4A[:], sigA)
	copy(s4B[:], sigB)

	b.Run(fmt.Sprintf("%d-bit loop array HD", len(s4A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist4Array64Unroll4Way(&s4A, &s4B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit loop slice HD", len(s4A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64Unroll4Way(s4A[:], s4B[:])
		}
	})

	var s5A, s5B [5]uint64
	copy(s5A[:], sigA)
	copy(s5B[:], sigB)

	b.Run(fmt.Sprintf("%d-bit loop array HD", len(s5A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist5Array64Unroll4Way(&s5A, &s5B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit loop slice HD", len(s5A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64Unroll4Way(s5A[:], s5B[:])
		}
	})

	var s6A, s6B [6]uint64
	copy(s6A[:], sigA)
	copy(s6B[:], sigB)

	b.Run(fmt.Sprintf("%d-bit loop array HD", len(s6A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist6Array64Unroll4Way(&s6A, &s6B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit loop slice HD", len(s6A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64Unroll4Way(s6A[:], s6B[:])
		}
	})

	var s7A, s7B [7]uint64
	copy(s7A[:], sigA)
	copy(s7B[:], sigB)

	b.Run(fmt.Sprintf("%d-bit loop array HD", len(s7A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist7Array64Unroll4Way(&s7A, &s7B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit loop slice HD", len(s7A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64Unroll4Way(s7A[:], s7B[:])
		}
	})

	var s8A, s8B [8]uint64
	copy(s8A[:], sigA)
	copy(s8B[:], sigB)

	b.Run(fmt.Sprintf("%d-bit loop array HD", len(s8A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist8Array64Unroll4Way(&s8A, &s8B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit loop array HD", len(s8A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist8Array64Unroll6Way(&s8A, &s8B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit loop slice HD", len(s8A)*64), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice64Unroll4Way(s8A[:], s8B[:])
		}
	})

}
