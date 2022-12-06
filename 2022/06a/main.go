package main

import (
	"fmt"
	"os"
)

const uniqueChars = 4

func main() {
	input, _ := os.ReadFile("2022/06b/input.txt")

	for i := 0; i < len(input); i++ {
		seen := make(map[byte]bool)

		for _, r := range input[i : i+uniqueChars] {
			seen[r] = true
		}

		if len(seen) == uniqueChars {
			fmt.Printf("Answer: %d", i+uniqueChars)
			break
		}
	}
}
