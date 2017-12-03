package main

import "fmt"

const input = 277678 // your puzzle input

func main() {

	// Right, Up, Left, Down
	directions := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	dir := 0

	// how many squares to move after changing directions
	hops := 1

	// how many squares we've filled
	moves := 1
	x, y := 0, 0
	for moves < input {
		for i := 0; i < hops && moves < input; i++ {
			x += directions[dir][0]
			y += directions[dir][1]
			moves++
		}
		// every other "line" increase number of squares we add
		if dir%2 == 0 {
			hops++
		}
		dir = (dir + 1) % 4
	}
	// just add our x and y distances from the center
	fmt.Println(abs(x) + abs(y))

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
