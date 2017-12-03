package main

import "fmt"

const input = 277678 // your puzzle input

func main() {

	// Right, Up, Left, Down
	directions := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	dir := 0

	// how many squares to move after changing directions
	hops := 1

	// arbitrarily large enough grid so i don't need to bounds check
	gridSize := 20
	offset := gridSize / 2

	// start in the middle
	x, y := offset, offset
	grid := make([][]int, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}

	// initialize center square to 1
	sum := 1
	grid[offset][offset] = sum

	for sum < input {
		for i := 0; i < hops && sum < input; i++ {
			x += directions[dir][0]
			y += directions[dir][1]
			sum = addAdjacent(x, y, grid)
			grid[y][x] = sum
		}

		// increase number of cells moved every other "line"
		if dir%2 == 1 {
			hops++
		}
		dir = (dir + 1) % 4 // change direction
	}

	// we found the target
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(sum)

}

func addAdjacent(x, y int, grid [][]int) int {
	//above
	sum := grid[y-1][x-1] + grid[y-1][x] + grid[y-1][x+1]
	// current
	sum += grid[y][x-1] + grid[y][x+1]
	//below
	sum += grid[y+1][x-1] + grid[y+1][x] + grid[y+1][x+1]
	return sum
}
