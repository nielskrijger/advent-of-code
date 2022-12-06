package main

import (
	"bufio"
	"fmt"
	"os"
)

// Hardcoding is quicker than writing a parser I thought
//
//	    [C]             [L]         [T]
//	    [V] [R] [M]     [T]         [B]
//	    [F] [G] [H] [Q] [Q]         [H]
//	    [W] [L] [P] [V] [M] [V]     [F]
//	    [P] [C] [W] [S] [Z] [B] [S] [P]
//	[G] [R] [M] [B] [F] [J] [S] [Z] [D]
//	[J] [L] [P] [F] [C] [H] [F] [J] [C]
//	[Z] [Q] [F] [L] [G] [W] [H] [F] [M]
//	 1   2   3   4   5   6   7   8   9
var stacks = [][]string{
	{"Z", "J", "G"},
	{"Q", "L", "R", "P", "W", "F", "V", "C"},
	{"F", "P", "M", "C", "L", "G", "R"},
	{"L", "F", "B", "W", "P", "H", "M"},
	{"G", "C", "F", "S", "V", "Q"},
	{"W", "H", "J", "Z", "M", "Q", "T", "L"},
	{"H", "F", "S", "B", "V"},
	{"F", "J", "Z", "S"},
	{"M", "C", "D", "P", "F", "H", "B", "T"},
}

func main() {
	f, _ := os.Open("2022/05a/input.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		var amount, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &amount, &from, &to)

		for i := 0; i < amount; i++ {
			stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-1])
			stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1] // pop last element from array
		}
	}

	fmt.Print("Answer: ")
	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}
}
