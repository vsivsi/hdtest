package hdtest

import (
	"encoding/binary"
	"math/bits"
)

// HammingDistSlice8 counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice8(b1, b2 []uint8) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}
	for x := range b1 {
		sum += bits.OnesCount8(b1[x] ^ b2[x])
	}
	return sum
}

// HammingDistSlice8Unroll2WayInc counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice8Unroll2WayInc(b1, b2 []uint8) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}
	const factor = 2
	startup := len(b1) % factor

	c := 0
	switch startup {
	case 1:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
		}
	}

	for x := startup; x < len(b1); x += factor {
		c = 0
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
	}

	return sum
}

// HammingDistSlice8Unroll3WayInc counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice8Unroll3WayInc(b1, b2 []uint8) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}
	const factor = 3
	startup := len(b1) % factor

	c := 0
	switch startup {
	case 1:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 2:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	}

	for x := startup; x < len(b1); x += factor {
		c = 0
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
	}

	return sum
}

// HammingDistSlice8Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice8Unroll4Way(b1, b2 []uint8) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}

	const factor = 4
	startup := len(b1) % factor

	switch startup {
	case 1:
		{
			sum += bits.OnesCount8(b1[0] ^ b2[0])
		}
	case 2:
		{
			sum += bits.OnesCount8(b1[0] ^ b2[0])
			sum += bits.OnesCount8(b1[1] ^ b2[1])
		}
	case 3:
		{
			sum += bits.OnesCount8(b1[0] ^ b2[0])
			sum += bits.OnesCount8(b1[1] ^ b2[1])
			sum += bits.OnesCount8(b1[2] ^ b2[2])
		}
	}

	for x := startup; x < len(b1); x += factor {
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		sum += bits.OnesCount8(b1[x+1] ^ b2[x+1])
		sum += bits.OnesCount8(b1[x+2] ^ b2[x+2])
		sum += bits.OnesCount8(b1[x+3] ^ b2[x+3])
	}

	return sum
}

// HammingDistSlice8Unroll4WayInc counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice8Unroll4WayInc(b1, b2 []uint8) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}
	const factor = 4
	startup := len(b1) % factor

	c := 0
	switch startup {
	case 1:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 2:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 3:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	}

	for x := startup; x < len(b1); x += factor {
		c = 0
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
	}

	return sum
}

// HammingDistSlice8Unroll8Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice8Unroll8Way(b1, b2 []uint8) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}
	const factor = 8
	startup := len(b1) % factor

	switch startup {
	case 1:
		{
			sum += bits.OnesCount8(b1[0] ^ b2[0])
		}
	case 2:
		{
			sum += bits.OnesCount8(b1[0] ^ b2[0])
			sum += bits.OnesCount8(b1[1] ^ b2[1])
		}
	case 3:
		{
			sum += bits.OnesCount8(b1[0] ^ b2[0])
			sum += bits.OnesCount8(b1[1] ^ b2[1])
			sum += bits.OnesCount8(b1[2] ^ b2[2])
		}
	case 4:
		{
			sum += bits.OnesCount8(b1[0] ^ b2[0])
			sum += bits.OnesCount8(b1[1] ^ b2[1])
			sum += bits.OnesCount8(b1[2] ^ b2[2])
			sum += bits.OnesCount8(b1[3] ^ b2[3])
		}
	case 5:
		{
			sum += bits.OnesCount8(b1[0] ^ b2[0])
			sum += bits.OnesCount8(b1[1] ^ b2[1])
			sum += bits.OnesCount8(b1[2] ^ b2[2])
			sum += bits.OnesCount8(b1[3] ^ b2[3])
			sum += bits.OnesCount8(b1[4] ^ b2[4])
		}
	case 6:
		{
			sum += bits.OnesCount8(b1[0] ^ b2[0])
			sum += bits.OnesCount8(b1[1] ^ b2[1])
			sum += bits.OnesCount8(b1[2] ^ b2[2])
			sum += bits.OnesCount8(b1[3] ^ b2[3])
			sum += bits.OnesCount8(b1[4] ^ b2[4])
			sum += bits.OnesCount8(b1[5] ^ b2[5])
		}
	case 7:
		{
			sum += bits.OnesCount8(b1[0] ^ b2[0])
			sum += bits.OnesCount8(b1[1] ^ b2[1])
			sum += bits.OnesCount8(b1[2] ^ b2[2])
			sum += bits.OnesCount8(b1[3] ^ b2[3])
			sum += bits.OnesCount8(b1[4] ^ b2[4])
			sum += bits.OnesCount8(b1[5] ^ b2[5])
			sum += bits.OnesCount8(b1[6] ^ b2[6])
		}
	}

	for x := startup; x < len(b1); x += factor {
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		sum += bits.OnesCount8(b1[x+1] ^ b2[x+1])
		sum += bits.OnesCount8(b1[x+2] ^ b2[x+2])
		sum += bits.OnesCount8(b1[x+3] ^ b2[x+3])
		sum += bits.OnesCount8(b1[x+4] ^ b2[x+4])
		sum += bits.OnesCount8(b1[x+5] ^ b2[x+5])
		sum += bits.OnesCount8(b1[x+6] ^ b2[x+6])
		sum += bits.OnesCount8(b1[x+7] ^ b2[x+7])
	}

	return sum
}

