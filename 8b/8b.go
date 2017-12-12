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

	// compute each operation
	max := 0
	m := make(map[string]int)
	for sc.Scan() {

		arr := strings.Split(sc.Text(), " ")
		reg, arg, amt, as, op, bs := arr[0], arr[1], arr[2], arr[4], arr[5], arr[6]
		a := m[as] // get register value for left operand
		b, _ := strconv.Atoi(bs)

		// check if the condition is true and apply the effect
		if cond(a, b, op) {
			v, _ := strconv.Atoi(amt)
			if arg == "inc" {
				m[reg] += v
			} else {
				m[reg] -= v
			}
			// update max
			if m[reg] > max {
				max = m[reg]
			}
		}
	}

	fmt.Println(max)
}

// check if the given condition is true
func cond(a, b int, op string) bool {
	switch op {
	case "<":
		return a < b
	case ">":
		return a > b
	case "<=":
		return a <= b
	case ">=":
		return a >= b
	case "!=":
		return a != b
	case "==":
		return a == b
	default:
		return false
	}
}
