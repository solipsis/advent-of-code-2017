package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read and parse input
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)

	arr := make([]int, 0)
	for sc.Scan() {
		i, _ := strconv.Atoi(sc.Text())
		arr = append(arr, i)
	}

	// map to track the first cycle that we see each permutation
	m := make(map[string]int)
	m[arrToKey(arr)] = 0

	// run through distribution cycles until we get a duplicate
	cycles := 0
	for {
		max, index := max(arr)
		arr[index] = 0
		for max > 0 {
			index = (index + 1) % len(arr)
			arr[index]++
			max--
		}
		cycles++

		// check for a permutation we have seen before and calculate how
		// many cycles since we last saw it
		key := arrToKey(arr)
		if m[key] > 0 {
			fmt.Println(cycles - m[key])
			break
		}
		m[key] = cycles
	}
}

// find the max and the index it first occurs
func max(arr []int) (max, index int) {

	for i, v := range arr {
		if v > max {
			max = v
			index = i
		}
	}
	return max, index
}

// convert array to suitable map key
func arrToKey(arr []int) string {
	s := ""
	for _, v := range arr {
		s += strconv.Itoa(v)
	}
	return s
}
