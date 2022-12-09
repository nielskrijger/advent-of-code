package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	matrix := loadData("2022/08b/input.txt")

	highest := 0
	for r := 1; r < len(matrix)-1; r++ {
		for c := 1; c < len(matrix)-1; c++ {
			col := getColumn(matrix, c)

			scenicScore := numVisible(reverse(matrix[r][:c]), matrix[r][c]) *
				numVisible(matrix[r][c+1:], matrix[r][c]) *
				numVisible(reverse(col[:r]), matrix[r][c]) *
				numVisible(col[r+1:], matrix[r][c])

			if scenicScore > highest {
				highest = scenicScore
			}
		}
	}

	fmt.Printf("Answer: %d", highest)
}

func getColumn(matrix [][]int, col int) []int {
	result := make([]int, 0, len(matrix[0]))
	for row := 0; row < len(matrix[0]); row++ {
		result = append(result, matrix[row][col])
	}
	return result
}

func numVisible(line []int, value int) int {
	score := 0
	for i := 0; i < len(line); i++ {
		score++

		if line[i] >= value {
			break
		}
	}
	return score
}

func reverse(s []int) []int {
	result := make([]int, 0, len(s))
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		result = append(result, s[j])
	}
	return result
}

func loadData(filename string) [][]int {
	f, _ := os.Open(filename)
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanRunes)
	matrix := make([][]int, 0)

	row := 0
	for sc.Scan() {
		if sc.Text() == "\n" {
			row++
			continue
		}

		if len(matrix) <= row {
			r1 := make([]int, 0)
			matrix = append(matrix, r1)
		}

		height, _ := strconv.Atoi(sc.Text())
		matrix[row] = append(matrix[row], height)
	}

	return matrix
}
