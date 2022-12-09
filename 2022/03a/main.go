package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("2022/03a/sample.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)
	total := 0

	for sc.Scan() {
		txt := sc.Text()
		left := createSet(txt[:len(txt)/2])
		right := createSet(txt[len(txt)/2:])

		for k := range left {
			if right[k] {
				total += int(calculatePriority(k))
			}
		}
	}

	fmt.Printf("Answer: %d", total)
}

func calculatePriority(chr rune) int8 {
	if chr >= 'a' { // Lowercase ranges 97 to 122
		return int8(chr - 'a' + 1) // a = 1, z = 26
	} else { // Uppercase ranges 65 to 90
		return int8(chr - 'A' + 27) // A = 27, Z = 52
	}
}

func createSet(items string) map[rune]bool {
	set := make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return set
}
