package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	val    int
	parent *node
}

func main() {
	file, _ := os.Open("input.txt")
	sc := bufio.NewScanner(file)

	m := make(map[int]*node)
	i := 0
	for sc.Scan() {
		arr := strings.Split(sc.Text(), " ")

		// create any nodes if they don't exist and attach them to this node
		m[i] = fetchOrCreate(i, m)
		for _, v := range arr[2:] {
			num, _ := strconv.Atoi(strings.Replace(v, ",", "", 1))
			m[num] = fetchOrCreate(num, m)
			attach(m[i], m[num])
		}
		i++
	}

	// every node that has the same parent as node zero is in the same set
	result := 0
	goal := head(m[0]).val
	for x := 1; x < i; x++ {
		if head(m[x]).val == goal {
			result++
		}
	}

	// add 1 because zero counts as being connected to itself
	fmt.Println(result + 1)
}

// union two trees of nodes
func attach(a, b *node) {

	ha := head(a)
	hb := head(b)
	// they are already part of the same set
	if ha.val == hb.val {
		return
	}
	// join the sets by setting a's parent to b
	ha.parent = b
}

// return the head of a tree
func head(n *node) *node {
	// base case
	if n == nil || n.parent == nil {
		return n
	}
	return head(n.parent)
}

func fetchOrCreate(i int, m map[int]*node) *node {
	if v, ok := m[i]; ok {
		return v
	}
	return &node{val: i}
}
