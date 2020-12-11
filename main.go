package main

import "fmt"

func main() {
	nodes := genNodes()
	links := genLinks()
	fmt.Println(nodes)

	nextNodes := genNextNodes(nodes, links)
	nodes = nextNodes

	fmt.Println(nodes)
}

func genNodes() []int {
	return []int{0, 1, 0, 1}
}

func genLinks() map[int][]int {
	return map[int][]int{
		0: {1, 2},
		1: {0, 3},
		2: {0, 1},
		3: {2, 1},
	}
}

func genNextNodes(nodes []int, links map[int][]int) (out []int) {
	for i := range nodes {
		l := links[i]
		res := nextBit(l[0], l[1])
		out = append(out, res)
	}
	return
}

func nextBit(x, y int) int {
	switch {
	case x == 0 && y == 0:
		return 0
	case x == 0 && y == 1:
		return 1
	case x == 1 && y == 0:
		return 1
	case x == 1 && y == 1:
		return 1
	}
	return 0
}
