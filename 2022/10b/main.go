package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("2022/10a/input.txt")
	lines := strings.Split(string(data), "\n")

	cycle, x := 0, 1

	tick := func() {
		pos := cycle % 40
		cycle++
		if pos == 0 {
			println("")
		}
		if pos-x >= -1 && pos-x <= 1 {
			print("#")
		} else {
			print(".")
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
}
