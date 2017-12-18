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

	sum := 0
	for i := 0; i <= end; i++ {
		// no scanner in this column
		if m[i] == 0 {
			continue
		}

		// calculate the period of each scanner
		cycle := (m[i] - 1) * 2
		if i%cycle == 0 {
			sum += i * m[i]
		}
	}

	fmt.Println(sum)
}
