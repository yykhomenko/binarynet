package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	nodesNum      = 50
	linkNum       = 2
	generationNum = 10
)

func main() {
	rand.Seed(time.Now().Unix())

	nodes := genNodes(nodesNum)
	links := genLinks(nodesNum, linkNum)
	fmt.Println(links)
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

func genLinks(size, linkNum int) [][]int {
	links := make([][]int, size, size)
	for i := 0; i < size; i++ {
		links[i] = genLink(size, linkNum)
	}

	return links
}

func genLink(size, linkNum int) []int {
	link := make([]int, linkNum, linkNum)
	for i := 0; i < linkNum; i++ {
		link[i] = rand.Intn(size)
	}
	return link
}

func genNextNodes(nodes []int, links [][]int) (out []int) {
	for i := range nodes {
		out = append(out, nextBit(nodes, links[i]))
	}
	return
}

func nextBit(nodes, links []int) int {
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
