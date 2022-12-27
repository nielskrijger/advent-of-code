package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	X, Y, Z int
}

func (p Point) Add(a Point) Point {
	return Point{X: p.X + a.X, Y: p.Y + a.Y, Z: p.Z + a.Z}
}

func (p Point) Sub(a Point) Point {
	return Point{X: p.X - a.X, Y: p.Y - a.Y, Z: p.Z - a.Z}
}

func (p Point) IsNegative() bool {
	return p.X < 0 || p.Y < 0 || p.Z < 0
}

func main() {
	points, min, max := loadData("2022/18b/answer.txt")

	// Expand search space by 1 so we scan also the outer edges
	min = min.Sub(Point{1, 1, 1})
	max = max.Add(Point{1, 1, 1})

	edges := 0
	visited := make(map[Point]bool)
	deltas := []Point{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}

	queue := []Point{min}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, delta := range deltas {
			next := p.Add(delta)
			if _, ok := points[next]; ok {
				edges++
			} else if _, ok := visited[next]; !ok && !next.Sub(min).IsNegative() && !max.Sub(next).IsNegative() {
				queue = append(queue, next)
				visited[next] = true
			}
		}
	}
	fmt.Printf("Answer: %d", edges)
}

func loadData(filename string) (map[Point]bool, Point, Point) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	min := Point{X: math.MaxInt, Y: math.MaxInt, Z: math.MaxInt}
	max := Point{X: math.MinInt, Y: math.MinInt, Z: math.MinInt}
	lava := make(map[Point]bool)
	for _, line := range lines {
		var p Point
		_, _ = fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		lava[p] = true
		min = Point{Min(min.X, p.X), Min(min.Y, p.Y), Min(min.Z, p.Z)}
		max = Point{Max(max.X, p.X), Max(max.Y, p.Y), Max(max.Z, p.Z)}
	}
	return lava, min, max
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
