package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)

	sum := 0
	garbage := false
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		// count number of garbage characters
		switch r {
		case '!':
			reader.ReadRune() // skip a character
		case '<':
			if garbage {
				sum++
			}
			garbage = true
		case '>':
			garbage = false
		default:
			if garbage {
				sum++
			}
		}
	}
	fmt.Println(sum)

}
