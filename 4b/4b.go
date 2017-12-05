package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	numValid := 0
	for sc.Scan() {

		m := make(map[string]bool)
		valid := true

		// convert each string to rune array and sort then add to map
		words := strings.Split(sc.Text(), " ")
		for _, word := range words {
			sorted := []rune(word)
			sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
			word := string(sorted)
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
