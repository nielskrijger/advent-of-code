package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	pos := findBeacon(loadSensors("2022/15b/sample.txt"))
	fmt.Printf("Answer: %d", pos.X*4_000_000+pos.Y)
}

func findBeacon(sensors []Sensor) image.Point {
	for _, s1 := range sensors {
	out:
		// We know the point must be on the boundary of one of the sensors since there is only a single
		// possible point in the [0,0][4000000,4000000] search space.
		for p := range s1.boundary() {
			if p.X < 0 || p.X > 4_000_000 || p.Y < 0 || p.Y > 4_000_000 {
				continue
			}
			for _, s2 := range sensors {
				if s1 == s2 {
					continue
				}
				if s2.inRange(p) {
					continue out
				}
			}
			return p
		}
	}

	panic("Did not find empty point")
}

type Sensor struct {
	point  image.Point
	beacon image.Point
}

func (s Sensor) distance(p image.Point) int {
	return abs(s.point.X-p.X) + abs(s.point.Y-p.Y)
}

func (s Sensor) inRange(p image.Point) bool {
	return s.distance(p) <= s.distance(s.beacon)
}

func (s Sensor) boundary() map[image.Point]bool {
	p := image.Pt(s.point.X-s.distance(s.beacon)-1, s.point.Y)
	points := map[image.Point]bool{}
	points[p] = true
	for _, d := range []image.Point{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}} {
		for i := 0; ; i++ {
			p = p.Add(d)
			points[p] = true
			if p.X == s.point.X || p.Y == s.point.Y {
				break
			}
		}
	}
	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func loadSensors(filename string) []Sensor {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	sensors := make([]Sensor, 0)
	for _, line := range lines {
		var s Sensor
		_, _ = fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.point.X, &s.point.Y, &s.beacon.X, &s.beacon.Y)
		sensors = append(sensors, s)
	}
	return sensors
}
