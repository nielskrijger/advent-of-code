package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Rock = iota
	Paper
	Scissors

	Loss = 0
	Draw = 3
	Win  = 6
)

type round struct {
	opponent int
	me       int
}

// calculateTotalScore returns the total score of a single round which
// is the symbol score (Rock/Paper/Scissors + 1) + outcome (Win/Draw/Loss)
func (r round) calculateTotalScore() int {
	return r.me + 1 + r.calculateWinScore()
}

func (r round) calculateWinScore() int {
	if r.me == r.opponent {
		return Draw
	}
	if r.me == Rock && r.opponent == Paper {
		return Loss
	}
	if r.me == Paper && r.opponent == Scissors {
		return Loss
	}
	if r.me == Scissors && r.opponent == Rock {
		return Loss
	}

	return Win
}

func translateHandSymbol(symbol string) int {
	switch symbol {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	}

	panic(fmt.Sprintf("Unknown hand symbol: %q", symbol))
}

func determineMyHand(opponent int, outcome string) int {
	switch outcome {
	case "X": // Lose
		result := opponent - 1
		if result < 0 {
			return 2
		}
		return result
	case "Y": // Draw
		return opponent
	case "Z": // Win
		return (opponent + 1) % 3
	}

	panic(fmt.Sprintf("Unknown symbol: %q", outcome))
}

func main() {
	f, _ := os.Open("2022/02b/input.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	totalScore := 0

	for sc.Scan() {
		split := strings.Split(sc.Text(), " ")

		opponent := translateHandSymbol(split[0])

		r := round{
			opponent: opponent,
			me:       determineMyHand(opponent, split[1]),
		}

		totalScore += r.calculateTotalScore()
	}

	fmt.Printf("\nAnswer: %d", totalScore)
}
