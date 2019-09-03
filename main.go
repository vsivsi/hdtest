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

// HammingDistSlice64Unroll6Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice64Unroll6Way(b1, b2 []uint64) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}

	startup := len(b1) % 6

	switch startup {
	case 1:
		{
			sum += bits.OnesCount64(b1[0] ^ b2[0])
		}
	case 2:
		{
			sum += bits.OnesCount64(b1[0] ^ b2[0])
			sum += bits.OnesCount64(b1[1] ^ b2[1])
		}
	case 3:
		{
			sum += bits.OnesCount64(b1[0] ^ b2[0])
			sum += bits.OnesCount64(b1[1] ^ b2[1])
			sum += bits.OnesCount64(b1[2] ^ b2[2])
		}
	case 4:
		{
			sum += bits.OnesCount64(b1[0] ^ b2[0])
			sum += bits.OnesCount64(b1[1] ^ b2[1])
			sum += bits.OnesCount64(b1[2] ^ b2[2])
			sum += bits.OnesCount64(b1[3] ^ b2[3])
		}
	case 5:
		{
			sum += bits.OnesCount64(b1[0] ^ b2[0])
			sum += bits.OnesCount64(b1[1] ^ b2[1])
			sum += bits.OnesCount64(b1[2] ^ b2[2])
			sum += bits.OnesCount64(b1[3] ^ b2[3])
			sum += bits.OnesCount64(b1[4] ^ b2[4])
		}
	}

	for x := startup; x < len(b1); x++ {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
	}

	return sum
}

// HammingDistSlice64Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice64Unroll4Way(b1, b2 []uint64) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}

	startup := len(b1) % 4

	switch startup {
	case 1:
		{
			sum += bits.OnesCount64(b1[0] ^ b2[0])
		}
	case 2:
		{
			sum += bits.OnesCount64(b1[0] ^ b2[0])
			sum += bits.OnesCount64(b1[1] ^ b2[1])
		}
	case 3:
		{
			sum += bits.OnesCount64(b1[0] ^ b2[0])
			sum += bits.OnesCount64(b1[1] ^ b2[1])
			sum += bits.OnesCount64(b1[2] ^ b2[2])
		}
	}

	for x := startup; x < len(b1); x++ {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
	}

	return sum
}

// HammingDistSlice64Unroll3Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDistSlice64Unroll3Way(b1, b2 []uint64) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}

	startup := len(b1) % 3

	switch startup {
	case 1:
		{
			sum += bits.OnesCount64(b1[0] ^ b2[0])
		}
	case 2:
		{
			sum += bits.OnesCount64(b1[0] ^ b2[0])
			sum += bits.OnesCount64(b1[1] ^ b2[1])
		}
	}
	for x := startup; x < len(b1); x++ {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
	}

	return sum
}

// HammingDist4Slice64Unroll counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as len 4 slices of uint64 values.
func HammingDist4Slice64Unroll(b1, b2 []uint64) (sum int) {
	if len(b1) != 4 || len(b1) != len(b2) {
		panic("invalid slices passed to HammingDist")
	}

	c := 0
	sum += bits.OnesCount64(b1[c] ^ b2[c])
	c++
	sum += bits.OnesCount64(b1[c] ^ b2[c])
	c++
	sum += bits.OnesCount64(b1[c] ^ b2[c])
	c++
	sum += bits.OnesCount64(b1[c] ^ b2[c])
	return sum
}

// HammingDist4Array64Unroll counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as len 4 arrays of uint64 values.
func HammingDist4Array64Unroll(b1, b2 *[4]uint64) (sum int) {
	c := 0
	sum = bits.OnesCount64(b1[c] ^ b2[c])
	c++
	sum += bits.OnesCount64(b1[c] ^ b2[c])
	c++
	sum += bits.OnesCount64(b1[c] ^ b2[c])
	c++
	sum += bits.OnesCount64(b1[c] ^ b2[c])
	return sum
}

// HammingDist1Array64Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist1Array64Unroll4Way(b1, b2 *[1]uint64) (sum int) {
	sum = bits.OnesCount64(b1[0] ^ b2[0])
	return sum
}

// HammingDist2Array64Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist2Array64Unroll4Way(b1, b2 *[2]uint64) (sum int) {
	sum = bits.OnesCount64(b1[0] ^ b2[0])
	sum += bits.OnesCount64(b1[1] ^ b2[1])
	return sum
}

// HammingDist3Array64Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist3Array64Unroll4Way(b1, b2 *[3]uint64) (sum int) {
	sum = bits.OnesCount64(b1[0] ^ b2[0])
	sum += bits.OnesCount64(b1[1] ^ b2[1])
	sum += bits.OnesCount64(b1[2] ^ b2[2])
	return sum
}

// HammingDist4Array64Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist4Array64Unroll4Way(b1, b2 *[4]uint64) (sum int) {
	const startup = len(b1) % 4
	switch startup {
	case 1:
		{
			sum = bits.OnesCount64(b1[0] ^ b2[0])
		}
	case 2:
		{
			sum = bits.OnesCount64(b1[0] ^ b2[0])
			sum += bits.OnesCount64(b1[1] ^ b2[1])
		}
	case 3:
		{
			sum = bits.OnesCount64(b1[0] ^ b2[0])
			sum += bits.OnesCount64(b1[1] ^ b2[1])
			sum += bits.OnesCount64(b1[2] ^ b2[2])
		}
	}

	for x := startup; x < len(b1); x += 4 {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		sum += bits.OnesCount64(b1[x+1] ^ b2[x+1])
		sum += bits.OnesCount64(b1[x+2] ^ b2[x+2])
		sum += bits.OnesCount64(b1[x+3] ^ b2[x+3])
	}

	return sum
}

// HammingDist5Array64Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist5Array64Unroll4Way(b1, b2 *[5]uint64) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}

	const startup = len(b1) % 4
	c := 0
	switch startup {
	case 1:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 2:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 3:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	}

	for x := startup; x < len(b1); x++ {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
	}

	return sum
}

// HammingDist6Array64Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist6Array64Unroll4Way(b1, b2 *[6]uint64) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}

	const startup = len(b1) % 4
	c := 0
	switch startup {
	case 1:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 2:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 3:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	}

	for x := startup; x < len(b1); x++ {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
	}

	return sum
}

// HammingDist7Array64Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist7Array64Unroll4Way(b1, b2 *[7]uint64) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}

	const startup = len(b1) % 4
	c := 0
	switch startup {
	case 1:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 2:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 3:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	}

	for x := startup; x < len(b1); x++ {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
	}

	return sum
}

// HammingDist8Array64Unroll4Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist8Array64Unroll4Way(b1, b2 *[8]uint64) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}

	const startup = len(b1) % 4
	c := 0
	switch startup {
	case 1:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 2:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 3:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	}

	for x := startup; x < len(b1); x++ {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
	}

	return sum
}

// HammingDist8Array64Unroll6Way counts the number bitwise differences between
// n-bit signatures b1 and b2 passed as slices of uint64 values.
// Function will panic if b1 and b2 aren't of equal length.
func HammingDist8Array64Unroll6Way(b1, b2 *[8]uint64) (sum int) {
	if len(b1) != len(b2) {
		panic("slices of different lengths passed to HammingDist")
	}

	const startup = len(b1) % 6
	c := 0
	switch startup {
	case 1:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 2:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 3:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 4:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	case 5:
		{
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
			c++
			sum += bits.OnesCount64(b1[c] ^ b2[c])
		}
	}

	for x := startup; x < len(b1); x++ {
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
		x++
		sum += bits.OnesCount64(b1[x] ^ b2[x])
	}

	return sum
}
