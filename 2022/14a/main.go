package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Answer: %d", fallingSand(loadGrid("2022/14a/sample.txt")))
}

func loadGrid(filename string) (map[image.Point]int, int) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	grid := make(map[image.Point]int, 0)
	maxY := 0
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		for i := 0; i < len(points)-1; i++ {
			var a, b image.Point
			_, _ = fmt.Sscanf(points[i], "%d,%d", &a.X, &a.Y)
			_, _ = fmt.Sscanf(points[i+1], "%d,%d", &b.X, &b.Y)

			for d := image.Pt(sgn(b.X-a.X), sgn(b.Y-a.Y)); a != b.Add(d); a = a.Add(d) {
				grid[a] = 1
				if a.Y > maxY {
					maxY = a.Y
				}
			}
		}
	}
	return grid, maxY
}

func fallingSand(grid map[image.Point]int, bottomY int) int {
	stillSandCount := 0

	// Three possible moves in order of preference: down, left+down, right+down. Y increases in downwards direction.
	moves := []image.Point{image.Pt(0, 1), image.Pt(-1, 1), image.Pt(1, 1)}

out:
	for {
		sand := image.Pt(500, 0)
		for {
			for i, drop := range moves {
				if _, ok := grid[sand.Add(drop)]; !ok {
					sand = sand.Add(drop)
					if sand.Y > bottomY {
						return stillSandCount // Reached the bottom
					}
					break
				}

				if i == len(moves)-1 { // Exhausted all options, sand can't fall any further
					stillSandCount++
					grid[sand] = 2
					continue out
				}
			}
		}
	}
}

func sgn(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	}
	return 0
}
