package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elm struct {
	Value int
	Index int
}

func main() {
	r, _ := loadData("2022/20a/sample.txt")
	for i := 0; i < r.Len(); i++ {
		for {
			e := r.Value.(Elm) // 2, 1, -3, 3, -2, 0, 4
			if e.Index == i {
				r = r.Prev()
				c := r.Unlink(1)
				r.Move(e.Value).Link(c)
				break
			}
			r = r.Next()
		}
	}

	for {
		if r = r.Next(); r.Value.(Elm).Value == 0 {
			break
		}
	}

	fmt.Printf("Answer: %v", r.Move(1000).Value.(Elm).Value+r.Move(2000).Value.(Elm).Value+r.Move(3000).Value.(Elm).Value)
}

func loadData(filename string) (*ring.Ring, []int) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	r := ring.New(len(lines))
	ids := make([]int, len(lines))
	for i, line := range lines {
		val, _ := strconv.Atoi(line)
		r.Value = Elm{Value: val, Index: i}
		r = r.Next()
		ids[i] = val
	}
	return r, ids
}
