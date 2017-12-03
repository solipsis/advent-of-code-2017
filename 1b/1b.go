package main

import (
	"fmt"
	"os"
)

func main() {
	// read input from Stdin
	captcha := os.Args[1]

	mid := len(captcha) / 2
	sum := 0
	for i := range captcha {
		if captcha[i] == captcha[(i+mid)%len(captcha)] {
			sum += int(captcha[i]) - '0' // add value of digit
		}
	}
	fmt.Println(sum)
}
