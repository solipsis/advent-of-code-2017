package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// map to track register values
var m map[string]int

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	inputs := make([]string, 0)
	for sc.Scan() {
		inputs = append(inputs, sc.Text())
	}
	lastSound := 0
	m = make(map[string]int)
	for i := 0; i >= 0 && i < len(inputs); i++ {
		cmd := inputs[i]
		arr := strings.Split(cmd, " ")
		reg := arr[1]
		switch arr[0] { // switch on command name
		case "set":
			m[reg] = val(arr[2])
		case "add":
			m[reg] += val(arr[2])
		case "mul":
			m[reg] *= val(arr[2])
		case "mod":
			v := val(arr[2])
			if v != 0 {
				m[reg] %= v
			}
		case "rcv":
			if val(reg) > 0 {
				log.Fatal(lastSound) // we are done
			}
		case "jgz":
			x := val(reg)
			y := val(arr[2])
			if x > 0 {
				i = (i + y) - 1
			}
		case "snd":
			lastSound = val(reg)
		default:
			log.Fatal("invalid")

		}

	}
	fmt.Println("done")
}

// return register value if register. else return int value
func val(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return m[str]
	}
	return v
}
