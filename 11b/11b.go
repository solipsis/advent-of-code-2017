package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// give each direction in a hexagon a value 0-5
const (
	START = iota
	RIGHT
	DOWNRIGHT
	DOWN
	DOWNLEFT
	LEFT
)

// string label for each face of hexagon
var wheel = []string{"n", "ne", "se", "s", "sw", "nw"}

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	sc.Scan()
	input := strings.Split(sc.Text(), ",")

	max := 0
	m := make(map[string]int)
	for _, dir := range input {
		m[dir]++
		// reduce complementary directions
		reduce(m)

		// calculate distance from start and update max
		sum := 0
		for _, v := range m {
			sum += v
		}
		if sum > max {
			max = sum
		}
	}

	fmt.Println(max)
}

// reduce directions by cancelling out complementary directions
func reduce(m map[string]int) {
	// repeat this process 3 times shifting the wheel 1 space
	for i := 0; i < 3; i++ {
		// keep reducing until nothing less to reduce in this position
		for {
			// opposites
			if m[wheel[START]] > 0 && m[wheel[DOWN]] > 0 {
				m[wheel[START]]--
				m[wheel[DOWN]]--
				continue
			}
			// right side. combine the 2 to form one of direction between them
			if m[wheel[START]] > 0 && m[wheel[DOWNRIGHT]] > 0 {
				m[wheel[START]]--
				m[wheel[DOWNRIGHT]]--
				m[wheel[RIGHT]]++
				continue
			}
			// left side. combine the 2 to form one of direction between them
			if m[wheel[START]] > 0 && m[wheel[DOWNLEFT]] > 0 {
				m[wheel[START]]--
				m[wheel[DOWNLEFT]]--
				m[wheel[LEFT]]++
				continue
			}
			break
		}
		shift(wheel, 1)
	}

}

// shift the array i spaces by reversing 3 times
func shift(arr []string, i int) {
	reverse(arr[:i])
	reverse(arr[i:])
	reverse(arr)
}

func reverse(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
