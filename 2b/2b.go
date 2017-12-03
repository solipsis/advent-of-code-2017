package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt") // ignoring error
	sc := bufio.NewScanner(file)

	sum := 0
	for sc.Scan() {
		words := strings.Split(sc.Text(), "\t")
		// convert input to int array
		nums := make([]int, 0, len(words))
		for _, w := range words {
			i, _ := strconv.Atoi(w)
			nums = append(nums, i)
		}

		// look for numbers that devide eachother
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]%nums[j] == 0 {
					sum += nums[i] / nums[j]
				} else if nums[j]%nums[i] == 0 {
					sum += nums[j] / nums[i]
				}
			}
		}
	}
	fmt.Println(sum)

}
