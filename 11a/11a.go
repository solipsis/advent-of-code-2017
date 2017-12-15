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

	sc.Scan()
	input := strings.Split(sc.Text(), ",")

	m := make(map[string]int)
	for _, dir := range input {
		m[dir]++
	}

	// cancel out complementary directions and count the remaining
	for m["nw"] > 0 && m["se"] > 0 {
		m["nw"]--
		m["se"]--
	}
	for m["ne"] > 0 && m["sw"] > 0 {
		m["ne"]--
		m["sw"]--
	}
	for m["nw"] > 0 && m["ne"] > 0 {
		m["n"]++
		m["nw"]--
		m["ne"]--
	}
	for m["sw"] > 0 && m["se"] > 0 {
		m["s"]++
		m["sw"]--
		m["se"]--
	}
	for m["n"] > 0 && m["s"] > 0 {
		m["n"]--
		m["s"]--
	}

	sum := 0
	for _, v := range m {
		sum += v
	}

	fmt.Println(sum)

}
