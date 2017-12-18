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

	m := make(map[int]int)
	end := 0
	for sc.Scan() {
		arr := strings.Split(sc.Text(), ": ")
		index, _ := strconv.Atoi(arr[0])
		size, _ := strconv.Atoi(arr[1])
		m[index] = size
		end = index
	}

	// keep incrementing delay until we don't get caught
	delay := 0
	for isCaught(delay, end, m) {
		delay++
	}

	fmt.Println(delay)
}

func isCaught(delay, end int, m map[int]int) bool {

	i := delay
	for ; i-delay <= end; i++ {
		// no scanner in this column
		if m[i-delay] == 0 {
			continue
		}

		// calculate the period of the scanner
		cycle := (m[i-delay] - 1) * 2
		if i%cycle == 0 {
			return true
		}
	}
	return false
}
