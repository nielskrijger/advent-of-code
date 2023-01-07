package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strings"
)

var directions = [][3]image.Point{
	{image.Pt(0, -1), image.Pt(1, -1), image.Pt(-1, -1)}, // N, NE, or NW
	{image.Pt(0, 1), image.Pt(1, 1), image.Pt(-1, 1)},    // S, SE, or SW
	{image.Pt(-1, 0), image.Pt(-1, -1), image.Pt(-1, 1)}, // W, NW, or SW
	{image.Pt(1, 0), image.Pt(1, -1), image.Pt(1, 1)},    // E, NE, or SE
}

func main() {
	m := loadData("2022/23a/sample.txt")
	for r := 0; r < 10; r++ {
		m = updateSpots(m, chooseSpots(m, r))
	}
	fmt.Printf("Answer: %d", m.FreeSpaces())
}

func chooseSpots(m Map, round int) map[image.Point][]image.Point {
	spots := make(map[image.Point][]image.Point)
	for pos := range m {
		allFree := true
		var p *image.Point

	skip:
		for i := 0; i < len(directions); i++ {
			direction := directions[(i+round)%len(directions)]
			for _, delta := range direction {
				if _, ok := m[pos.Add(delta)]; ok {
					allFree = false
					continue skip
				}
			}
			if p == nil {
				elm := pos.Add(direction[0])
				p = &elm
			}
		}
		if !allFree && p != nil { // p is nil if all surrounding directions are occupied
			if _, ok := spots[*p]; !ok {
				spots[*p] = make([]image.Point, 0)
			}
			spots[*p] = append(spots[*p], pos)
		}
	}
	return spots
}

func updateSpots(m map[image.Point]bool, spots map[image.Point][]image.Point) map[image.Point]bool {
	for newSpot, elves := range spots {
		if len(elves) > 1 {
			continue // Skip because multiple elves chose the same spot
		}
		delete(m, elves[0])
		m[newSpot] = true
	}
	return m
}

type Map map[image.Point]bool

func (m Map) SmallestRectangle() (image.Point, image.Point) {
	min, max := image.Pt(math.MaxInt, math.MaxInt), image.Pt(math.MinInt, math.MinInt)
	for k := range m {
		max = image.Pt(Max(max.X, k.X), Max(max.Y, k.Y))
		min = image.Pt(Min(min.X, k.X), Min(min.Y, k.Y))
	}
	return min, max
}

func (m Map) FreeSpaces() int {
	min, max := m.SmallestRectangle()
	free := 0
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			if _, ok := m[image.Pt(x, y)]; !ok {
				free++
			}
		}
	}
	return free
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

func loadData(filename string) Map {
	m := make(map[image.Point]bool)
	data, _ := os.ReadFile(filename)
	for r, line := range strings.Split(string(data), "\n") {
		for c := 0; c < len(line); c++ {
			if line[c] == '#' {
				m[image.Pt(c, r)] = true
			}
		}
	}
	return m
}