// HammingDistSlice8Unroll8WayInc counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice8Unroll8WayInc(b1, b2 []uint8) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}
	const factor = 8
	startup := len(b1) % factor

	c := 0
	switch startup {
	case 1:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 2:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 3:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 4:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 5:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 6:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 7:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	}

	for x := startup; x < len(b1); x += factor {
		c = 0
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
	}

	return sum
}

// HammingDistSlice8Unroll16WayInc counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice8Unroll16WayInc(b1, b2 []uint8) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}
	const factor = 16
	startup := len(b1) % factor

	c := 0
	switch startup {
	case 1:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 2:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 3:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 4:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 5:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 6:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 7:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 8:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 9:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 10:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 11:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 12:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 13:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 14:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	case 15:
		{
			sum = bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount8(b1[c] ^ b2[c])
		}
	}

	for x := startup; x < len(b1); x += factor {
		c = 0
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
	}

	return sum
}

// HammingDist32Array8 counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Array8(b1, b2 *[32]uint8) (sum int) {
	for x := range b1 {
		sum += bits.OnesCount8(b1[x] ^ b2[x])
	}
	return sum
}

// HammingDist32Array8Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Array8Unroll4Way(b1, b2 *[32]uint8) (sum int) {

	const factor = 4
	for x := 0; x < len(b1); x += factor {
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		sum += bits.OnesCount8(b1[x+1] ^ b2[x+1])
		sum += bits.OnesCount8(b1[x+2] ^ b2[x+2])
		sum += bits.OnesCount8(b1[x+3] ^ b2[x+3])
	}

	return sum
}

// HammingDist32Array8Unroll4WayInc counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Array8Unroll4WayInc(b1, b2 *[32]uint8) (sum int) {

	const factor = 4
	c := 0
	for x := 0; x < len(b1); x += factor {
		c = 0
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
	}

	return sum
}

// HammingDist32Array8Unroll8Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Array8Unroll8Way(b1, b2 *[32]uint8) (sum int) {

	const factor = 8
	for x := 0; x < len(b1); x += factor {
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		sum += bits.OnesCount8(b1[x+1] ^ b2[x+1])
		sum += bits.OnesCount8(b1[x+2] ^ b2[x+2])
		sum += bits.OnesCount8(b1[x+3] ^ b2[x+3])
		sum += bits.OnesCount8(b1[x+4] ^ b2[x+4])
		sum += bits.OnesCount8(b1[x+5] ^ b2[x+5])
		sum += bits.OnesCount8(b1[x+6] ^ b2[x+6])
		sum += bits.OnesCount8(b1[x+7] ^ b2[x+7])
	}

	return sum
}

