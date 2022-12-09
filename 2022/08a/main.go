package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	matrix := loadData("2022/08a/input.txt")

	visible := 0
	for r := 1; r < len(matrix)-1; r++ {
		for c := 1; c < len(matrix)-1; c++ {
			col := getColumn(matrix, c)
			if isVisible(matrix[r][:c], matrix[r][c]) ||
				isVisible(matrix[r][c+1:], matrix[r][c]) ||
				isVisible(col[:r], matrix[r][c]) ||
				isVisible(col[r+1:], matrix[r][c]) {
				visible++
			}
		}
	}

	fmt.Printf("Answer: %d", len(matrix)*2+(len(matrix[0])*2-4)+visible)
}

func getColumn(matrix [][]int, col int) []int {
	result := make([]int, 0, len(matrix[0]))
	for row := 0; row < len(matrix[0]); row++ {
		result = append(result, matrix[row][col])
	}
	return result
}

func isVisible(row []int, value int) bool {
	for i := 0; i < len(row); i++ {
		if row[i] >= value {
			return false
		}
	}
	return true
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
