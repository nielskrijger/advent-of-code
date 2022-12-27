package main

import (
	"fmt"
	"image"
	"os"
)

var rocks = [][]image.Point{
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},         // horizontal line
	{{1, 2}, {0, 1}, {1, 1}, {2, 1}, {1, 0}}, // plus
	{{2, 2}, {2, 1}, {0, 0}, {1, 0}, {2, 0}}, // L in reverse
	{{0, 3}, {0, 2}, {0, 1}, {0, 0}},         // vertical line
	{{0, 1}, {1, 1}, {0, 0}, {1, 0}},         // block
}

func main() {
	jets, _ := os.ReadFile("2022/17a/sample.txt")

	grid := &Grid{occupied: map[image.Point]bool{}}
	moved, height, j := false, 0, 0
	for i := 0; i < 2022; i++ {
		rock := rocks[i%len(rocks)]
		rock, _ = grid.moveRock(rock, image.Pt(2, height+3)) // start position

		for {
			jet := jets[j%len(jets)]
			j++
			rock, _ = grid.moveRock(rock, image.Pt(int(jet)-61, 0))
			rock, moved = grid.moveRock(rock, image.Pt(0, -1))

			if !moved {
				if maxY(rock)+1 > height {
					height = maxY(rock) + 1
				}
				for _, p := range rock {
					grid.occupied[p] = true
				}
				break
			}
		}
	}

	fmt.Printf("Answer: %d", height)
}

type Grid struct {
	occupied map[image.Point]bool
}

func maxY(pts []image.Point) (maxY int) {
	for _, p := range pts {
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return maxY
}

func (g *Grid) moveRock(rock []image.Point, delta image.Point) ([]image.Point, bool) {
	newPos := make([]image.Point, 0, len(rock)-1)
	for _, point := range rock {
		p := point.Add(delta)
		if _, occupied := g.occupied[p]; occupied || p.X < 0 || p.X > 6 || p.Y < 0 {
			return rock, false // can't move
		}
		newPos = append(newPos, p)
	}
	return newPos, true
}
