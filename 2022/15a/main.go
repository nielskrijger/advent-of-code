package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	sensors := loadSensors("2022/15a/sample.txt")

	hits := make(map[image.Point]int, 0)
	for x := -1_000_000; x < 6_000_000; x++ {
		for _, sensor := range sensors {
			p := image.Pt(x, 2_000_000)
			if sensor.beacon == p {
				hits[p] = 2
			} else if _, ok := hits[p]; !ok && sensor.inRange(p) {
				hits[p] = 1
			}
		}
	}

	total := 0
	for _, hit := range hits {
		if hit == 1 {
			total++
		}
	}

	fmt.Printf("Answer: %d", total)
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
