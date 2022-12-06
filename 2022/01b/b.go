package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("2022/01b/input.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	maxCalories := make([]int, 4)
	currentCalories := 0

	for sc.Scan() {
		calories, err := strconv.Atoi(sc.Text())
		currentCalories += calories

		if err != nil {
			// Overwrite last element in the sorted array
			maxCalories[3] = currentCalories

			// Sort from maxCalories to lowest
			sort.Slice(maxCalories, func(i, j int) bool {
				return maxCalories[i] > maxCalories[j]
			})

			// Reset for next iteration
			currentCalories = 0
		}
	}

	var total int
	for _, v := range maxCalories[:3] {
		total += v
	}

	fmt.Printf("Answer: %d", total)
}
