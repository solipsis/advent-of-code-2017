package main

import "fmt"

const (
	seedA   = 634
	seedB   = 301
	factorA = 16807
	factorB = 48271
	mod     = 1<<31 - 1 // 2147483647
	mask    = 0xFFFF    // bottom 16 bits
)

func main() {

	a := seedA
	b := seedB
	count := 0

	for pairs := 0; pairs < 5000000; pairs++ {

		// wait till we get a valid input from each generator
		for a = (a * factorA) % mod; a%4 != 0; a = (a * factorA) % mod {
		}
		for b = (b * factorB) % mod; b%8 != 0; b = (b * factorB) % mod {
		}

		// compare bottom 16 bits
		if a&mask == b&mask {
			count++
		}
	}
	fmt.Println(count)

}
