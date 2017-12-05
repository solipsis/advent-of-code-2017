package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	// read in data
	arr := make([]int, 0)
	for sc.Scan() {
		i, _ := strconv.Atoi(sc.Text())
		arr = append(arr, i)
	}

	i := 0
	jumps := 0
	// follow instructions until we are out of the bounds of the array
	for i < len(arr) {
		temp := arr[i] + i
		if arr[i] >= 3 {
			arr[i]--
		} else {
			arr[i]++
		}
		i = temp
		jumps++
	}

	fmt.Println(jumps)
}
