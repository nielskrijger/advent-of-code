package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	char       rune
	neighbours []*node
}

func main() {
	start, end := loadData("2022/12a/sample.txt")
	fmt.Printf("Answer: %d", bfs(start, end))
}

func bfs(start *node, end *node) int {
	visited := make(map[*node]int)
	visited[start] = 0

	queue := make(chan *node, 100) // 100 is arbitrary, should suffice for our grid size
	defer close(queue)
	queue <- start

	for n := range queue {
		if n == end {
			return visited[n]
		}

		for _, neighbour := range n.neighbours {
			if _, ok := visited[neighbour]; !ok {
				visited[neighbour] = visited[n] + 1
				queue <- neighbour
			}
		}
	}

	return -1
}

func loadData(file string) (*node, *node) {
	var start, end *node
	data, _ := os.ReadFile(file)
	matrix := make([][]*node, 0)
	for r, line := range strings.Split(string(data), "\n") {
		matrix = append(matrix, []*node{})
		for _, v := range line {
			n := &node{char: v}
			if v == 'S' {
				start = n
				n.char = 'a'
			} else if v == 'E' {
				end = n
				n.char = 'z'
			}
			matrix[r] = append(matrix[r], n)
		}
	}
	updateNeighbours(matrix)
	return start, end
}

func updateNeighbours(matrix [][]*node) {
	appendIfReachable := func(n *node, r int, c int) {
		if r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[r]) && matrix[r][c].char-n.char <= 1 {
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
