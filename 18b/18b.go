package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	inputs := make([]string, 0)
	for sc.Scan() {
		inputs = append(inputs, sc.Text())
	}

	// two different groups of registers
	m0 := map[string]int{"p": 0}
	m1 := map[string]int{"p": 1}

	// 2 buffered channels for communicating
	in := make(chan int, 10000)
	out := make(chan int, 10000)

	// Wait for both processes to finish or deadlock
	var wg sync.WaitGroup
	sends0, sends1 := 0, 0
	wg.Add(2)
	go func() {
		defer wg.Done()
		prog(inputs, m0, in, out, &sends0, 0)
	}()
	go func() {
		defer wg.Done()
		prog(inputs, m1, out, in, &sends1, 1)
	}()

	wg.Wait()
	fmt.Println(sends1)
}

func prog(inputs []string, m map[string]int, in, out chan int, sends *int, id int) {

	for i := 0; i >= 0 && i < len(inputs); i++ {
		cmd := inputs[i]
		arr := strings.Split(cmd, " ")
		reg := arr[1]
		switch arr[0] {
		case "set":
			m[reg] = val(arr[2], m)
		case "add":
			m[reg] += val(arr[2], m)
		case "mul":
			m[reg] *= val(arr[2], m)
		case "mod":
			v := val(arr[2], m)
			if v != 0 {
				m[reg] %= v
			}
		case "rcv":
			// Timeout channel to detect deadlock
			timeout := make(chan bool, 1)
			go func() {
				time.Sleep(1 * time.Second)
				timeout <- true
			}()
			// If we block for more than a second assume we are deadlocked and return
			select {
			case v := <-in:
				m[reg] = v
			case <-timeout:
				return
			}
		case "jgz":
			x := val(reg, m)
			y := val(arr[2], m)
			if x > 0 {
				i = (i + y) - 1
			}
		case "snd":
			out <- val(reg, m)
			*sends++
		default:
			log.Fatal("invalid")

		}
	}
}

// return register value if register. else return int value
func val(str string, m map[string]int) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return m[str]
	}
	return v

}
