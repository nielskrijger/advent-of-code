package main

import (
	"fmt"
	"os"
	"strings"
)

type Monkey struct {
	Name     string
	Number   int
	MonkeyA  string
	ValueA   int
	MonkeyB  string
	ValueB   int
	Operator string
}

var expressions = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func (m *Monkey) Value() int {
	if m.Number > 0 { // no monkey has number 0
		return m.Number
	}
	if m.ValueA > 0 && m.ValueB > 0 {
		m.Number = expressions[m.Operator](m.ValueA, m.ValueB)
	}
	return m.Number
}

func main() {
	monkeys, queue := loadData("2022/21a/sample.txt")
	for len(queue) > 0 {
		m := queue[0]
		queue = queue[1:]
		if m.ValueA == 0 && monkeys[m.MonkeyA].Value() > 0 {
			m.ValueA = monkeys[m.MonkeyA].Value()
		}
		if m.ValueB == 0 && monkeys[m.MonkeyB].Value() > 0 {
			m.ValueB = monkeys[m.MonkeyB].Value()
		}
		if m.ValueA == 0 || m.ValueB == 0 {
			queue = append(queue, m)
		} else if m.Name == "root" {
			fmt.Printf("Answer: %d", m.Value())
			break
		}
	}
}

func loadData(file string) (map[string]*Monkey, []*Monkey) {
	data, _ := os.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	result := make(map[string]*Monkey)
	mathMonkeys := make([]*Monkey, 0)
	for _, line := range lines {
		monkey := Monkey{}
		if len(line) < 10 {
			_, _ = fmt.Sscanf(line, "%4s: %d", &monkey.Name, &monkey.Number)
		} else {
			_, _ = fmt.Sscanf(line, "%4s: %s %s %s", &monkey.Name, &monkey.MonkeyA, &monkey.Operator, &monkey.MonkeyB)
			mathMonkeys = append(mathMonkeys, &monkey)
		}
		result[monkey.Name] = &monkey
	}
	return result, mathMonkeys
}
