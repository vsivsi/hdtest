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

// HammingDist4Array64 counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as len 4 arrays of uint64 values.
func HammingDist4Array64(b1, b2 *[4]uint64) (sum int) {
	for x := range b1 {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
	}
	return sum
}
