package main

import (
	"fmt"
	"image"
	"math"
	"os"
)

const MaxRounds = 1_000_000_000_000

var rocks = [][]image.Point{
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},         // horizontal line
	{{1, 2}, {0, 1}, {1, 1}, {2, 1}, {1, 0}}, // plus
	{{2, 2}, {2, 1}, {0, 0}, {1, 0}, {2, 0}}, // L in reverse
	{{0, 3}, {0, 2}, {0, 1}, {0, 0}},         // vertical line
	{{0, 1}, {1, 1}, {0, 0}, {1, 0}},         // block
}

func main() {
	jets, _ := os.ReadFile("2022/17b/input.txt")

	grid := &Grid{occupied: map[image.Point]bool{}}
	moved, height, j, totalCycleHeight := false, 0, 0, 0
	states := make([]State, 0)
	lastRound := MaxRounds

	for i := 0; i < lastRound; i++ {
		rock := rocks[i%len(rocks)]
		rock, _ = grid.moveRock(rock, image.Pt(2, height+3)) // start position

		for {
			jet := jets[j%len(jets)]
			j++
			rock, _ = grid.moveRock(rock, image.Pt(int(jet)-61, 0))
			if rock, moved = grid.moveRock(rock, image.Pt(0, -1)); !moved { // rock has settled
				if maxY(rock)+1 > height {
					height = maxY(rock) + 1
				}
				for _, p := range rock {
					grid.occupied[p] = true
				}
				if totalCycleHeight > 0 {
					break
				}

				states = append(states, NewState(maxY(rock), grid.occupied, height))
				if cycleOffset := findCycle(states, 10); cycleOffset > -1 {
					cycleHeight := states[len(states)-1].height - states[len(states)-cycleOffset-1].height
					remainingRounds := MaxRounds - i
					totalCycleHeight = remainingRounds / cycleOffset * cycleHeight
					lastRound = i + remainingRounds%cycleOffset
				}
				break
			}
		}
	}
	fmt.Printf("Answer: %d", height+totalCycleHeight)
}

type State struct {
	hash   int
	height int
}

func NewState(y int, occupiedPts map[image.Point]bool, height int) State {
	hash := 0
	for x := 0; x < 7; x++ {
		if _, ok := occupiedPts[image.Pt(x, y)]; ok {
			hash += int(math.Pow(2, float64(x+1)))
		}
	}
	return State{hash, height}
}

func findCycle(states []State, repeat int) int {
	for i := repeat; i < len(states)-repeat; i++ {
		if isCycle(states[len(states)-repeat:], states[len(states)-i-repeat:len(states)-i]) {
			return i
		}
	}
	return -1
}

func isCycle(first []State, second []State) bool {
	for i, m := range first {
		if second[i].hash != m.hash {
			return false
		}
	}
	return true
}

func maxY(pts []image.Point) (maxY int) {
	for _, p := range pts {
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return maxY
}

type Grid struct {
	occupied map[image.Point]bool
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
