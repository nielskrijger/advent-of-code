package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var hand = map[string]int{"A": 0, "B": 1, "C": 2}

func calculateWinScore(me int, opponent int) int {
	if me == opponent {
		return 3
	}
	if opponent == (me+1)%3 {
		return 0
	}
	return 6
}

func determineMyHand(opponent int, outcome string) int {
	result := opponent
	if outcome == "X" { // Lose
		result -= 1
	}
	if outcome == "Z" { // Win
		result += 1
	}
	return ((result % 3) + 3) % 3 // apply mod twice to yield positive result for negative modulo (python-like)
}

func main() {
	f, _ := os.Open("2022/02b/input.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	totalScore := 0

	for sc.Scan() {
		split := strings.Split(sc.Text(), " ")

		opponent := hand[split[0]]
		me := determineMyHand(opponent, split[1])
		totalScore += me + 1 + calculateWinScore(me, opponent)
	}

	fmt.Printf("Answer: %d", totalScore)
}
