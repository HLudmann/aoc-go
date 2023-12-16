package y2023

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Lens struct {
	Label string
	Focal int
}

func HolidayHash(input string) int {
	hash := 0
	for _, c := range input {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

func day15Part1(input string) string {
	var sum int
	for _, step := range strings.Split(toLines(input)[0], ",") {
		sum += HolidayHash(step)
	}

	return fmt.Sprint(sum)
}

func day15Part2(input string) string {
	boxes := make([][]Lens, 256)
	for _, step := range strings.Split(toLines(input)[0], ",") {
		if step[len(step)-1] == '-' {
			label := step[:len(step)-1]
			hash := HolidayHash(label)
			index := slices.IndexFunc(boxes[hash], func(l Lens) bool { return l.Label == label })
			if index == -1 {
				continue
			}
			boxes[hash] = slices.Delete(boxes[hash], index, index+1)
		} else {
			labelAndFocal := strings.Split(step, "=")
			lens := Lens{labelAndFocal[0], toInt(labelAndFocal[1])}
			hash := HolidayHash(lens.Label)
			if index := slices.IndexFunc(boxes[hash], func(l Lens) bool { return l.Label == lens.Label }); index != -1 {
				boxes[hash][index] = lens
			} else {
				boxes[hash] = append(boxes[hash], lens)
			}
		}
	}

	var sum int
	for i, box := range boxes {
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * lens.Focal
		}
	}

	return fmt.Sprint(sum)
}

func Day15(test bool) {
	path := "inputs/2023/day15.txt"
	if test {
		path = strings.Replace(path, "day15", "day15-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day15Part1(string(input))
	p2 := day15Part2(string(input))

	fmt.Printf("Day 15\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
