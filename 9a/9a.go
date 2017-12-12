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

	depth := 0
	sum := 0
	garbage := false
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		switch r {
		case '!':
			reader.ReadRune() // skip a character
		case '<':
			garbage = true
		case '>':
			garbage = false
		case '{':
			if !garbage {
				depth++
			}
		case '}':
			if !garbage {
				sum += depth
				depth--
			}
		default:
		}
	}
	fmt.Println(sum)

}