// HammingDist32Array8Unroll8WayInc counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Array8Unroll8WayInc(b1, b2 *[32]uint8) (sum int) {

	const factor = 8
	c := 0
	for x := 0; x < len(b1); x += factor {
		c = 0
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
	}

	return sum
}

// HammingDist32Array8Unroll16Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Array8Unroll16Way(b1, b2 *[32]uint8) (sum int) {

	const factor = 16
	for x := 0; x < len(b1); x += factor {
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		sum += bits.OnesCount8(b1[x+1] ^ b2[x+1])
		sum += bits.OnesCount8(b1[x+2] ^ b2[x+2])
		sum += bits.OnesCount8(b1[x+3] ^ b2[x+3])
		sum += bits.OnesCount8(b1[x+4] ^ b2[x+4])
		sum += bits.OnesCount8(b1[x+5] ^ b2[x+5])
		sum += bits.OnesCount8(b1[x+6] ^ b2[x+6])
		sum += bits.OnesCount8(b1[x+7] ^ b2[x+7])
		sum += bits.OnesCount8(b1[x+8] ^ b2[x+8])
		sum += bits.OnesCount8(b1[x+9] ^ b2[x+9])
		sum += bits.OnesCount8(b1[x+10] ^ b2[x+10])
		sum += bits.OnesCount8(b1[x+11] ^ b2[x+11])
		sum += bits.OnesCount8(b1[x+12] ^ b2[x+12])
		sum += bits.OnesCount8(b1[x+13] ^ b2[x+13])
		sum += bits.OnesCount8(b1[x+14] ^ b2[x+14])
		sum += bits.OnesCount8(b1[x+15] ^ b2[x+15])
	}

	return sum
}

// HammingDist32Array8Unroll16WayIncSame counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Array8Unroll16WayIncSame(b1, b2 *[32]uint8) (sum int) {

	for x := 0; x < len(b1); x++ {
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount8(b1[x] ^ b2[x])
	}

	return sum
}

// HammingDist32Array8Unroll16WayInc counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Array8Unroll16WayInc(b1, b2 *[32]uint8) (sum int) {
	const factor = 16
	c := 0
	for x := 0; x < len(b1); x += factor {
		c = 0
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
		c++
		sum += bits.OnesCount8(b1[c] ^ b2[c])
	}

	return sum
}

// HammingDist32Array8UnrollConst counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
func HammingDist32Array8UnrollConst(b1, b2 *[32]uint8) (sum int) {

	sum = bits.OnesCount8(b1[0] ^ b2[0])
	sum += bits.OnesCount8(b1[1] ^ b2[1])
	sum += bits.OnesCount8(b1[2] ^ b2[2])
	sum += bits.OnesCount8(b1[3] ^ b2[3])
	sum += bits.OnesCount8(b1[4] ^ b2[4])
	sum += bits.OnesCount8(b1[5] ^ b2[5])
	sum += bits.OnesCount8(b1[6] ^ b2[6])
	sum += bits.OnesCount8(b1[7] ^ b2[7])
	sum += bits.OnesCount8(b1[8] ^ b2[8])
	sum += bits.OnesCount8(b1[9] ^ b2[9])
	sum += bits.OnesCount8(b1[10] ^ b2[10])
	sum += bits.OnesCount8(b1[11] ^ b2[11])
	sum += bits.OnesCount8(b1[12] ^ b2[12])
	sum += bits.OnesCount8(b1[13] ^ b2[13])
	sum += bits.OnesCount8(b1[14] ^ b2[14])
	sum += bits.OnesCount8(b1[15] ^ b2[15])
	sum += bits.OnesCount8(b1[16] ^ b2[16])
	sum += bits.OnesCount8(b1[17] ^ b2[17])
	sum += bits.OnesCount8(b1[18] ^ b2[18])
	sum += bits.OnesCount8(b1[19] ^ b2[19])
	sum += bits.OnesCount8(b1[20] ^ b2[20])
	sum += bits.OnesCount8(b1[21] ^ b2[21])
	sum += bits.OnesCount8(b1[22] ^ b2[22])
	sum += bits.OnesCount8(b1[23] ^ b2[23])
	sum += bits.OnesCount8(b1[24] ^ b2[24])
	sum += bits.OnesCount8(b1[25] ^ b2[25])
	sum += bits.OnesCount8(b1[26] ^ b2[26])
	sum += bits.OnesCount8(b1[27] ^ b2[27])
	sum += bits.OnesCount8(b1[28] ^ b2[28])
	sum += bits.OnesCount8(b1[29] ^ b2[29])
	sum += bits.OnesCount8(b1[30] ^ b2[30])
	sum += bits.OnesCount8(b1[31] ^ b2[31])

	return sum
}

