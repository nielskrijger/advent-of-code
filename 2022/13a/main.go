package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	data, _ := os.ReadFile("2022/13a/sample.txt")
	groups := strings.Split(string(data), "\n\n")

	total := 0
	for i, group := range groups {
		lines := strings.Split(group, "\n")
		result := compare(parseLine(lines[0]), parseLine(lines[1]))
		if result < 0 {
			total += i + 1
		}
	}

	fmt.Printf("Answer: %d", total)
}

type Elm struct {
	value *int
	list  []*Elm
}

func parseLine(line string) *Elm {
	var s scanner.Scanner
	s.Init(strings.NewReader(line))
	current := &Elm{list: []*Elm{}}
	root, parent := current, current
	for s.Scan() != scanner.EOF {
		value, err := strconv.ParseInt(s.TokenText(), 10, 0)
		switch {
		case err == nil:
			val := int(value)
			current.list = append(current.list, &Elm{value: &val})
		case s.TokenText() == "[":
			parent = current
			newList := &Elm{list: []*Elm{}}
			current.list = append(current.list, newList)
			current = newList
		case s.TokenText() == "]":
			current = parent
		}
	}
	return root
}

func normalize(a, b *Elm) *Elm {
	if a.value != nil && b.value == nil { // if types don't match convert it to a list
		return &Elm{list: []*Elm{{value: a.value, list: []*Elm{}}}}
	}
	return a
}

func compare(leftElm, rightElm *Elm) int {
	left, right := normalize(leftElm, rightElm), normalize(rightElm, leftElm)

	if left.value != nil && right.value != nil {
		return *left.value - *right.value
	}

	for i, l := range left.list {
		if i >= len(right.list) {
			return 1 // right ran out of items
		}

		if result := compare(l, right.list[i]); result != 0 {
			return result
		}
	}

	return len(left.list) - len(right.list)
}
