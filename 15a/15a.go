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
	for i := 0; i < 40000000; i++ {
		a = (a * factorA) % mod
		b = (b * factorB) % mod

		// compare bottom 16 bits
		if a&mask == b&mask {
			count++
		}
	}
	fmt.Println(count)
}
