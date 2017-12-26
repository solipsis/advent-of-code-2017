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

	arr := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}

	sc.Scan()
	inputs := strings.Split(sc.Text(), ",")

	for _, v := range inputs {
		switch v[0] {
		case 's':
			spin(v[1:], arr)
		case 'x':
			exchange(v[1:], arr)
		case 'p':
			partner(v[1:], arr)

		}
	}
	for _, v := range arr {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}

// spin array by reversing 3 times
func spin(v string, arr []rune) {
	shift, _ := strconv.Atoi(v)
	shift = len(arr) - shift // change direction we are shifting
	reverse(arr[:shift])
	reverse(arr[shift:])
	reverse(arr)

}

func reverse(arr []rune) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func exchange(v string, arr []rune) {
	pair := strings.Split(v, "/")
	a, _ := strconv.Atoi(pair[0])
	b, _ := strconv.Atoi(pair[1])
	arr[a], arr[b] = arr[b], arr[a]
}

func partner(v string, arr []rune) {
	pair := strings.Split(v, "/")

	// find index of both characters
	ai, bi := 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == rune(pair[0][0]) {
			ai = i
		}
		if arr[i] == rune(pair[1][0]) {
			bi = i
		}
	}
	arr[ai], arr[bi] = arr[bi], arr[ai]
}
