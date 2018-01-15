package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type particle struct {
	pos, vel, acc int
}

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	i := 0
	bestIndex := 0
	bestParticle := particle{99999, 99999, 99999}
	for sc.Scan() {

		vectors := strings.Split(sc.Text(), ", ")

		pos := parse(vectors[0][3 : len(vectors[0])-1])
		vel := parse(vectors[1][3 : len(vectors[1])-1])
		acc := parse(vectors[2][3 : len(vectors[2])-1])
		part := particle{val(pos), val(vel), val(acc)}
		if Less(part, bestParticle) {
			bestParticle = part
			bestIndex = i
		}
		i++
	}
	fmt.Println(bestParticle)
	fmt.Println(bestIndex)
}

func Less(i, j particle) bool {
	if i.acc < j.acc {
		return true
	}
	if i.acc > j.acc {
		return false
	}
	if i.vel < j.vel {
		return true
	}
	if i.vel > j.vel {
		return false
	}
	if i.pos < j.pos {
		return true
	}

	return false
}

func parse(vec string) []int {
	arr := strings.Split(vec, ",")
	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])
	z, _ := strconv.Atoi(arr[2])
	return []int{x, y, z}
}

func val(arr []int) int {
	return abs(arr[0]) + abs(arr[1]) + abs(arr[2])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
