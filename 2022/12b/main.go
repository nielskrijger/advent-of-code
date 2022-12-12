package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	symbol     rune
	elevation  rune
	neighbours []*node
}

func main() {
	end := loadData("2022/12b/input.txt")
	fmt.Printf("Answer: %d", bfs(end, 'a')) // Start at the end and work backwards so it does one traversal
}

func bfs(start *node, findSymbol rune) int {
	type nodeLevel struct {
		node  *node
		depth int
	}

	visited := make(map[*node]bool)
	visited[start] = true

	queue := make(chan nodeLevel, 100) // 100 is arbitrary, should suffice for our grid size
	defer close(queue)
	queue <- nodeLevel{node: start, depth: 0}

	for len(queue) > 0 {
		n := <-queue
		if n.node.symbol == findSymbol {
			return n.depth
		}

		for _, neighbour := range n.node.neighbours {
			if _, ok := visited[neighbour]; !ok {
				visited[neighbour] = true
				queue <- nodeLevel{node: neighbour, depth: n.depth + 1}
			}
		}
	}

	return -1
}

func loadData(file string) *node {
	var end *node
	data, _ := os.ReadFile(file)
	matrix := make([][]*node, 0)
	for r, line := range strings.Split(string(data), "\n") {
		matrix = append(matrix, []*node{})
		for _, v := range line {
			n := &node{symbol: v, elevation: v}
			if v == 'S' {
				n.elevation = 'a'
			} else if v == 'E' {
				end = n
				n.elevation = 'z'
			}
			matrix[r] = append(matrix[r], n)
		}
	}
	updateNeighbours(matrix)
	return end
}

func updateNeighbours(matrix [][]*node) {
	appendIfReachable := func(n *node, r int, c int) {
		if r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[r]) && matrix[r][c].elevation-n.elevation >= -1 {
			n.neighbours = append(n.neighbours, matrix[r][c])
		}
	}
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			appendIfReachable(matrix[r][c], r-1, c)
			appendIfReachable(matrix[r][c], r+1, c)
			appendIfReachable(matrix[r][c], r, c-1)
			appendIfReachable(matrix[r][c], r, c+1)
		}
	}
}
