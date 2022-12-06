package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var hand = map[string]int{"A": 0, "X": 0, "B": 1, "Y": 1, "C": 2, "Z": 2}

func calculateWinScore(me int, opponent int) int {
	if me == opponent {
		return 3
	}
	if opponent == (me+1)%3 {
		return 0
	}
	return 6
}

func main() {
	f, _ := os.Open("2022/02a/input.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	totalScore := 0

	for sc.Scan() {
		split := strings.Split(sc.Text(), " ")
		totalScore += hand[split[1]] + 1 + calculateWinScore(hand[split[1]], hand[split[0]])
	}

	fmt.Printf("Answer: %d", totalScore)
}
