package main

import (
	"fmt"
	"os"
	"strings"
)

type ResourceIndex int

const (
	Ore ResourceIndex = iota
	Clay
	Obsidian
	Geode
)

type Resources [4]int // ore, clay, obsidian, geode

func (r Resources) Sub(sub Resources) Resources {
	for i, v := range sub {
		r[i] -= v
	}
	return r
}

func (r Resources) IsPositive() bool {
	for _, v := range r {
		if v < 0 {
			return false
		}
	}
	return true
}

func (r Resources) Add(robots Robots) Resources {
	for i, v := range robots {
		r[i] += v
	}
	return r
}

type Robots [4]int // for semantic purposes, same type as robots

func (r Robots) Add(resIndex ResourceIndex) Robots {
	arr := r
	arr[resIndex] += 1
	return arr
}

type Blueprint [4]Resources // ore robot cost, clay robot cost, obsidian robot cost, geode robot cost

func (b Blueprint) MaxCost(resource ResourceIndex) int {
	var max int
	for _, blueprint := range b {
		if blueprint[resource] > max {
			max = blueprint[resource]
		}
	}
	return max
}

func main() {
	total := 1
	for i, blueprint := range loadData("2022/19b/sample.txt") {
		geodes := dfs(0, blueprint, Resources{0, 0, 0, 0}, Robots{1, 0, 0, 0})
		fmt.Printf("Blueprint %d produces %d geodes\n", i+1, geodes)
		total *= geodes
	}
	fmt.Printf("Answer: %d", total)
}

// TODO performance can be improved a lot with memoization, but I didn't want to spend more time on it
func dfs(time int, blueprint Blueprint, wallet Resources, robots Robots) int {
	// Add funds and increase time until we can afford a robot or until max time was reached
	for {
		time++
		if time > 32 {
			return wallet[Geode]
		}

		ids := canBuyRobots(blueprint, wallet, robots)
		wallet = wallet.Add(robots) // Resource are increased after deciding to build anything

		if len(ids) > 0 {
			if ids[0] == Geode { // When it's a Geode robot always do that
				return dfs(time, blueprint, wallet.Sub(blueprint[Geode]), robots.Add(Geode))
			}

			var max int

			// Try spending nothing unless we can afford all non-Geode robots in which case
			// we should always buy something.
			if len(ids) < 3 {
				max = dfs(time, blueprint, wallet, robots)
			}

			// Try spending all possible robots and save the highest yielding Geodes value.
			for _, id := range ids {
				if geodes := dfs(time, blueprint, wallet.Sub(blueprint[id]), robots.Add(id)); geodes > max {
					max = geodes
				}
			}

			return max
		}
	}
}

func canBuyRobots(blueprint Blueprint, wallet Resources, robots Robots) []ResourceIndex {
	var ids []ResourceIndex
	for i := Geode; i >= Ore; i-- { // Prioritize buying a Geode over everything else

		// We can only buy 1 robot per tick, so if we already produce or have the max cost
		// of said resource it makes no sense to buy another robot of that resource type.
		if i != Geode && (robots[i] >= blueprint.MaxCost(i) || wallet[i] > blueprint.MaxCost(i)+1) {
			continue
		}

		if newRes := wallet.Sub(blueprint[i]); newRes.IsPositive() {
			ids = append(ids, i)
			if i == Geode { // no need to check further, we only use the Geode
				return ids
			}
		}
	}
	return ids
}

func loadData(filename string) []Blueprint {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	blueprints := make([]Blueprint, len(lines))
	for _, line := range lines {
		var id int
		var oreRobot, clayRobot, obsidianRobot, geodeRobot Resources
		_, _ = fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &id, &oreRobot[0], &clayRobot[0], &obsidianRobot[0], &obsidianRobot[1], &geodeRobot[0], &geodeRobot[2])
		blueprints[id-1] = [4]Resources{oreRobot, clayRobot, obsidianRobot, geodeRobot}
	}
	return blueprints
}
