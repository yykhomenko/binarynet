package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	nodesNum      = 10
	generationNum = 10
)

func main() {
	rand.Seed(time.Now().Unix())

	nodes := genNodes(nodesNum)
	links := genLinks()
	fmt.Println(nodes)

	for i := 0; i < generationNum; i++ {
		nodes = genNextNodes(nodes, links)
		fmt.Println(nodes)
	}
}

func genNodes(size int) []int {
	nodes := make([]int, size, size)
	for i := 0; i < size; i++ {
		nodes[i] = rand.Intn(2)
	}
	return nodes
}

func genLinks() map[int][]int {
	return map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {1, 1},
		3: {2, 4},
		4: {5, 0},
		5: {1, 3},
	}
}

func genNextNodes(nodes []int, links map[int][]int) (out []int) {
	for i := range nodes {
		out = append(out, nextBit(nodes, links[i]))
	}
	return
}

func nextBit(nodes []int, links []int) int {
	var p int
	for i, node := range links {
		p |= nodes[node] << i
	}

	switch p {
	case 0, 1:
		return 1
	default:
		return 0
	}
}
