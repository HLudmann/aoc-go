package y2023

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func energies(contraption []string, start Pos, startDir rune) int {
	beamsPos := []Pos{start}
	maxX := len(contraption) - 1
	maxY := len(contraption[0]) - 1
	energy := make([][][]rune, maxX+1)
	for i := 0; i < len(energy); i++ {
		energy[i] = make([][]rune, maxY+1)
	}
	energy[start.x][start.y] = []rune{startDir}

	for len(beamsPos) > 0 {
		nextBeamsPos := []Pos{}
		for _, pos := range beamsPos {
			newPos := []Pos{}
			newBeams := []rune{}
			c := rune(contraption[pos.x][pos.y])
			for _, b := range energy[pos.x][pos.y] {
				if pos.x > 0 && (b == '^' && (c == '.' || c == '|') || b == '>' && (c == '/' || c == '|') || b == '<' && (c == '\\' || c == '|')) {
					newPos = append(newPos, Pos{pos.x - 1, pos.y})
					newBeams = append(newBeams, '^')
				}
				if pos.x < maxX && (b == 'v' && (c == '.' || c == '|') || b == '<' && (c == '/' || c == '|') || b == '>' && (c == '\\' || c == '|')) {
					newPos = append(newPos, Pos{pos.x + 1, pos.y})
					newBeams = append(newBeams, 'v')
				}
				if pos.y > 0 && (b == '<' && (c == '.' || c == '-') || b == 'v' && (c == '/' || c == '-') || b == '^' && (c == '\\' || c == '-')) {
					newPos = append(newPos, Pos{pos.x, pos.y - 1})
					newBeams = append(newBeams, '<')
				}
				if pos.y < maxY && (b == '>' && (c == '.' || c == '-') || b == '^' && (c == '/' || c == '-') || b == 'v' && (c == '\\' || c == '-')) {
					newPos = append(newPos, Pos{pos.x, pos.y + 1})
					newBeams = append(newBeams, '>')
				}

			}
			for i, b := range newBeams {
				p := newPos[i]
				if slices.Contains(energy[p.x][p.y], b) {
					continue
				}
				energy[p.x][p.y] = append(energy[p.x][p.y], b)
				nextBeamsPos = append(nextBeamsPos, p)
			}
		}
		beamsPos = nextBeamsPos
	}

	var count int
	for _, row := range energy {
		for _, nrg := range row {
			if len(nrg) > 0 {
				count++
			}
		}
	}
	return count
}

func day16Part1(input string) string {
	return fmt.Sprint(energies(toLines(input), Pos{0, 0}, '>'))
}

func day16Part2(input string) string {
	contraption := toLines(input)

	counts := make(chan int, 2*(len(contraption)+len(contraption[0])))

	for i := 0; i < len(contraption); i++ {
		x := i
		go func() {
			counts <- energies(contraption, Pos{x, 0}, '>')
		}()
		go func() {
			counts <- energies(contraption, Pos{x, len(contraption[0]) - 1}, '<')
		}()
	}

	for i := 0; i < len(contraption[0]); i++ {
		y := i
		go func() {
			counts <- energies(contraption, Pos{0, y}, 'v')
		}()
		go func() {
			counts <- energies(contraption, Pos{len(contraption) - 1, y}, '^')
		}()
	}

	maxCount := 0
	for i := 0; i < 2*(len(contraption)+len(contraption[0])); i++ {
		maxCount = max(maxCount, <-counts)
	}

	return fmt.Sprint(maxCount)
}

func Day16(test bool) {
	path := "inputs/2023/day16.txt"
	if test {
		path = strings.Replace(path, "day16", "day16-test", 1)
	}

	input, err := os.ReadFile(path)
	check(err)

	fmt.Println("Day 16")
	fmt.Println("\tPart 1:", day16Part1(string(input)))
	fmt.Println("\tPart 2:", day16Part2(string(input)))
}
