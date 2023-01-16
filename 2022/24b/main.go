package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func main() {
	maps := generateAllMaps(loadData("2022/24b/sample.txt"))
	start := image.Pt(1, 0)
	end := image.Pt(maps[0].max.X-1, maps[0].max.Y)
	fmt.Printf("Answer: %+v", findRoute(maps, findRoute(maps, findRoute(maps, 0, start, end), end, start), start, end))
}

var deltas = map[uint8]image.Point{
	'>': image.Pt(1, 0), 'v': image.Pt(0, 1), '<': image.Pt(-1, 0), '^': image.Pt(0, -1), '.': image.Pt(0, 0),
}

type State struct {
	time  int
	point image.Point
}

func findRoute(maps []*Map, t int, start image.Point, end image.Point) int {
	seen := make(map[State]bool)
	queue := []State{{t, start}}
	seen[State{0, start}] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		m := maps[cur.time%len(maps)]

		for _, d := range deltas {
			np := cur.point.Add(d)
			s := State{cur.time + 1, np}
			if np == end {
				return cur.time
			}
			if np != start && (np.X <= 0 || np.Y <= 0 || np.X >= m.max.X || np.Y >= m.max.Y) {
				continue
			}
			if _, ok := seen[s]; ok {
				continue
			}
			if _, ok := m.points[np]; ok {
				continue
			}
			seen[s] = true
			queue = append(queue, s)
		}
	}
	return -1
}

type Points map[image.Point][]uint8

type Map struct {
	points Points
	max    image.Point
}

func (m *Map) updateBlizzards() *Map {
	newPoints := make(Points)
	for point, values := range m.points {
		for _, e := range values {
			if e > '.' {
				newPoint := m.nextPoint(point, deltas[e])
				if _, ok := newPoints[newPoint]; !ok {
					newPoints[newPoint] = make([]uint8, 0)
				}
				newPoints[newPoint] = append(newPoints[newPoint], e)
			}
		}
	}
	return &Map{newPoints, m.max}
}

func (m *Map) nextPoint(p image.Point, delta image.Point) image.Point {
	for {
		p = p.Add(delta)
		p.X = (p.X + m.max.X) % m.max.X
		p.Y = (p.Y + m.max.Y) % m.max.Y
		if p.X > 0 && p.Y > 0 {
			return p
		}
	}
}

func generateAllMaps(m *Map) []*Map {
	lookup := make(map[string]bool)
	maps := make([]*Map, 0)
	for {
		if _, ok := lookup[m.String()]; ok {
			break
		}
		lookup[m.String()] = true
		maps = append(maps, m)
		m = m.updateBlizzards()
	}
	return maps
}

func (m *Map) String() string {
	r := ""
	for y := 0; y <= m.max.Y; y++ {
		for x := 0; x <= m.max.X; x++ {
			if v, ok := m.points[image.Pt(x, y)]; ok {
				if len(v) > 1 {
					r += strconv.Itoa(len(v))
				} else {
					r += string(v[0])
				}
			} else if x == 0 || y == 0 || x == m.max.X || y == m.max.Y {
				r += "#"
			} else {
				r += "."
			}
		}
		r += "\n"
	}
	return r
}

func (m *Map) updateMax(x, y int) {
	if x > m.max.X {
		m.max.X = x
	}
	if y > m.max.Y {
		m.max.Y = y
	}
}

func loadData(filename string) *Map {
	m := &Map{points: map[image.Point][]uint8{}}
	data, _ := os.ReadFile(filename)
	for y, line := range strings.Split(string(data), "\n") {
		for x := 0; x < len(line); x++ {
			if line[x] != '.' && line[x] != '#' {
				m.points[image.Pt(x, y)] = make([]uint8, 1, 1)
				m.points[image.Pt(x, y)][0] = line[x]
			}
			m.updateMax(x, y)
		}
	}
	return m
}
