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

	// count the number of distinct parent nodes
	parents := make(map[int]bool)
	for _, v := range m {
		parents[head(v).val] = true
	}

	fmt.Println(len(parents))
}

// union two trees of nodes
func attach(a, b *node) {

	ha := head(a)
	hb := head(b)
	// they are already part of the same set
	if ha.val == hb.val {
		return
	}
	// union the two trees by attaching the head of a to b
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
