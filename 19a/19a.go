package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	grid := make([][]string, 0)

	for sc.Scan() {
		grid = append(grid, strings.Split(sc.Text(), ""))
	}

	// find start
	r, c := 0, 0
	for i, v := range grid[0] {
		if v == "|" {
			c = i
		}
	}

	dirs := map[string][]int{"DOWN": {1, 0}, "UP": {-1, 0}, "RIGHT": {0, 1}, "LEFT": {0, -1}}
	cur := "DOWN"

	follow(r, c, &cur, dirs, grid)
}

func follow(r, c int, cur *string, dirs map[string][]int, grid [][]string) {
	cell := ""
	path := ""
	steps := 0
	//  keep reading until we get to a plus or are done
	for cell != " " {

		r += dirs[*cur][0]
		c += dirs[*cur][1]

		steps++
		cell = grid[r][c]
		// if the cell is a letter, add it to the path
		if cell != " " && cell != "+" && cell != "-" && cell != "|" {
			path += cell
		}
		// if the cell is a joint then we need to turn
		if cell == "+" {
			*cur = updateDirection(r, c, *cur, grid)
		}
	}
	fmt.Println(path)
}

func updateDirection(r, c int, dir string, grid [][]string) string {

	// Look for the path 90 degrees from the direction we are currently traveling
	if dir == "DOWN" || dir == "UP" {
		if c > 0 && grid[r][c-1] == "-" {
			return "LEFT"
		} else {
			return "RIGHT"
		}
	} else {
		if r > 0 && grid[r-1][c] == "|" {
			return "UP"
		} else {
			return "DOWN"
		}
	}
}
