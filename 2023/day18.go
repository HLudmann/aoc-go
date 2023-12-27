package y2023

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LavaTrenches struct {
	M                      map[Pos]int
	P                      Pos
	MinX, MaxX, MinY, MaxY int
}

func NewLavaTranches() *LavaTrenches {
	return &LavaTrenches{map[Pos]int{{0, 0}: 1}, Pos{0, 0}, 0, 0, 0, 0}
}

func (lt *LavaTrenches) Add(dir string, meters int) {
	var mv Pos
	switch dir {
	case "U":
		mv = Pos{-1, 0}
	case "D":
		mv = Pos{1, 0}
	case "L":
		mv = Pos{0, -1}
	default:
		mv = Pos{0, 1}
	}

	p := lt.P
	for i := 1; i <= meters; i++ {
		p = p.Add(mv)
		lt.M[p] = 1
	}
	lt.P = p
	lt.MaxX = max(lt.MaxX, p.x)
	lt.MaxY = max(lt.MaxY, p.y)
	lt.MinX = min(lt.MinX, p.x)
	lt.MinY = min(lt.MinY, p.y)
}

func (lt *LavaTrenches) PoolSize() int {
	var sum int
	for x := lt.MinX; x <= lt.MaxX; x++ {
		in := false
		for y := lt.MinY; y <= lt.MaxY; y++ {
			v := lt.M[Pos{x, y}]
			if v == 1 && lt.M[Pos{x + 1, y}] == 1 {
				in = !in
			}
			if v == 1 || in {
				sum++
			}
		}
	}
	return sum
}

func day18Part1(input string) string {
	lt := NewLavaTranches()
	for _, line := range toLines(input) {
		f := strings.Fields(line)
		lt.Add(f[0], toInt(f[1]))
	}

	return fmt.Sprint(lt.PoolSize())
}

func day18Part2(input string) string {
	lineParser := func(line string) (int, int) {
		hex := strings.Fields(line)[2]
		meters, _ := strconv.ParseInt(hex[2:7], 16, 64)
		return toInt(string(hex[7])), int(meters)
	}

	lines := toLines(input)
	var y, area int
	cw := 1
	dir, len := lineParser(lines[0])
	for _, line := range lines[1:] {
		ndir, nlen := lineParser(line)
		ncw := 0
		if ndir == (dir+1)&3 {
			ncw = 1
		}
		len += ncw + cw - 1
		if (dir & 1) == 1 {
			if dir < 2 {
				y += len
			} else {
				y -= len
			}
		} else {
			if dir > 1 {
				area += len * y
			} else {
				area -= len * y
			}
		}
		cw = ncw
		len = nlen
		dir = ndir
	}

	return fmt.Sprint(area)
}

func Day18(test bool) {
	path := "inputs/2023/day18.txt"
	if test {
		path = strings.Replace(path, "day18", "day18-test", 1)
	}

	input, err := os.ReadFile(path)
	check(err)

	fmt.Println("Day 18")
	fmt.Println("\tPart 1:", day18Part1(string(input)))
	fmt.Println("\tPart 2:", day18Part2(string(input)))
}