// HammingDist32Array8UnrollConstBC counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
func HammingDist32Array8UnrollConstBC(b1, b2 *[32]uint8) (sum int) {
	_ = b1[31]
	_ = b2[31]
	sum = bits.OnesCount8(b1[0] ^ b2[0])
	sum += bits.OnesCount8(b1[1] ^ b2[1])
	sum += bits.OnesCount8(b1[2] ^ b2[2])
	sum += bits.OnesCount8(b1[3] ^ b2[3])
	sum += bits.OnesCount8(b1[4] ^ b2[4])
	sum += bits.OnesCount8(b1[5] ^ b2[5])
	sum += bits.OnesCount8(b1[6] ^ b2[6])
	sum += bits.OnesCount8(b1[7] ^ b2[7])
	sum += bits.OnesCount8(b1[8] ^ b2[8])
	sum += bits.OnesCount8(b1[9] ^ b2[9])
	sum += bits.OnesCount8(b1[10] ^ b2[10])
	sum += bits.OnesCount8(b1[11] ^ b2[11])
	sum += bits.OnesCount8(b1[12] ^ b2[12])
	sum += bits.OnesCount8(b1[13] ^ b2[13])
	sum += bits.OnesCount8(b1[14] ^ b2[14])
	sum += bits.OnesCount8(b1[15] ^ b2[15])
	sum += bits.OnesCount8(b1[16] ^ b2[16])
	sum += bits.OnesCount8(b1[17] ^ b2[17])
	sum += bits.OnesCount8(b1[18] ^ b2[18])
	sum += bits.OnesCount8(b1[19] ^ b2[19])
	sum += bits.OnesCount8(b1[20] ^ b2[20])
	sum += bits.OnesCount8(b1[21] ^ b2[21])
	sum += bits.OnesCount8(b1[22] ^ b2[22])
	sum += bits.OnesCount8(b1[23] ^ b2[23])
	sum += bits.OnesCount8(b1[24] ^ b2[24])
	sum += bits.OnesCount8(b1[25] ^ b2[25])
	sum += bits.OnesCount8(b1[26] ^ b2[26])
	sum += bits.OnesCount8(b1[27] ^ b2[27])
	sum += bits.OnesCount8(b1[28] ^ b2[28])
	sum += bits.OnesCount8(b1[29] ^ b2[29])
	sum += bits.OnesCount8(b1[30] ^ b2[30])
	sum += bits.OnesCount8(b1[31] ^ b2[31])

	return sum
}

// HammingDist32Array8UnrollInc counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Array8UnrollInc(b1, b2 *[32]uint8) (sum int) {

	c := 0
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	return sum
}

