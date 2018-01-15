package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type particle struct {
	pos, vel, acc []int
}

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	parts := make([]particle, 0)
	for sc.Scan() {
		arr := strings.Split(sc.Text(), ", ")
		p := parse(arr[0])
		v := parse(arr[1])
		a := parse(arr[2])
		parts = append(parts, particle{p, v, a})
	}

	for {
		positions := make(map[string]int)
		toRemove := make(map[int]bool)
		// see overlaps
		for i, part := range parts {
			posKey := toKey(part.pos)

			if positions[posKey] > 0 {
				toRemove[i] = true
				toRemove[positions[posKey]-1] = true
				fmt.Println("collision")
			}
			positions[posKey] = i + 1
		}
		// remove overlaps
		newParts := make([]particle, 0)
		for i, part := range parts {
			if !toRemove[i] {
				newParts = append(newParts, part)
			}
		}
		parts = newParts

		// udpdate
		for _, part := range parts {
			//update velocity
			for i := 0; i < 3; i++ {
				part.vel[i] += part.acc[i]
			}
			//update position
			for i := 0; i < 3; i++ {
				part.pos[i] += part.vel[i]
			}
		}
		fmt.Println(len(parts))

	}
}

func toKey(arr []int) string {
	return strconv.Itoa(arr[0]) + "," + strconv.Itoa(arr[1]) + "," + strconv.Itoa(arr[2])
}

func parse(vec string) []int {
	arr := strings.Split(vec[3:len(vec)-1], ",")
	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])
	z, _ := strconv.Atoi(arr[2])
	return []int{x, y, z}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
