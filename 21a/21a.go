package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Rule struct {
	rule, res [][]string
}

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)
	rules := make([]Rule, 0)
	for sc.Scan() {
		arr := strings.Split(sc.Text(), " => ")
		ruleArr := strings.Split(arr[0], "/")
		rule := make([][]string, 0)
		for i := 0; i < len(ruleArr); i++ {
			rule = append(rule, strings.Split(ruleArr[i], ""))
		}
		resArr := strings.Split(arr[1], "/")
		res := make([][]string, 0)
		for i := 0; i < len(resArr); i++ {
			res = append(res, strings.Split(resArr[i], ""))
		}

		rules = append(rules, Rule{rule, res})
	}

	start := [][]string{[]string{".", "#", "."}, []string{".", ".", "#"}, []string{"#", "#", "#"}}
	// how many expansion iterations
	for t := 0; t < 18; t++ {
		// split into size 3 chunks
		if len(start)%3 == 0 && len(start)%2 != 0 {
			newGrid := make([][]string, 4*(len(start)/3))

			for i := 0; i < len(start)/3; i++ {
				for j := 0; j < len(start)/3; j++ {
					block := make([][]string, 3)
					for x := 0; x < 3; x++ {
						temp := start[i*3+x][j*3 : j*3+3]
						block[x] = make([]string, len(temp))
						copy(block[x], temp)
					}

					for idx, line := range match(block, rules) {
						newGrid[i*4+idx] = append(newGrid[i*4+idx], line...)
					}
				}
			}
			start = newGrid

		} else {

			newGrid := make([][]string, 3*(len(start)/2))
			for i := 0; i < len(start)/2; i++ {
				for j := 0; j < len(start)/2; j++ {
					block := make([][]string, 2)
					for x := 0; x < 2; x++ {
						temp := start[i*2+x][j*2 : j*2+2]
						block[x] = make([]string, len(temp))
						copy(block[x], temp)
					}

					for idx, line := range match(block, rules) {
						newGrid[i*3+idx] = append(newGrid[i*3+idx], line...)
					}
				}
			}
			start = newGrid
		}
		print(start)
	}

	count := 0
	for _, row := range start {
		for _, c := range row {
			if c == "#" {
				count++
			}
		}
	}
	fmt.Println(count)

}

// Check if a block matches any rules
func match(block [][]string, rules []Rule) [][]string {
	for num, rule := range rules {
		if len(block) != len(rule.rule) {
			continue
		}
		for r := 0; r < 4; r++ {

			match := true
			for i := 0; i < len(rule.rule); i++ {
				for j := 0; j < len(rule.rule[0]); j++ {
					if block[i][j] != rule.rule[i][j] {
						match = false
					}
				}
			}

			if match {
				//print(block)
				fmt.Println("using rule: ", num)
				return rule.res
			}
			rotate(block)
		}
		for _, line := range block {
			reverse(line)
		}
		for r := 0; r < 4; r++ {

			match := true
			for i := 0; i < len(rule.rule); i++ {
				for j := 0; j < len(rule.rule[0]); j++ {
					if block[i][j] != rule.rule[i][j] {
						match = false
					}
				}
			}

			if match {
				//print(block)
				fmt.Println("using rule: ", num)
				return rule.res
			}
			rotate(block)
		}
	}
	log.Fatal("No Match")
	return nil
}

func print(arr [][]string) {
	for i := 0; i < len(arr); i++ {
		s := ""
		for j := 0; j < len(arr[0]); j++ {
			s += arr[i][j]
		}
		fmt.Println(s)
	}
	fmt.Println()
}

// Rotate a block by transposing the matrix and reversing the rows
func rotate(arr [][]string) {
	transpose(arr)
	for r := 0; r < len(arr); r++ {
		reverse(arr[r])
	}
}

func transpose(arr [][]string) {
	for r := 0; r < len(arr)-1; r++ {
		for c := r; c < len(arr[0]); c++ {
			arr[r][c], arr[c][r] = arr[c][r], arr[r][c]
		}
	}
}

func reverse(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
