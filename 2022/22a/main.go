package main

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	m, steps := loadData("2022/22a/sample.txt")
	facing := 0
	current := m.findWrapAroundPoint(image.Pt(1, 1), image.Pt(1, 0))
	for _, step := range steps {
		current = m.findFinishingPoint(current, step[0], facing)
		if step[1] == 'L' {
			facing = facing - 1
			if facing < 0 {
				facing = 3
			}
		} else if step[1] == 'R' {
			facing = (facing + 1) % 4
		}
	}
	fmt.Printf("Answer: %d", 1000*current.Y+4*current.X+facing)
}

type Map struct {
	grid map[image.Point]uint8
	max  image.Point
}

func (m *Map) setPoint(p image.Point, c uint8) {
	m.grid[p] = c
	if p.X > m.max.X {
		m.max = image.Pt(p.X, m.max.Y)
	}
	if p.Y > m.max.Y {
		m.max = image.Pt(m.max.X, p.Y)
	}
}

func (m *Map) findWrapAroundPoint(p image.Point, delta image.Point) image.Point {
	for {
		p = p.Add(delta)
		if p.X > m.max.X {
			p = image.Pt(1, p.Y)
		}
		if p.Y > m.max.Y {
			p = image.Pt(p.X, 1)
		}
		if p.X < 1 {
			p = image.Pt(m.max.X, p.Y)
		}
		if p.Y < 1 {
			p = image.Pt(p.X, m.max.Y)
		}
		if _, ok := m.grid[p]; ok {
			return p
		}
	}
}

var deltas = map[int]image.Point{0: image.Pt(1, 0), 1: image.Pt(0, 1), 2: image.Pt(-1, 0), 3: image.Pt(0, -1)}

func (m *Map) findFinishingPoint(current image.Point, step uint8, facing int) image.Point {
	for i := uint8(0); i < step; i++ {
		next := current.Add(deltas[facing])
		v, ok := m.grid[next]
		if !ok {
			next = m.findWrapAroundPoint(current, deltas[facing])
			v, _ = m.grid[next]
		}
		if v == '#' {
			return current
		}
		current = next
	}
	return current
}

func loadData(filename string) (*Map, [][2]uint8) {
	m := &Map{grid: make(map[image.Point]uint8)}
	data, _ := os.ReadFile(filename)
	groups := strings.Split(string(data), "\n\n")
	for r, line := range strings.Split(groups[0], "\n") {
		for c := 0; c < len(line); c++ {
			if line[c] != ' ' {
				m.setPoint(image.Pt(c+1, r+1), line[c])
			}
		}
	}

	r := regexp.MustCompile(`([0-9]+)([LR]{1})*`)
	directions := make([][2]uint8, 0)
	for _, submatch := range r.FindAllStringSubmatch(groups[1], -1) {
		num, _ := strconv.Atoi(submatch[1])
		direction := [2]uint8{uint8(num), 0}
		if len(submatch[2]) == 1 {
			direction[1] = submatch[2][0]
		}
		directions = append(directions, direction)
	}

	return m, directions
}
