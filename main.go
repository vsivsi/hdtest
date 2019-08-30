package hdtest

import "math/bits"

// HammingDistSlice64 counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice64(b1, b2 []uint64) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}
	for x := range b1 {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
	}
	return sum
}

// HammingDist4Array64Unroll counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as len 4 arrays of uint64 values.
func HammingDist4Array64Unroll(b1, b2 *[4]uint64) (sum int) {
	sum = bits.OnesCount64(b1[0] ^ b2[0])
	sum += bits.OnesCount64(b1[1] ^ b2[1])
	sum += bits.OnesCount64(b1[2] ^ b2[2])
	sum += bits.OnesCount64(b1[3] ^ b2[3])
	return sum
}
