package main

import "fmt"

const (
	nodesNum = 10
)

func main() {
	nodes := genNodes()
	links := genLinks()
	fmt.Println(nodes)

	for i := 0; i < nodesNum; i++ {
		nodes = genNextNodes(nodes, links)
		fmt.Println(nodes)
	}
}

func genNodes() []int {
	return []int{0, 1, 0, 1, 0, 1}
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
