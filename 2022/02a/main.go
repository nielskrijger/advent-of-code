package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Rock = iota + 1
	Paper
	Scissors

	Win  = 6
	Draw = 3
	Loss = 0
)

type round struct {
	opponent int
	me       int
}

// calculateTotalScore returns the total score of a single round which
// is the symbol score (Rock/Paper/Scissors) + outcome (Win/Draw/Loss)
func (r round) calculateTotalScore() int {
	return r.me + r.calculateWinScore()
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

func determineHand(symbol string) int {
	switch symbol {
	case "A":
		fallthrough
	case "X":
		return Rock
	case "B":
		fallthrough
	case "Y":
		return Paper
	case "C":
		fallthrough
	case "Z":
		return Scissors
	}

	panic(fmt.Sprintf("Unknown hand symbol: %q", symbol))
}

func main() {
	f, _ := os.Open("2022/02a/input.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	totalScore := 0

	for sc.Scan() {
		split := strings.Split(sc.Text(), " ")
		r := round{
			opponent: determineHand(split[0]),
			me:       determineHand(split[1]),
		}
		totalScore += r.calculateTotalScore()
	}

	fmt.Printf("Answer: %d", totalScore)
}
