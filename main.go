package main

import "fmt"

func main() {
	nodes := genNodes()
	links := genLinks()

	var nextNodes []int
	for i := range nodes {
		l := links[i]
		res := nextBit(l[0], l[1])
		nextNodes = append(nextNodes, res)
	}
	nodes = nextNodes

	fmt.Println(nextNodes)
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
