package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// the bottom tower has a list but appears in no lists
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	candidates := make(map[string]bool)
	children := make(map[string]bool)
	for sc.Scan() {

		// a node is a candidate if it has children
		arr := strings.Split(sc.Text(), " ")
		if len(arr) > 2 {
			candidates[arr[0]] = true
			// track everything that is a child of another node
			for _, child := range arr[3:] {
				if strings.HasSuffix(child, ",") {
					child = child[:len(child)-1]
				}
				children[child] = true
			}
		}
	}
	// the solution is the candidate that was not a child
	for candidate := range candidates {
		if !children[candidate] {
			fmt.Println(candidate)
		}
	}
}
