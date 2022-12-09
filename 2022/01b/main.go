package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("2022/01b/sample.txt")
	groups := strings.Split(string(data), "\n\n")

	bags := make([]int, len(groups))

	for i, group := range groups {
		for _, calories := range strings.Fields(group) {
			cal, _ := strconv.Atoi(calories)
			bags[i] += cal
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(bags)))

	fmt.Printf("Answer: %d", bags[0]+bags[1]+bags[2])
}
