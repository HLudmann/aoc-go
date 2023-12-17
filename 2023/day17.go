package y2023

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type HeatLossGraph map[Pos]int

type CityBlock struct {
	P    Pos
	Dirs string
}

type CityBlockQueue struct {
	Blocks []CityBlock
	Size   int
}

func (cbq *CityBlockQueue) Push(new CityBlock, dist map[CityBlock]int) {
	if cbq.Size == 0 {
		cbq.Blocks = []CityBlock{new}
		cbq.Size = 1
		return
	}
	index := slices.IndexFunc(cbq.Blocks, func(cb CityBlock) bool { return dist[cb] < dist[new] })
	if index == -1 {
		cbq.Blocks = append(cbq.Blocks, new)
	} else {
		cbq.Blocks = slices.Insert(cbq.Blocks, index, new)
	}
	cbq.Size++
}

func (cbq *CityBlockQueue) Pop() CityBlock {
	cb := cbq.Blocks[cbq.Size-1]
	cbq.Blocks = cbq.Blocks[:cbq.Size-1]
	cbq.Size--

	return cb
}

func (hlm HeatLossGraph) Neighbours(node CityBlock, minF, maxF int) (neigh []CityBlock) {
	d := "."
	if l := len(node.Dirs); l > 0 {
		d = string(node.Dirs[l-1])
	}

	if d == "." || minF <= len(node.Dirs) {
		if p := (Pos{node.P.x - 1, node.P.y}); hlm[p] != 0 && (d == "." || d == "<" || d == ">") {
			neigh = append(neigh, CityBlock{p, "^"})
		}
		if p := (Pos{node.P.x + 1, node.P.y}); hlm[p] != 0 && (d == "." || d == "<" || d == ">") {
			neigh = append(neigh, CityBlock{p, "v"})
		}
		if p := (Pos{node.P.x, node.P.y - 1}); hlm[p] != 0 && (d == "." || d == "^" || d == "v") {
			neigh = append(neigh, CityBlock{p, "<"})
		}
		if p := (Pos{node.P.x, node.P.y + 1}); hlm[p] != 0 && (d == "." || d == "^" || d == "v") {
			neigh = append(neigh, CityBlock{p, ">"})
		}
	}

	if d == "." {
		return
	}

	if len(node.Dirs) < maxF {
		var p Pos
		switch d {
		case "^":
			p = Pos{node.P.x - 1, node.P.y}
		case "v":
			p = Pos{node.P.x + 1, node.P.y}
		case "<":
			p = Pos{node.P.x, node.P.y - 1}
		case ">":
			p = Pos{node.P.x, node.P.y + 1}

		}
		if hlm[p] != 0 {
			neigh = append(neigh, CityBlock{p, node.Dirs + d})
		}
	}
	return
}

func findPath(graph HeatLossGraph, start, goal Pos, minF, maxF int) int {
	startBlock := CityBlock{start, ""}
	dist := make(map[CityBlock]int)
	prev := make(map[CityBlock]CityBlock)
	dist[startBlock] = 0

	queue := &CityBlockQueue{[]CityBlock{}, 0}
	for _, n := range graph.Neighbours(startBlock, minF, maxF) {
		dist[n] = graph[n.P]
		prev[n] = startBlock
		queue.Push(n, dist)
	}

	cb := startBlock
	for queue.Size != 0 {
		cb = queue.Pop()
		if minF <= len(cb.Dirs) && cb.P == goal {
			return dist[cb]
		}
		neigh := graph.Neighbours(cb, minF, maxF)
		for _, n := range neigh {
			if d := dist[cb] + graph[n.P]; dist[n] == 0 || d < dist[n] {
				dist[n] = d
				queue.Push(n, dist)
				prev[n] = cb
			}
		}
	}

	return 0
}

func day17Part1(input string) string {
	graph := make(HeatLossGraph)
	for i, line := range toLines(input) {
		for j, val := range line {
			graph[Pos{i, j}] = toInt(string(val))
		}
	}
	maxX := int(math.Sqrt(float64(len(graph)))) - 1
	return fmt.Sprint(findPath(graph, Pos{0, 0}, Pos{maxX, maxX}, 1, 3))
}

func day17Part2(input string) string {
	graph := make(HeatLossGraph)
	for i, line := range toLines(input) {
		for j, val := range line {
			graph[Pos{i, j}] = toInt(string(val))
		}
	}
	maxX := int(math.Sqrt(float64(len(graph)))) - 1
	return fmt.Sprint(findPath(graph, Pos{0, 0}, Pos{maxX, maxX}, 4, 10))
}

func Day17(test bool) {
	path := "inputs/2023/day17.txt"
	if test {
		path = strings.Replace(path, "day17", "day17-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day17Part1(string(input))
	p2 := day17Part2(string(input))

	fmt.Printf("Day 17\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
