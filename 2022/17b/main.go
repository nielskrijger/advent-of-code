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
	jets, _ := os.ReadFile("2022/17b/sample.txt")

	//cycleHeights := make([]int, 0)
	heights := make([]int, 1)

	grid := &Grid{occupied: map[image.Point]bool{}}
	moved, height, prevHeight, j := false, 0, 0, 0

	for i := 0; i < 2000; i++ {
		rock := rocks[i%len(rocks)]
		rock, _ = grid.moveRock(rock, image.Pt(2, height+3)) // start position
		grid.print(rock)

		for {
			c := findCycle(heights, 10)
			if c > -1 {
				fmt.Printf("Cycle found at %d\n", c)
				totalCycles := 2000 / i
				fmt.Printf("%d\n", (totalCycles-1)*height)
				return
			}
			//fmt.Printf("test: %v %v - %v %v\n", i%len(rocks), i%len(rocks) == 0, j%uint64(len(jets)), j%uint64(len(jets)) == 0)
			//if i%len(rocks) == 0 && j%uint64(len(jets)) == uint64(0) {
			//	println("-----------------")
			//	cycles++
			//	if len(cycleHeights) > 0 {
			//		cycleHeights = append(cycleHeights, height-cycleHeights[len(cycleHeights)-1])
			//		fmt.Printf("cycleHeights = %+v\n", cycleHeights)
			//		return
			//	} else {
			//		println(height)
			//		cycleHeights = append(cycleHeights, height)
			//	}
			//
			//	//
			//	//println("CYCLE")
			//	//totalCycles := 1000000000000 / i
			//	//fmt.Printf("height = %+v\n", height)
			//	//fmt.Printf("totalCycles = %+v\n", totalCycles)
			//	//height = (totalCycles - 1) * height
			//	//i = 1000000000000 - 1000000000000%i
			//	//fmt.Printf("i = %+v\n", i)
			//	//for x := 0; x < 7; x++ {
			//	//	grid.occupied[image.Pt(x, height)] = true // create new floor
			//	//}
			//	//return
			//	continue out // Do the last remainder of the last cycle
			//}
			jet := jets[j%len(jets)]
			j++
			grid.print(rock)
			rock, _ = grid.moveRock(rock, image.Pt(int(jet)-61, 0))
			grid.print(rock)
			rock, moved = grid.moveRock(rock, image.Pt(0, -1))

			grid.print(rock)
			if !moved {
				prevHeight = height
				if maxY(rock)+1 > height {
					height = maxY(rock) + 1
				}
				heights = append(heights, height-prevHeight)
				for _, p := range rock {
					grid.occupied[p] = true
				}
				break
			}
		}
	}

	fmt.Printf("Answer: %d", height)
}

func findCycle(heights []int, nr int) int {
	if len(heights) < 100*nr {
		return -1 // not enough elements to find cycle
	}
	for i := len(heights) - nr; i > nr; i-- {
		//fmt.Printf("%v\n", heights[len(heights)-nr:])
		if isCycle(heights[len(heights)-nr:], heights[i-nr:i]) {
			fmt.Printf("Cycle found at %v\n", i)
			//if isCycle(heights[len(heights)-nr:], heights[len(heights)-i*2:len(heights)-i]) {
			//	println("---------------")
			//	return len(heights) - i
			//}
		}
	}
	return -1
}

func isCycle(match []int, heights []int) bool {
	for i, m := range match {
		if heights[i] != m {
			return false
		}
	}
	return true
}

type Grid struct {
	occupied  map[image.Point]bool
	lastMoves [][]image.Point
}

//func (g *Grid) findCycle(length int) bool {
//out:
//	for j := len(g.lastMoves) - length; j > len(g.lastMoves)-100*length && j > 0; j-- {
//		//fmt.Printf("j = %+v\n", j)
//		for i, l := range g.lastMoves[len(g.lastMoves)-length:] {
//			for pi, p := range g.lastMoves[i] {
//				if g.lastMoves[j][pi].X != l.X {
//					continue out
//				}
//			}
//		}
//		fmt.Printf("ALL THE SAME AT %d", j)
//		return true
//	}
//	return false
//}

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
			g.lastMoves = append(g.lastMoves, rock)
			return rock, false // can't move
		}
		newPos = append(newPos, p)
	}
	return newPos, true
}

func (g *Grid) print(rock []image.Point) {
	//println("")
	//for y := 20; y >= 0; y-- {
	//	for x := 0; x < 7; x++ {
	//		if contains(rock, image.Pt(x, y)) {
	//			fmt.Printf("@")
	//		} else if g.occupied[image.Pt(x, y)] {
	//			fmt.Print("#")
	//		} else {
	//			fmt.Print(".")
	//		}
	//	}
	//	fmt.Print("\n")
	//}
}

func contains(pts []image.Point, p image.Point) bool {
	for _, point := range pts {
		if p == point {
			return true
		}
	}
	return false
}
