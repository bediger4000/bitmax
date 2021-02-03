package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
 * Given an integer n, find the next biggest integer with the same number of
 * 1-bits on.  For example, given the number 6 (0110 in binary), return 9 (1001).
 */

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	bitCount := bits(n)
	fmt.Printf("%d has %d bits on\n", n, bitCount)

	for i := n + 1; true; i++ {
		m := bits(i)
		if m == bitCount {
			fmt.Printf("%d\n", i)
			break
		}
	}
}
