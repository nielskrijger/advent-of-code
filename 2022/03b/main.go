package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("2022/03b/sample.txt")
	defer f.Close()

	total := 0

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		s1 := createSet(sc.Text())
		sc.Scan()
		s2 := createSet(sc.Text())
		sc.Scan()
		s3 := createSet(sc.Text())

		for k := range s1 {
			if s2[k] && s3[k] {
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
