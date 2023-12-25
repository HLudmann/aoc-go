package y2023

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Neighbour struct {
	p Pos
	w int
}

func trailNeighbours(trailMat map[Pos]rune, visited map[Pos]bool, p Pos, part1 bool) (neighs []Pos) {
	val := trailMat[p]
	var shifts []Pos
	switch {
	case part1 && val == '<':
		shifts = []Pos{{0, -1}}
	case part1 && val == '<':
		shifts = []Pos{{0, 1}}
	case part1 && val == '<':
		shifts = []Pos{{-1, 0}}
	case part1 && val == '<':
		shifts = []Pos{{1, 0}}
	default:
		shifts = []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	}
	for _, s := range shifts {
		if part1 {
			switch {
			case s.x == 1 && trailMat[p.Add(s)] == '^':
				continue
			case s.x == -1 && trailMat[p.Add(s)] == 'v':
				continue
			case s.y == 1 && trailMat[p.Add(s)] == '<':
				continue
			case s.y == -1 && trailMat[p.Add(s)] == '>':
				continue
			}
		}
		if n := p.Add(s); trailMat[n] != 0 && !visited[n] {
			neighs = append(neighs, n)
		}
	}

	return
}

func toTrailMap(input string, part1 bool) (map[Pos]map[Neighbour]bool, Pos, Pos) {
	mat := map[Pos]rune{}
	var size int
	for i, line := range toLines(input) {
		size++
		for j, val := range line {
			if val != '#' {
				mat[Pos{i, j}] = val
			}
		}
	}
	start := Pos{0, 1}
	stop := Pos{size - 1, size - 2}

	graph := map[Pos]map[Neighbour]bool{}
	queue := [][2]interface{}{{start, Neighbour{start, 0}}}

	for len(queue) != 0 {
		visited := make(map[Pos]bool)
		s := queue[0][0].(Pos)
		cur := queue[0][1].(Neighbour)
		queue = queue[1:]
		if graph[s] == nil {
			graph[s] = make(map[Neighbour]bool)
		}
		neighs := slices.DeleteFunc(trailNeighbours(mat, visited, cur.p, part1), func(p Pos) bool { return graph[p] != nil })
		for len(neighs) == 1 {
			n := neighs[0]
			if graph[n] != nil {
				break
			}
			visited[cur.p] = true
			cur.p = n
			cur.w++
			neighs = trailNeighbours(mat, visited, cur.p, part1)
		}
		graph[s][cur] = true
		if !part1 {
			if graph[cur.p] == nil {
				graph[cur.p] = make(map[Neighbour]bool)
			}
			graph[cur.p][Neighbour{s, cur.w}] = true
		}
		for _, p := range neighs {
			if graph[p] == nil {
				graph[p] = make(map[Neighbour]bool)
			}
			queue = append(queue, [2]interface{}{cur.p, Neighbour{p, 1}})
		}
	}

	return graph, start, stop
}

func dfs23(graph map[Pos]map[Neighbour]bool, start, stop Pos, length int, visited map[Pos]bool, path []Pos) int {
	if start == stop {
		return length
	}
	best := 0
	for n := range graph[start] {
		if visited[n.p] {
			continue
		}
		visited[n.p] = true
		best = max(best, dfs23(graph, n.p, stop, length+n.w, visited, append(path, n.p)))
		delete(visited, n.p)
	}
	return best
}

func day23Part1(input string) string {
	graph, start, stop := toTrailMap(input, true)
	return fmt.Sprint(dfs23(graph, start, stop, 0, make(map[Pos]bool), []Pos{start}))
}

func day23Part2(input string) string {
	graph, start, stop := toTrailMap(input, false)
	fmt.Println(start, "->", stop)
	return fmt.Sprint(dfs23(graph, start, stop, 0, make(map[Pos]bool), []Pos{start}))
}

func Day23(test bool) {
	path := "inputs/2023/day23.txt"
	if test {
		path = strings.Replace(path, "day23", "day23-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day23Part1(string(input))
	p2 := day23Part2(string(input))

	fmt.Printf("Day 23\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
