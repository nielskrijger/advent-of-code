package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items       []int
	operator    string
	operand     string // can be "old" or int
	test        int
	monkeyTrue  int
	monkeyFalse int
	inspections int
}

func (m *monkey) inspect(i int) int {
	m.inspections++

	operand := 0
	if m.operand == "old" {
		operand = m.items[i]
	} else {
		operand = mustAtoi(m.operand)
	}

	newWorry := m.items[i]
	if m.operator == "+" { // Either + or *
		return m.items[i] + operand
	}
	return newWorry * operand
}

func main() {
	monkeys := loadData()

	lcm := leastCommonMultiple(monkeys)

	for round := 0; round < 10000; round++ {
		for i := 0; i < len(monkeys); i++ {
			for j := 0; j < len(monkeys[i].items); j++ {
				newWorry := monkeys[i].inspect(j)

				newWorry %= lcm

				newMonkey := monkeys[i].monkeyFalse
				if newWorry%monkeys[i].test == 0 {
					newMonkey = monkeys[i].monkeyTrue
				}
				monkeys[newMonkey].items = append(monkeys[newMonkey].items, newWorry)
			}
			monkeys[i].items = []int{}
		}
	}

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})
	fmt.Printf("Answer: %v", monkeys[0].inspections*monkeys[1].inspections)
}

func leastCommonMultiple(monkeys []*monkey) int {
	result := 1
	for _, monkey := range monkeys {
		result *= monkey.test
	}
	return result
}

func loadData() []*monkey {
	data, _ := os.ReadFile("2022/11b/sample.txt")
	inputs := strings.Split(string(data), "\n\n")

	monkeys := make([]*monkey, len(inputs), len(inputs))
	for i, input := range inputs {
		lines := strings.Split(input, "\n")
		monkeys[i] = &monkey{ // Quick & dirty
			items:       sliceAtoi(strings.Split(lines[1][18:], ", ")),
			operator:    lines[2][23:24],
			operand:     lines[2][25:],
			test:        mustAtoi(lines[3][21:]),
			monkeyTrue:  mustAtoi(lines[4][29:]),
			monkeyFalse: mustAtoi(lines[5][30:]),
		}
	}

	return monkeys
}

func sliceAtoi(sa []string) []int {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, _ := strconv.Atoi(a)
		si = append(si, i)
	}
	return si
}

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
