package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("2022/01a/input.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	maxCalories := 0
	currentCalories := 0

	for sc.Scan() {
		calories, err := strconv.Atoi(sc.Text())
		currentCalories += calories

		if err != nil {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}

			currentCalories = 0
		}
	}

	fmt.Printf("Answer: %d", maxCalories)
}
