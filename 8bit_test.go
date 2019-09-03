package hdtest

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"
)

func generateRandomSliceSig8(n int) (sig []uint8) {
	sig = make([]uint8, n)
	for x := range sig {
		sig[x] = uint8(rand.Uint64())
	}
	return sig
}

func BenchmarkArrayHD8(b *testing.B) {

	const max = 128
	const bits = 8

	sigA := generateRandomSliceSig8(max)
	sigB := generateRandomSliceSig8(max)

	// testLens := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	// 12, 15, 18, 21, 24, 27, 30, 40, 50, 75, 100}

	var s256A, s256B [32]uint8
	copy(s256A[:], sigA)
	copy(s256B[:], sigB)

	b.Run(fmt.Sprintf("IGNORE spin-up"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8(&s256A, &s256B)
		}
	})

	b.Run(fmt.Sprintf("%d-bit %d loop slice HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice8(s256A[:], s256B[:])
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop slice 2-way inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice8Unroll2WayInc(s256A[:], s256B[:])
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop slice 3-way inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice8Unroll3WayInc(s256A[:], s256B[:])
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop slice 4-way HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice8Unroll4Way(s256A[:], s256B[:])
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop slice 4-way inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice8Unroll4WayInc(s256A[:], s256B[:])
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop slice 8-way HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice8Unroll8Way(s256A[:], s256B[:])
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop slice 8-way inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice8Unroll8WayInc(s256A[:], s256B[:])
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop slice 16-way inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDistSlice8Unroll16WayInc(s256A[:], s256B[:])
		}
	})

	b.Run(fmt.Sprintf("%d-bit %d loop array HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array 4-way HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8Unroll4Way(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array 4-way inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8Unroll4WayInc(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array 8-way HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8Unroll8Way(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array 8-way inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8Unroll8WayInc(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array 16-way HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8Unroll16Way(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array 16-way inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8Unroll16WayInc(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array 16-way inc same HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8Unroll16WayIncSame(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array full const HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8UnrollConst(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array full const BC HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8UnrollConstBC(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array full inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8UnrollInc(&s256A, &s256B)
		}
	})

	b.Run(fmt.Sprintf("%d-bit %d loop slice full const HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Slice8UnrollConst(s256A[:], s256B[:])
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop slice full inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Slice8UnrollInc(s256A[:], s256B[:])
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop slice full dec HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Slice8UnrollDec(s256A[:], s256B[:])
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array convert 64-bit HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var sA, sB [4]uint64
			sA[0] = binary.LittleEndian.Uint64(s256A[0:])
			sA[1] = binary.LittleEndian.Uint64(s256A[8:])
			sA[2] = binary.LittleEndian.Uint64(s256A[16:])
			sA[3] = binary.LittleEndian.Uint64(s256A[24:])
			sB[0] = binary.LittleEndian.Uint64(s256B[0:])
			sB[1] = binary.LittleEndian.Uint64(s256B[8:])
			sB[2] = binary.LittleEndian.Uint64(s256B[16:])
			sB[3] = binary.LittleEndian.Uint64(s256B[24:])
			HammingDist4Array64Unroll(&sA, &sB)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array convert 64-bit inc HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var sA, sB [4]uint64
			c := 0
			s := 0
			sA[c] = binary.LittleEndian.Uint64(s256A[s:])
			sB[c] = binary.LittleEndian.Uint64(s256B[s:])
			c++
			s += 8
			sA[c] = binary.LittleEndian.Uint64(s256A[s:])
			sB[c] = binary.LittleEndian.Uint64(s256B[s:])
			c++
			s += 8
			sA[c] = binary.LittleEndian.Uint64(s256A[s:])
			sB[c] = binary.LittleEndian.Uint64(s256B[s:])
			c++
			s += 8
			sA[c] = binary.LittleEndian.Uint64(s256A[s:])
			sB[c] = binary.LittleEndian.Uint64(s256B[s:])

			HammingDist4Array64Unroll(&sA, &sB)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array convert 64-bit inc shift HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var sA, sB [4]uint64
			c := 0
			sA[c] = binary.LittleEndian.Uint64(s256A[c<<3:])
			sB[c] = binary.LittleEndian.Uint64(s256B[c<<3:])
			c++
			sA[c] = binary.LittleEndian.Uint64(s256A[c<<3:])
			sB[c] = binary.LittleEndian.Uint64(s256B[c<<3:])
			c++
			sA[c] = binary.LittleEndian.Uint64(s256A[c<<3:])
			sB[c] = binary.LittleEndian.Uint64(s256B[c<<3:])
			c++
			sA[c] = binary.LittleEndian.Uint64(s256A[c<<3:])
			sB[c] = binary.LittleEndian.Uint64(s256B[c<<3:])

			HammingDist4Array64Unroll(&sA, &sB)
		}
	})

	b.Run(fmt.Sprintf("%d-bit %d loop array convert 64-bit range HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var sA, sB [4]uint64
			for c := range sA {
				sA[c] = binary.LittleEndian.Uint64(s256A[c<<3:])
				sB[c] = binary.LittleEndian.Uint64(s256B[c<<3:])
			}
			HammingDist4Array64Unroll(&sA, &sB)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array convert 64-bit function HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8Convert64(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop array convert 64-bit function const HD", len(s256A)*bits, bits), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingDist32Array8Convert64Const(&s256A, &s256B)
		}
	})
	b.Run(fmt.Sprintf("%d-bit %d loop slice convert 64-bit separate functions HD", len(s256A)*bits, bits), func(b *testing.B) {
		sA := s256A[:]
		sB := s256B[:]
		var aA, aB [4]uint64
		for i := 0; i < b.N; i++ {
			HammingDist4Array64Unroll(ConvertByteSliceToUint64Array4(sA, &aA), ConvertByteSliceToUint64Array4(sB, &aB))
		}
	})
}
