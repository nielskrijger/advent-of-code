package main

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Sample edges
var edges = Edges{
	{{2, image.Pt(9, 1), image.Pt(9, 4)}, {3, image.Pt(5, 5), image.Pt(8, 5)}},
	{{3, image.Pt(9, 1), image.Pt(12, 1)}, {3, image.Pt(4, 5), image.Pt(1, 5)}},
	{{0, image.Pt(12, 1), image.Pt(12, 4)}, {0, image.Pt(16, 12), image.Pt(16, 9)}},
	{{0, image.Pt(12, 5), image.Pt(12, 8)}, {3, image.Pt(16, 9), image.Pt(13, 9)}},
	{{2, image.Pt(1, 5), image.Pt(1, 8)}, {1, image.Pt(16, 12), image.Pt(13, 12)}},
	{{1, image.Pt(1, 8), image.Pt(4, 8)}, {1, image.Pt(12, 12), image.Pt(9, 12)}},
	{{1, image.Pt(5, 8), image.Pt(8, 8)}, {1, image.Pt(9, 12), image.Pt(9, 9)}},
}

// Input edges
//var edges = Edges{
//	{{3, image.Pt(51, 1), image.Pt(100, 1)}, {2, image.Pt(1, 151), image.Pt(1, 200)}},
//	{{3, image.Pt(101, 1), image.Pt(150, 1)}, {1, image.Pt(1, 200), image.Pt(50, 200)}},
//	{{0, image.Pt(150, 1), image.Pt(150, 50)}, {0, image.Pt(100, 150), image.Pt(100, 101)}},
//	{{1, image.Pt(101, 50), image.Pt(150, 50)}, {0, image.Pt(100, 51), image.Pt(100, 100)}},
//	{{1, image.Pt(51, 150), image.Pt(100, 150)}, {0, image.Pt(50, 151), image.Pt(50, 200)}},
//	{{2, image.Pt(51, 1), image.Pt(51, 50)}, {2, image.Pt(1, 150), image.Pt(1, 101)}},
//	{{2, image.Pt(51, 51), image.Pt(51, 100)}, {3, image.Pt(1, 101), image.Pt(50, 101)}},
//}

func main() {
	m, steps := loadData("2022/22b/sample.txt")
	facing := 0
	current := image.Pt(9, 1) // (51,1) for input, (9,1) for sample
	for _, step := range steps {
		current, facing = m.findFinishingPoint(current, step[0], facing)
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

type Edge struct {
	facing int // Facing when leaving the edge: Right = 0, Down = 1, Left = 2, Up = 3
	start  image.Point
	end    image.Point
}

type Edges [][2]Edge

func (c Edges) findFoldingPoint(grid map[image.Point]uint8, p image.Point, facing int) (image.Point, int) {
	for _, pair := range c {
		newPoint, newFacing := c.findNewEdge(p, pair[0], pair[1], facing)
		if newFacing == -1 {
			newPoint, newFacing = c.findNewEdge(p, pair[1], pair[0], facing)
		}
		if newFacing > -1 {
			if v, _ := grid[newPoint]; v == '#' {
				return p, facing
			}
			return newPoint, newFacing
		}
	}
	panic("Didn't find new edge")
}

func (c Edges) findNewEdge(p image.Point, first, second Edge, facing int) (image.Point, int) {
	if first.facing == facing && isInLine(first.start, p, first.end) {
		dist := distance(first.start, p)
		newFacing := (second.facing + 2) % 4

		if second.start.X-second.end.X == 0 {
			if second.start.Y > second.end.Y {
				return image.Pt(second.start.X, second.start.Y-dist), newFacing
			}
			return image.Pt(second.start.X, second.start.Y+dist), newFacing
		}

		if second.start.X > second.end.X {
			return image.Pt(second.start.X-dist, second.start.Y), newFacing
		}
		return image.Pt(second.start.X+dist, second.start.Y), newFacing
	}

	return image.Point{}, -1
}

func isInLine(a, b, c image.Point) bool {
	if (a.X == b.X && b.X == c.X) || (a.Y == b.Y && b.Y == c.Y) {
		return distance(a, c)-distance(a, b) == distance(b, c)
	}
	return false
}

func distance(a, b image.Point) int {
	return abs(b.Y - a.Y + b.X - a.X) // always straight so no need to account for diagonals
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

var deltas = map[int]image.Point{0: image.Pt(1, 0), 1: image.Pt(0, 1), 2: image.Pt(-1, 0), 3: image.Pt(0, -1)}

func (m *Map) findFinishingPoint(current image.Point, step uint8, facing int) (image.Point, int) {
	for i := uint8(0); i < step; i++ {
		next := current.Add(deltas[facing])

		v, ok := m.grid[next]
		if !ok {
			next, facing = edges.findFoldingPoint(m.grid, current, facing)
			v, _ = m.grid[next]
		}
		if v == '#' {
			return current, facing
		}

		current = next
	}
	return current, facing
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
