package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name     string
	weight   int
	children map[string]bool
}

func main() {

	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	// track all nodes by name
	m := make(map[string]*node)

	// create an entry for each node and populate its children
	for sc.Scan() {

		arr := strings.Split(sc.Text(), " ")
		weightStr := arr[1]
		weight, _ := strconv.Atoi(weightStr[1 : len(weightStr)-1]) // strip parenthesis

		// create node and children
		node := &node{name: arr[0], weight: weight, children: make(map[string]bool)}
		if len(arr) > 2 {
			for _, child := range arr[3:] {
				// strip trailing comma
				if strings.HasSuffix(child, ",") {
					child = child[:len(child)-1]
				}
				node.children[child] = true
			}
		}
		m[node.name] = node
	}

	// calculate the weight of every child starting with the root node we found in part A
	_, result := getWeight(m["bpvhwhh"], m)
	fmt.Println(result)

}

func getWeight(n *node, m map[string]*node) (sum, result int) {

	// we have found a solution. Exit recursion
	if result > 0 {
		return sum, result
	}

	// leaf node
	if len(n.children) == 0 {
		return n.weight, result
	}

	// get weight of each child.
	cWeights := make(map[string]int)
	for k := range n.children {
		cWeights[k], result = getWeight(m[k], m)
		// break out if we found a solution. Should probably find a cleaner way to do this
		if result > 0 {
			return sum, result
		}
	}

	// track how many children have each weight
	weightCounts := make(map[int][]string)
	for k, v := range cWeights {
		weightCounts[v] = append(weightCounts[v], k)
		sum += v
	}

	// if we have more than 1 result the entry with only 1 occurence is the offender
	if len(weightCounts) > 1 {
		correctWeight, incorrectWeight := 0, 0
		for k, v := range weightCounts {
			if len(v) == 1 {
				incorrectWeight = k
			} else {
				correctWeight = k
			}
		}
		result = correctWeight - incorrectWeight + m[weightCounts[incorrectWeight][0]].weight
		return sum, result
	}

	return sum + n.weight, result
}
