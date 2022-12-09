package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("2022/09a/sample.txt")
	lines := strings.Split(string(data), "\n")

	tail := image.Pt(0, 0)
	head := image.Pt(0, 0)
	visited := map[image.Point]bool{}

	change := map[rune]image.Point{'U': {0, 1}, 'D': {0, -1}, 'R': {1, 0}, 'L': {-1, 0}}

	for _, line := range lines {
		var direction rune
		var steps int
		_, _ = fmt.Sscanf(line, "%c %d", &direction, &steps)

		for i := 0; i < steps; i++ {
			head = head.Add(change[direction])

			diff := head.Sub(tail)
			if abs(diff.X) > 1 || abs(diff.Y) > 1 {
				tail = tail.Add(image.Pt(signum(diff.X), signum(diff.Y)))
			}

			visited[tail] = true
		}
	}

	fmt.Printf("Answer: %d", len(visited))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func signum(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
