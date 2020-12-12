package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

const (
	nodesNum      = 20
	linkNum       = 3
	generationNum = 100
)

func main() {
	rand.Seed(time.Now().Unix())

	nodes := genNodes(nodesNum)
	links := genLinks(nodesNum, linkNum)
	fmt.Println(links)
	fmt.Println(toString(nodes))

	for i := 0; i < generationNum; i++ {
		nodes = genNextNodes(nodes, links)
		fmt.Println(toString(nodes))
	}
}

func toString(nodes []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, data := range nodes {
		if data == 0 {
			buf.WriteByte('0')
		} else {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte(']')

	return buf.String()
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
		links[i] = genLink(size, linkNum, i)
	}

	return links
}

func genLink(size, linkNum, node int) []int {
	link := make([]int, linkNum, linkNum)
	for i := 0; i < linkNum; i++ {
	retry:
		n := rand.Intn(size)
		if n == node {
			goto retry
		}
		link[i] = n
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
	case 0, 3, 6, 7:
		return 1
	default:
		return 0
	}
}
