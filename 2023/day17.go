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
	P   Pos
	Dir rune
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

func (hlm HeatLossGraph) Neighbours(node CityBlock, dist, minF, maxF int) map[CityBlock]int {
	neigh := make(map[CityBlock]int)
	var side Pos
	var plus, minus rune
	switch node.Dir {
	case '^', 'v':
		side = Pos{0, 1}
		minus = '<'
		plus = '>'
	case '<', '>':
		side = Pos{1, 0}
		minus = '^'
		plus = 'v'
	}
	dPlus := dist
	dMinus := dist
	for i := 1; i <= maxF; i++ {
		pPlus := node.P.Add(side.Multiply(i))
		pMinus := node.P.Add(side.Multiply(-i))
		if d := hlm[pPlus]; d != 0 {
			dPlus += d
			if minF <= i {
				neigh[CityBlock{pPlus, plus}] = dPlus
			}
		}
		if d := hlm[pMinus]; d != 0 {
			dMinus += d
			if minF <= i {
				neigh[CityBlock{pMinus, minus}] = dMinus
			}
		}
	}

	return neigh
}

func findPath(graph HeatLossGraph, minF, maxF int) int {
	dist := make(map[CityBlock]int)
	queue := &CityBlockQueue{[]CityBlock{}, 0}
	for n, d := range graph.Neighbours(CityBlock{Pos{0, 0}, '>'}, 0, minF, maxF) {
		dist[n] = d
		queue.Push(n, dist)
	}
	for n, d := range graph.Neighbours(CityBlock{Pos{0, 0}, 'v'}, 0, minF, maxF) {
		dist[n] = d
		queue.Push(n, dist)
	}

	maxX := int(math.Sqrt(float64(len(graph)))) - 1
	goal := Pos{maxX, maxX}
	for queue.Size != 0 {
		cb := queue.Pop()
		if cb.P == goal {
			return dist[cb]
		}
		neigh := graph.Neighbours(cb, dist[cb], minF, maxF)
		for n, d := range neigh {
			if dist[n] == 0 || d < dist[n] {
				dist[n] = d
				queue.Push(n, dist)
			}
		}
	}

	return -1
}

func day17Part1(input string) string {
	graph := make(HeatLossGraph)
	for i, line := range toLines(input) {
		for j, val := range line {
			graph[Pos{i, j}] = toInt(string(val))
		}
	}
	return fmt.Sprint(findPath(graph, 1, 3))
}

func day17Part2(input string) string {
	graph := make(HeatLossGraph)
	for i, line := range toLines(input) {
		for j, val := range line {
			graph[Pos{i, j}] = toInt(string(val))
		}
	}
	return fmt.Sprint(findPath(graph, 4, 10))
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
