package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	numValid := 0
	for sc.Scan() {
		// add each word to map
		words := strings.Split(sc.Text(), " ")
		m := make(map[string]bool)
		valid := true
		// if there is a dupe the passphrase is not valid
		for _, word := range words {
			if m[word] {
				valid = false
				break
			}
			m[word] = true
		}
		if valid {
			numValid++
		}
	}
	fmt.Println(numValid)
}
