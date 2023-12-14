package y2023

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type MirrorPlatform struct {
	Map        map[Pos]int
	RoundRocks []Pos
	Xlen, Ylen int
	Key        string
}

func NewMirrorPlatform(input string) *MirrorPlatform {
	lines := toLines(input)

	m := make(map[Pos]int)
	rocks := []Pos{}

	for x, line := range lines {
		for y, r := range line {
			p := Pos{x, y}
			switch r {
			case '.':
				m[p] = 0
			case 'O':
				m[p] = 1
				rocks = append(rocks, p)
			default:
				m[p] = 2
			}
		}
	}

	mp := MirrorPlatform{m, rocks, len(lines), len(lines[0]), ""}
	mp.computeKey()
	return &mp
}

func (mp *MirrorPlatform) Println() {
	for x := 0; x < mp.Xlen; x++ {
		line := ""
		for y := 0; y < mp.Ylen; y++ {
			p := Pos{x, y}

			switch mp.Map[p] {
			case 0:
				line += "."
			case 1:
				line += "O"
			case 2:
				line += "#"
			}
		}
		fmt.Println(line)
	}
}

func (mp *MirrorPlatform) computeKey() {
	rocks := mp.RoundRocks
	slices.SortStableFunc(rocks, func(a, b Pos) int { return (a.x-b.x)*100 + (a.y - b.y) })
	key := ""
	for _, r := range rocks {
		key += fmt.Sprintf("%d%d", r.x, r.y)
	}
	mp.Key = key
}

func (mp *MirrorPlatform) verticalTilt(dir int) {
	oldRocks := mp.RoundRocks
	slices.SortStableFunc(oldRocks, func(a, b Pos) int { return -dir * (a.x - b.x) })
	newRocks := make([]Pos, len(oldRocks))

	for i, p := range oldRocks {
		newX := p.x

		for mp.Map[Pos{newX + dir, p.y}] == 0 && (dir == -1 && 0 < newX || dir == 1 && newX < mp.Xlen-1) {
			newX += dir
		}
		newP := Pos{newX, p.y}
		mp.Map[p] = 0
		mp.Map[newP] = 1
		newRocks[i] = newP
	}

	mp.RoundRocks = newRocks
}

func (mp *MirrorPlatform) horizontalTilt(dir int) {
	oldRocks := mp.RoundRocks
	slices.SortStableFunc(oldRocks, func(a, b Pos) int { return -dir * (a.y - b.y) })
	newRocks := make([]Pos, len(oldRocks))

	for i, p := range oldRocks {
		newY := p.y

		for mp.Map[Pos{p.x, newY + dir}] == 0 && (dir == -1 && 0 < newY || dir == 1 && newY < mp.Ylen-1) {
			newY += dir
		}
		newP := Pos{p.x, newY}
		mp.Map[p] = 0
		mp.Map[newP] = 1
		newRocks[i] = newP
	}

	mp.RoundRocks = newRocks
}

func (mp *MirrorPlatform) TiltNorth() {
	mp.verticalTilt(-1)
}

func (mp *MirrorPlatform) TiltSouth() {
	mp.verticalTilt(1)
}

func (mp *MirrorPlatform) TiltWest() {
	mp.horizontalTilt(-1)
}

func (mp *MirrorPlatform) TiltEast() {
	mp.horizontalTilt(1)
}

func (mp *MirrorPlatform) oneSpin() {
	mp.TiltNorth()
	mp.TiltWest()
	mp.TiltSouth()
	mp.TiltEast()
}

func (mp *MirrorPlatform) Spin(times int) {
	memory := []string{}
	for i := 0; i < times; i++ {
		mp.oneSpin()

		mp.computeKey()
		if m := slices.Index(memory, mp.Key); m != -1 {
			rem := (times - m - 1) % (i - m)
			for j := 0; j < rem; j++ {
				mp.oneSpin()
			}
			break
		}
		memory = append(memory, mp.Key)
	}
}

func (mp *MirrorPlatform) NorthLoad() (sum int) {
	for _, pos := range mp.RoundRocks {
		sum += mp.Xlen - pos.x
	}
	return
}

func day14Part1(input string) string {
	platform := NewMirrorPlatform(input)
	platform.TiltNorth()
	return fmt.Sprint(platform.NorthLoad())
}

func day14Part2(input string) string {
	platform := NewMirrorPlatform(input)
	platform.Spin(1_000_000_000)
	return fmt.Sprint(platform.NorthLoad())
}

func Day14(test bool) {
	path := "inputs/2023/day14.txt"
	if test {
		path = strings.Replace(path, "day14", "day14-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day14Part1(string(input))
	p2 := day14Part2(string(input))

	fmt.Printf("Day 14\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
