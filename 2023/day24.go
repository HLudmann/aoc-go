package y2023

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type HailStone struct {
	x, y, z, vx, vy, vz float64
}

func toHailStones(input string) []HailStone {
	var stones []HailStone
	re := regexp.MustCompile(`(-?\d+), (-?\d+), (-?\d+) @ (-?\d+), (-?\d+), (-?\d+)`)
	for _, line := range toLines(input) {
		grps := re.FindStringSubmatch(line)
		stones = append(stones, HailStone{
			toFloat(grps[1]),
			toFloat(grps[2]),
			toFloat(grps[3]),
			toFloat(grps[4]),
			toFloat(grps[5]),
			toFloat(grps[6]),
		})
	}
	return stones
}

func day24Part1(input string, test bool) string {
	mn, mx := 200000000000000.0, 400000000000000.0
	if test {
		mn, mx = 7, 27
	}
	stones := toHailStones(input)
	var count int
	for i, a := range stones[:len(stones)-1] {
		for _, b := range stones[i:] {
			if a.vx*b.vy-a.vy*b.vx == 0 {
				continue // colinear
			}
			x := (a.vx*b.vx*(a.y-b.y) + a.vx*b.vy*b.x - a.vy*b.vx*a.x) / (a.vx*b.vy - b.vx*a.vy)
			if x < mn || mx < x || (x-a.x)*a.vx < 0 || (x-b.x)*b.vx < 0 {
				continue
			}
			y := a.vy/a.vx*x + a.y - a.x*a.vy/a.vx
			if mn <= y && y <= mx {
				count++
			}
		}
	}
	return fmt.Sprint(count)
}

func day24Part2(input string, test bool) string {
	stones := toHailStones(input)

	matchingVelocity := func(l, vel int) map[int]bool {
		res := make(map[int]bool)
		for v := -1000; v <= 1000; v++ {
			if v != vel && l%(v-vel) == 0 {
				res[v] = true
			}
		}
		return res
	}

	intersectOrAll := func(old, new map[int]bool) map[int]bool {
		if len(old) == 0 {
			return new
		}
		for k := range new {
			if !old[k] {
				delete(new, k)
			}
		}
		return new
	}

	maybeX, maybeY, maybeZ := make(map[int]bool), make(map[int]bool), make(map[int]bool)
out:
	for i, a := range stones[:len(stones)-1] {
		for _, b := range stones[i+1:] {
			if len(maybeX) == 1 && len(maybeX) == len(maybeY) && len(maybeY) == len(maybeZ) {
				break out
			}
			if a.vx == b.vx {
				maybeX = intersectOrAll(maybeX, matchingVelocity(int(a.x-b.x), int(a.vx)))
			}
			if a.vy == b.vy {
				maybeY = intersectOrAll(maybeY, matchingVelocity(int(a.y-b.y), int(a.vy)))
			}
			if a.vz == b.vz {
				maybeZ = intersectOrAll(maybeZ, matchingVelocity(int(a.z-b.z), int(a.vz)))
			}
		}
	}

	rock := HailStone{}
	for k := range maybeX {
		rock.vx = float64(k)
	}
	for k := range maybeY {
		rock.vy = float64(k)
	}
	for k := range maybeZ {
		rock.vz = float64(k)
	}
	a := stones[0]
	b := stones[1]
	a.vx = a.vx - rock.vx
	a.vy = a.vy - rock.vy
	a.vz = a.vz - rock.vz
	b.vx = b.vx - rock.vx
	b.vy = b.vy - rock.vy
	x := (a.vx*b.vx*(a.y-b.y) + a.vx*b.vy*b.x - a.vy*b.vx*a.x) / (a.vx*b.vy - b.vx*a.vy)
	y := a.vy/a.vx*x + a.y - a.x*a.vy/a.vx
	t := (x - a.x) / a.vx
	z := a.vz*t + a.z

	return fmt.Sprint(int(x + y + z))
}

func Day24(test bool) {
	path := "inputs/2023/day24.txt"
	if test {
		path = strings.Replace(path, "day24", "day24-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day24Part1(string(input), test)
	p2 := day24Part2(string(input), test)

	fmt.Printf("Day 24\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
