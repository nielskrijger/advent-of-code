package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("2022/04a/input.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	contained := 0

	for sc.Scan() {
		var a1, a2, b1, b2 int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &a1, &a2, &b1, &b2)

		if (a1 >= b1 && a2 <= b2) || (b1 >= a1 && b2 <= a2) {
			contained += 1
		}
	}

	fmt.Printf("Answer: %d", contained)
}