// HammingDist32Slice8UnrollConst counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Slice8UnrollConst(b1, b2 []uint8) (sum int) {
	const factor = 32
	if len(b1) != factor || len(b2) != factor {
		panic("slices of incorrect length passed to HammingDist")
	}

	sum = bits.OnesCount8(b1[0] ^ b2[0])
	sum += bits.OnesCount8(b1[1] ^ b2[1])
	sum += bits.OnesCount8(b1[2] ^ b2[2])
	sum += bits.OnesCount8(b1[3] ^ b2[3])
	sum += bits.OnesCount8(b1[4] ^ b2[4])
	sum += bits.OnesCount8(b1[5] ^ b2[5])
	sum += bits.OnesCount8(b1[6] ^ b2[6])
	sum += bits.OnesCount8(b1[7] ^ b2[7])
	sum += bits.OnesCount8(b1[8] ^ b2[8])
	sum += bits.OnesCount8(b1[9] ^ b2[9])
	sum += bits.OnesCount8(b1[10] ^ b2[10])
	sum += bits.OnesCount8(b1[11] ^ b2[11])
	sum += bits.OnesCount8(b1[12] ^ b2[12])
	sum += bits.OnesCount8(b1[13] ^ b2[13])
	sum += bits.OnesCount8(b1[14] ^ b2[14])
	sum += bits.OnesCount8(b1[15] ^ b2[15])
	sum += bits.OnesCount8(b1[16] ^ b2[16])
	sum += bits.OnesCount8(b1[17] ^ b2[17])
	sum += bits.OnesCount8(b1[18] ^ b2[18])
	sum += bits.OnesCount8(b1[19] ^ b2[19])
	sum += bits.OnesCount8(b1[20] ^ b2[20])
	sum += bits.OnesCount8(b1[21] ^ b2[21])
	sum += bits.OnesCount8(b1[22] ^ b2[22])
	sum += bits.OnesCount8(b1[23] ^ b2[23])
	sum += bits.OnesCount8(b1[24] ^ b2[24])
	sum += bits.OnesCount8(b1[25] ^ b2[25])
	sum += bits.OnesCount8(b1[26] ^ b2[26])
	sum += bits.OnesCount8(b1[27] ^ b2[27])
	sum += bits.OnesCount8(b1[28] ^ b2[28])
	sum += bits.OnesCount8(b1[29] ^ b2[29])
	sum += bits.OnesCount8(b1[30] ^ b2[30])
	sum += bits.OnesCount8(b1[31] ^ b2[31])

	return sum
}

// HammingDist32Slice8UnrollInc counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist32Slice8UnrollInc(b1, b2 []uint8) (sum int) {
	const factor = 32
	if len(b1) != factor || len(b2) != factor {
		panic("slices of incorrect length passed to HammingDist")
	}
	c := 0
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	c++
	sum = bits.OnesCount8(b1[c] ^ b2[c])
	return sum
}

// HammingDist32Slice8Convert64 counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
func HammingDist32Slice8Convert64(b1, b2 []uint8) (sum int) {
	const factor = 32
	if len(b1) != factor || len(b2) != factor {
		panic("slices of incorrect length passed to HammingDist")
	}
	var sA, sB [4]uint64
	c := 0
	s := 0
	sA[c] = binary.LittleEndian.Uint64(b1[s:])
	sB[c] = binary.LittleEndian.Uint64(b2[s:])
	c++
	s += 8
	sA[c] = binary.LittleEndian.Uint64(b1[s:])
	sB[c] = binary.LittleEndian.Uint64(b2[s:])
	c++
	s += 8
	sA[c] = binary.LittleEndian.Uint64(b1[s:])
	sB[c] = binary.LittleEndian.Uint64(b2[s:])
	c++
	s += 8
	sA[c] = binary.LittleEndian.Uint64(b1[s:])
	sB[c] = binary.LittleEndian.Uint64(b2[s:])

	HammingDist4Array64Unroll(&sA, &sB)
	return sum
}

// HammingDist32Array8Convert64 counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint8 values.
func HammingDist32Array8Convert64(b1, b2 *[32]uint8) (sum int) {
	var sA, sB [4]uint64
	c := 0
	s := 0
	sA[c] = binary.LittleEndian.Uint64(b1[s:])
	sB[c] = binary.LittleEndian.Uint64(b2[s:])
	c++
	s += 8
	sA[c] = binary.LittleEndian.Uint64(b1[s:])
	sB[c] = binary.LittleEndian.Uint64(b2[s:])
	c++
	s += 8
	sA[c] = binary.LittleEndian.Uint64(b1[s:])
	sB[c] = binary.LittleEndian.Uint64(b2[s:])
	c++
	s += 8
	sA[c] = binary.LittleEndian.Uint64(b1[s:])
	sB[c] = binary.LittleEndian.Uint64(b2[s:])

	HammingDist4Array64Unroll(&sA, &sB)
	return sum
}
