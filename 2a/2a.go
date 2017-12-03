package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt") // ignoring error
	sc := bufio.NewScanner(file)

	// total checksum
	sum := 0

	// read each line
	for sc.Scan() {
		words := strings.Split(sc.Text(), "\t")

		// find min and max for each line
		min, max := 9999999999, 0
		for _, w := range words {
			i, _ := strconv.Atoi(w)
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}
		sum += max - min
	}
	fmt.Println(sum)

}
