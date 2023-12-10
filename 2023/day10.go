package y2023

import (
	"fmt"
	"os"
	"strings"
)

type Tile struct {
	P Pos
	D Pos
}

type Field map[Pos]rune

func NewFieldAndS(input string) (Field, Pos) {
	f := make(Field)
	var sPos Pos
	for i, row := range toLines(input) {
		for j, r := range row {
			f[Pos{i, j}] = r
			if r == 'S' {
				sPos = Pos{i, j}
			}
		}
	}
	return f, sPos
}

func nextPipe(previous, current Pos, shape rune) Pos {
	switch shape {
	case '|', '-':
		return Pos{2*current.x - previous.x, 2*current.y - previous.y}
	case 'L', '7':
		return Pos{current.x + current.y - previous.y, current.y + current.x - previous.x}
	case 'J', 'F':
		return Pos{current.x + previous.y - current.y, current.y + previous.x - current.x}
	}

	return current
}

func startBranches(field Field, sPos Pos) (pipes []Pos) {

	pPos := Pos{sPos.x - 1, sPos.y}
	if r := field[pPos]; r == '7' || r == '|' || r == 'F' {
		pipes = append(pipes, pPos)
	}

	pPos = Pos{sPos.x + 1, sPos.y}
	if r := field[pPos]; r == 'J' || r == '|' || r == 'L' {
		pipes = append(pipes, pPos)
	}

	pPos = Pos{sPos.x, sPos.y - 1}
	if r := field[pPos]; r == 'F' || r == '-' || r == 'L' {
		pipes = append(pipes, pPos)
	}

	pPos = Pos{sPos.x, sPos.y + 1}
	if r := field[pPos]; r == '7' || r == '-' || r == 'J' {
		pipes = append(pipes, pPos)
	}

	if len(pipes) != 2 {
		panic(pipes)
	}

	return
}

func startShape(start, nextA, nextB Pos) rune {
	if nextA.x-start.x == 0 {
		return '-'
	}
	if nextB.y-start.y == 0 {
		return '|'
	}
	if nextA.x-start.x == -1 {
		if nextB.y-start.y == -1 {
			return 'J'
		}
		return 'L'
	}
	if nextB.y-start.y == -1 {
		return '7'
	}
	return 'F'
}

func day10Part1(input string) string {
	field, sPos := NewFieldAndS(input)

	starts := startBranches(field, sPos)

	dist := 1
	previouses := []Pos{sPos, sPos}
	for starts[0] != starts[1] {
		for i, pos := range starts {
			prev := previouses[i]
			shape := field[pos]
			next := nextPipe(prev, pos, shape)

			previouses[i] = pos
			starts[i] = next
		}
		dist++
	}

	return fmt.Sprint(dist)
}

func day10Part2(input string) string {
	field, sPos := NewFieldAndS(input)

	starts := startBranches(field, sPos)
	sShape := startShape(sPos, starts[0], starts[1])
	field[sPos] = sShape

	loop := map[Pos]bool{sPos: true}
	loop[starts[0]] = true
	loop[starts[1]] = true

	previouses := []Pos{sPos, sPos}
	for starts[0] != starts[1] {
		for i, pos := range starts {
			prev := previouses[i]
			shape := field[pos]
			next := nextPipe(prev, pos, shape)

			previouses[i] = pos
			starts[i] = next
			loop[next] = true
		}
	}

	var count int
	var inside bool
	var i, j int

	for field[Pos{i, j}] != 0 {
		for field[Pos{i, j}] != 0 {
			p := Pos{i, j}
			s := field[p]
			if !loop[p] && inside {
				count++
			}
			if loop[p] && s != '-' && s != '7' && s != 'F' {
				inside = !inside
			}
			j++
		}
		j = 0
		i++
	}

	return fmt.Sprint(count)
}

func Day10(test bool) {
	path := "inputs/2023/day10.txt"
	if test {
		path = strings.Replace(path, "day10", "day10-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day10Part1(string(input))
	p2 := day10Part2(string(input))

	fmt.Printf("Day 10\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
