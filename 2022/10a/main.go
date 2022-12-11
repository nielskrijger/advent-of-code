package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("2022/10a/sample.txt")
	lines := strings.Split(string(data), "\n")

	x := 1
	var strength, cycle int

	tick := func() {
		cycle++
		if (cycle+20)%40 == 0 {
			strength += cycle * x
		}
	}

	for i := 0; i < len(lines); i++ {
		tick()
		if lines[i][:4] == "addx" {
			tick()
			num, _ := strconv.Atoi(lines[i][5:])
			x += num
		}
	}

	fmt.Printf("Answer: %d", strength)
}
