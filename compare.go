package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unsafe"
)

/*
 * Compare the 2 algorithms with an input from the command line.
 */

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	bitCount := bits2(n)
	fmt.Printf("%d has %d bits on\n", n, bitCount)
	fmt.Printf("%d\t%064b\t%d\n", n, n, bits2(n))
	x := nextBiggest1(n)
	fmt.Printf("%d\t%064b\t%d\n", x, x, bits2(x))
	y := nextBiggest2(n)
	fmt.Printf("%d\t%064b\t%d\n", y, y, bits2(y))
}

func nextBiggest2(n int) int {

	bytes := int(unsafe.Sizeof(n))
	maxbits := bytes * 8

	var x int

	for i := 0; i < maxbits; i++ {
		bit := (n >> i) & 0x01
		if bit == 1 {
			nextbit := (n >> (i + 1)) & 0x01
			if nextbit == 0 {
				// zero bit i
				x = n ^ (1 << i)
				// set bit i+1
				x |= (1 << (i + 1))
				if (x & 0x01) == 0 {
					for j := i - 1; j > 0; j-- {
						if ((n >> j) & 0x01) == 1 {
							// zero bit j
							x ^= (1 << j)
							// set bit 0
							x |= 1
							break
						}
					}
				}
				break
			}
		}
	}
	return x
}

func nextBiggest1(n int) int {
	bitCount := bits2(n)
	var i int
	for i = n + 1; true; i++ {
		m := bits2(i)
		if m == bitCount {
			break
		}
	}
	return i
}
