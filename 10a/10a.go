package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	// convert input to integer slice
	sc.Scan()
	strInput := strings.Split(sc.Text(), ",")
	input := make([]int, len(strInput))
	for i := 0; i < len(input); i++ {
		input[i], _ = strconv.Atoi(strInput[i])
	}

	// init list
	list := make([]int, 256)
	for i := 0; i < 256; i++ {
		list[i] = i
	}

	pos := 0
	skip := 0
	for _, v := range input {
		// create a sublist and reverse it
		sub := sublist(list, pos, v)
		reverse(sub)

		// use reversed sublist to update the interval
		for i := 0; i < len(sub); i++ {
			list[(i+pos)%len(list)] = sub[i]
		}

		// update the start position and increment the skip length
		pos = (pos + v + skip) % len(list)
		skip++
	}

	fmt.Println(list[0] * list[1])

}

// Sublist creates a sublist of a given length that wraps around to the beginning if necessary
func sublist(arr []int, start, length int) []int {
	fmt.Println(arr, start, length)
	sub := make([]int, length)
	for i := 0; i < length; i++ {
		sub[i] = arr[start]
		start = (start + 1) % len(arr)
	}
	return sub
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
