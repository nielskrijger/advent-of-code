package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	data, _ := os.ReadFile("2022/13b/sample.txt")

	two, six := &Elm{value: pointer(2)}, &Elm{value: pointer(6)}
	packets := []*Elm{two, six}
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			packets = append(packets, parseLine(line))
		}
	}

	sort.SliceStable(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})

	total := 1
	for i, packet := range packets {
		if packet == two || packet == six {
			total *= i + 1
		}
	}

	fmt.Printf("Answer: %d", total)
}

func pointer(i int) *int {
	return &(i)
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
