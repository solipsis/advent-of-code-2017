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

	sc.Scan()
	base := strings.TrimSpace(sc.Text())

	// turn knot hashes into a grid
	arr := make([][]string, 0)
	for i := 0; i < 128; i++ {
		arr = append(arr, strings.Split(knotHash(base+"-"+strconv.Itoa(i)), ""))
	}

	// the number of fill calls we make is the answer
	blocks := 0
	for r := 0; r < len(arr); r++ {
		for c := 0; c < len(arr[0]); c++ {
			// fill if the cell is a "1"
			if arr[r][c] == "1" {
				blocks++
				fill(r, c, arr)
			}
		}
	}
	fmt.Println(blocks)
}

// keep filling "1"s until the block is complete
func fill(r, c int, arr [][]string) {
	// bounds check
	if r < 0 || c < 0 || r >= len(arr) || c >= len(arr[0]) {
		return
	}
	if arr[r][c] == "1" {
		arr[r][c] = "2"
		fill(r-1, c, arr)
		fill(r+1, c, arr)
		fill(r, c-1, arr)
		fill(r, c+1, arr)
	}

}

//*******************************************************
// Knot hash from challenge 10
//*******************************************************
func knotHash(str string) string {

	input := []byte(str)

	// bonus bytes to append given in the problem statement
	end := []byte{17, 31, 73, 47, 23}
	input = append(input, end...) // append the bonus bytes

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
			sub := sublist(list, pos, int(v))
			reverse(sub)

			// update the list using our sublist
			for i := 0; i < len(sub); i++ {
				list[(i+pos)%len(list)] = sub[i]
			}

			// update the start postion and increment the skip length
			pos = (pos + int(v) + skip) % len(list)
			skip++
		}
	}

	// create 16 parts of the dense hash by XORing the list segments
	s := ""
	for i := 0; i < 16; i++ {
		dense := 0
		for j := 0; j < 16; j++ {
			dense ^= list[(i*16)+j]
		}
		s += fmt.Sprintf("%08b", dense)
	}
	fmt.Println(s)
	return s
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
