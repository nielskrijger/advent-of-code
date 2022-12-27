package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X, Y, Z int
}

func (p Point) Add(a Point) Point {
	return Point{X: p.X + a.X, Y: p.Y + a.Y, Z: p.Z + a.Z}
}

func main() {
	points := loadData("2022/18a/sample.txt")
	edges := 0
	mutations := []Point{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}
	for k := range points {
		for _, mut := range mutations {
			if _, ok := points[k.Add(mut)]; !ok {
				edges++
			}
		}
	}
	fmt.Printf("Answer: %d", edges)
}

func loadData(filename string) map[Point]bool {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	lava := make(map[Point]bool)
	for _, line := range lines {
		var p Point
		_, _ = fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		lava[p] = true
	}
	return lava
}
