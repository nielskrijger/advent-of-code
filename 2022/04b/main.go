package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("2022/04b/input.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	overlapping := 0

	for sc.Scan() {
		var a1, a2, b1, b2 int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &a1, &a2, &b1, &b2)

		if b1 <= a2 && b2 >= a1 {
			overlapping += 1
		}
	}

	fmt.Printf("Answer: %d", overlapping)
}
