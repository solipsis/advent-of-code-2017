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

	// map to track seen permutations
	m := make(map[string]bool)
	m[arrToKey(arr)] = true

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

		// check for permutation we have already seen
		key := arrToKey(arr)
		if m[key] {
			break
		}
		m[key] = true
	}

	fmt.Println(cycles)
}

// find the max value and the index it occurs at
func max(arr []int) (max, index int) {

	for i, v := range arr {
		if v > max {
			max = v
			index = i
		}
	}
	return max, index
}

// convert array to suitable key string
func arrToKey(arr []int) string {
	s := ""
	for _, v := range arr {
		s += strconv.Itoa(v)
	}
	return s
}
