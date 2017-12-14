package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanBytes)

	// bonus bytes to append given in the problem statement
	end := []int{17, 31, 73, 47, 23}
	input := make([]int, 0)
	for sc.Scan() {
		input = append(input, int(sc.Bytes()[0]))
	}
	input = append(input[:len(input)-1], end...) // append the bonus bytes skipping the newline character from the input

	// init list with values 0 - 255
	list := make([]int, 256)
	for i := 0; i < 256; i++ {
		list[i] = i
	}

	pos := 0
	skip := 0
	// 64 rounds of hashing
	for r := 0; r < 64; r++ {
		for _, v := range input {
			// get a sublist and reverse it
			sub := sublist(list, pos, v)
			reverse(sub)

			// update the list using our sublist
			for i := 0; i < len(sub); i++ {
				list[(i+pos)%len(list)] = sub[i]
			}

			// update the start postion and increment the skip length
			pos = (pos + v + skip) % len(list)
			skip++
		}
	}

	// create 16 parts of the dense hash by XORing the list segments
	for i := 0; i < 16; i++ {
		dense := 0
		for j := 0; j < 16; j++ {
			dense ^= list[(i*16)+j]
		}
		fmt.Printf("%x", dense)
	}
	fmt.Println()
}

// Sublist creates a sublist of a given length wrapping around to the beginning if necessary
func sublist(arr []int, start, length int) []int {
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
