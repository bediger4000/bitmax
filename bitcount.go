package main

import "unsafe"

// bits counts 1-bits in the formal argument
// the hard way, by shifting and bitwise-and.
func bits(n int) int {

	count := 0

	bytes := int(unsafe.Sizeof(n))

	for i := 0; i < bytes*8; i++ {
		if (n & 0x01) != 0 {
			count++
		}
		n >>= 1
	}
	return count
}
