package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	network := loadData("2022/16b/input.txt")

	candidates := network.candidates()
	max := 0
	for _, human := range dfs(network.find("AA"), candidates, 26) {
		elephantCandidates := make([]*Valve, 0)
		for _, candidate := range candidates {
			if _, ok := human[candidate.name]; !ok {
				elephantCandidates = append(elephantCandidates, candidate)
			}
		}

		// Takes a few minutes, should be optimized
		for _, elephant := range dfs(network.find("AA"), elephantCandidates, 26) {
			total := 0
			for _, v := range human {
				total += v
			}
			for _, v := range elephant {
				total += v
			}
			if total > max {
				max = total
			}
		}
	}
	fmt.Printf("Answer2: %d", max)
}

func dfs(valve *Valve, candidates []*Valve, t int) []map[string]int {
	var paths []map[string]int

	var _dfs func(valve *Valve, t int, visited map[string]int)
	_dfs = func(valve *Valve, t int, visited map[string]int) {
		if t <= 0 {
			return
		}

		for _, candidate := range candidates {
			if _, ok := visited[candidate.name]; ok {
				continue
			}

			visitedCopy := make(map[string]int)
			for k, v := range visited {
				visitedCopy[k] = v
			}

			time := t - valve.distanceTo(candidate) - 1
			if time <= 0 {
				continue
			}

			visitedCopy[candidate.name] = time * candidate.flowRate
			_dfs(candidate, time, visitedCopy)
		}

		paths = append(paths, visited)
	}

	_dfs(valve, t, map[string]int{})

	return paths
}

type Valve struct {
	name          string
	flowRate      int
	neighbours    []*Valve
	distanceCache map[*Valve]int
}

func (v *Valve) distanceTo(end *Valve) int { // Breadth first search algorithm
	if value, ok := v.distanceCache[end]; ok {
		return value
	}
	visited := make(map[*Valve]int)
	visited[v] = 0

	queue := []*Valve{v}

	result := -1
	for {
		if len(queue) == 0 {
			break
		}

		n := queue[0]
		queue = queue[1:]

		if n == end {
			result = visited[n]
			break
		}

		for _, neighbour := range n.neighbours {
			if _, ok := visited[neighbour]; !ok {
				visited[neighbour] = visited[n] + 1
				queue = append(queue, neighbour)
			}
		}
	}

	v.distanceCache[end] = result
	return result
}

func (v *Valve) String() string {
	result := v.name + "["
	for i, valve := range v.neighbours {
		if i > 0 {
			result += ","
		}
		result += valve.name
	}
	result += "]"
	return result
}

type Network struct {
	valves []*Valve
}

func (vs *Network) getOrCreateValve(name string) *Valve {
	for _, v := range vs.valves {
		if v.name == name {
			return v
		}
	}

	valve := &Valve{name: name, neighbours: []*Valve{}, distanceCache: map[*Valve]int{}}
	vs.valves = append(vs.valves, valve)
	return valve
}

func (vs *Network) find(name string) *Valve {
	for _, v := range vs.valves {
		if v.name == name {
			return v
		}
	}
	panic(fmt.Sprintf("Did not find %s", name))
}

// candidates returns all valves with flow rate > 0
func (vs *Network) candidates() []*Valve {
	valves := make([]*Valve, 0)
	for _, v := range vs.valves {
		if v.flowRate > 0 {
			valves = append(valves, v)
		}
	}
	return valves
}

func loadData(file string) *Network {
	data, _ := os.ReadFile(file)
	lines := strings.Split(string(data), "\n")
	network := &Network{}

	for _, line := range lines {
		result := regexp.MustCompile("Valve ([A-Z]+) has flow rate=([0-9]+); tunnels* leads* to valves* (.+)").FindStringSubmatch(line)
		valve := network.getOrCreateValve(result[1])
		flowRate, _ := strconv.Atoi(result[2])
		valve.flowRate = flowRate
		for _, neighbourName := range strings.Split(result[3], ", ") {
			valve.neighbours = append(valve.neighbours, network.getOrCreateValve(neighbourName))
		}
	}

	return network
}
