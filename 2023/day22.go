package y2023

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pos3D struct{ x, y, z int }

type SandBlock struct {
	x, y, z       int
	size          int
	dir           string
	OnTop, Bellow map[*SandBlock]bool
}

func (sb SandBlock) Range() []Pos3D {
	var r []Pos3D
	for i := 0; i < sb.size; i++ {
		switch sb.dir {
		case "x":
			r = append(r, Pos3D{sb.x + i, sb.y, sb.z})
		case "y":
			r = append(r, Pos3D{sb.x, sb.y + i, sb.z})
		default:
			r = append(r, Pos3D{sb.x, sb.y, sb.z + i})
		}
	}
	return r
}

func toJentris(input string) []*SandBlock {
	var blocks []*SandBlock
	bMap := make(map[Pos3D]*SandBlock)
	for _, line := range toLines(input) {
		s := strings.Split(line, "~")
		s1 := strings.Split(s[0], ",")
		x1, y1, z1 := toInt(s1[0]), toInt(s1[1]), toInt(s1[2])
		s2 := strings.Split(s[1], ",")
		xDiff, yDiff, zDiff := toInt(s2[0])-x1, toInt(s2[1])-y1, toInt(s2[2])-z1
		var dir string
		size := 1
		switch {
		case xDiff != 0:
			dir = "x"
			size = xDiff + 1
		case yDiff != 0:
			dir = "y"
			size = yDiff + 1
		case zDiff != 0:
			dir = "z"
			size = zDiff + 1
		}
		sb := SandBlock{x1, y1, z1, size, dir, make(map[*SandBlock]bool), make(map[*SandBlock]bool)}
		blocks = append(blocks, &sb)
		for i := 0; i < xDiff+1; i++ {
			for j := 0; j < yDiff+1; j++ {
				for k := 0; k < zDiff+1; k++ {
					bMap[Pos3D{x1 + i, y1 + j, z1 + k}] = &sb
				}
			}
		}
	}
	slices.SortStableFunc(blocks, func(a, b *SandBlock) int {
		if z := a.z - b.z; z != 0 {
			return z
		}
		if y := a.y - b.y; y != 0 {
			return y
		}
		return a.x - b.x
	})

	for _, b := range blocks {
		if b.z == 1 {
			continue
		}
		i := 0
		cont := true
		for cont {
			if b.z-i-1 < 1 {
				break
			}
			for _, p := range b.Range() {
				new := p
				new.z -= (i + 1)
				if b2 := bMap[new]; b2 != nil && b2 != b {
					cont = false
					b.Bellow[b2] = true
					b2.OnTop[b] = true
				}
			}
			if cont {
				i++
			}
		}
		b.z = b.z - i
		for _, p := range b.Range() {
			old := p
			old.z += i
			bMap[old] = nil
			bMap[p] = b
		}
	}
	return blocks
}

func day22Part1(input string) string {
	blocks := toJentris(input)
	var count int
	for _, b := range blocks {
		destroyable := true
		for b2 := range b.OnTop {
			if b2.dir == "z" || b2.dir == "" || len(b2.Bellow) == 1 {
				destroyable = false
				break
			}
		}
		if destroyable {
			count++
		}
	}
	return fmt.Sprint(count)
}

func day22Part2(input string) string {
	blocks := toJentris(input)
	var count int
	for _, b := range blocks {
		fell := make(map[*SandBlock]bool)
		queue := b.OnTop
		for len(queue) != 0 {
			nq := make(map[*SandBlock]bool)
			for b2 := range queue {
				f := true
				for b3 := range b2.Bellow {
					if b3 != b && !fell[b3] {
						f = false
						break
					}
				}
				if f {
					fell[b2] = true
					for b3 := range b2.OnTop {
						nq[b3] = true
					}
				}
			}
			queue = nq
		}
		count += len(fell)
	}
	return fmt.Sprint(count)
}

func Day22(test bool) {
	path := "inputs/2023/day22.txt"
	if test {
		path = strings.Replace(path, "day22", "day22-test", 1)
	}

	input, err := os.ReadFile(path)
	check(err)

	fmt.Println("Day 22")
	fmt.Println("\tPart 1:", day22Part1(string(input)))
	fmt.Println("\tPart 2:", day22Part2(string(input)))
}
