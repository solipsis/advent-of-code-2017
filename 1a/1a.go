package main

import (
	"fmt"
	"os"
)

func main() {
	captcha := os.Args[1]
	c := 0
	// start with last char as previous to handle wrap around
	prev := rune(captcha[len(captcha)-1])
	for _, v := range captcha {
		if v == prev {
			c += int(v - '0')
		}
		prev = v
	}
	fmt.Println(c)
}
