package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("2022/09b/sample.txt")
	lines := strings.Split(string(data), "\n")

	rope := make([]image.Point, 10)
	visited := map[image.Point]bool{}

	change := map[uint8]image.Point{'U': {0, 1}, 'D': {0, -1}, 'R': {1, 0}, 'L': {-1, 0}}

	for _, line := range lines {
		var direction uint8
		var steps int
		_, _ = fmt.Sscanf(line, "%c %d", &direction, &steps)

		for i := 0; i < steps; i++ {
			rope[0] = rope[0].Add(change[direction])

			for i := 1; i < len(rope); i++ {
				diff := rope[i-1].Sub(rope[i])
				if abs(diff.X) > 1 || abs(diff.Y) > 1 {
					rope[i] = rope[i].Add(image.Pt(signum(diff.X), signum(diff.Y)))
				}
			}

			visited[rope[9]] = true
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
