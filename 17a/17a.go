package main

import "fmt"

const steps = 349 // puzzle input
func main() {
	fmt.Println("vim-go")

	pos := 0
	arr := make([]int, 1, 2017)

	for i := 1; i <= 2017; i++ {
		pos = (pos+steps)%len(arr) + 1
		// insert element at pos of value i
		arr = append(arr[:pos], append([]int{i}, arr[pos:]...)...)
	}
	fmt.Println(arr[pos+1]) // to lazy to bounds check

}
