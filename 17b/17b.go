package main

import "fmt"

const steps = 349 // puzzle input
func main() {
	pos := 0
	result := 0
	// track each time the position lands on position 0
	for i := 1; i <= 50000000; i++ {
		pos = (pos+steps)%i + 1
		if pos == 1 {
			result = i
		}
	}
	fmt.Println(result)
}
