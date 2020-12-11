package main

import "fmt"

func main() {
	nodes := []bool{false, true}
	// links := map[byte][]byte{
	// 	0: {1},
	// 	1: {0},
	// }

	var nextNodes []bool
	for _, node := range nodes {
		nextNodes = append(nextNodes, node)
	}

	fmt.Println(nextNodes)
}

func nextBit(x, y bool) bool {
	switch {
	case x == false && y == false:
		return false
	case x == false && y == true:
		return true
	case x == true && y == false:
		return true
	case x == true && y == true:
		return true
	}
	return false
}
