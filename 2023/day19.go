package y2023

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type MachinePart struct {
	X, M, A, S int
}

func (mp MachinePart) Value() int {
	return mp.X + mp.M + mp.A + mp.S
}

func parseMachineParts(input string) (parts []MachinePart) {
	re := regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
	for _, line := range toLines(input) {
		grps := re.FindStringSubmatch(line)
		parts = append(parts, MachinePart{toInt(grps[1]), toInt(grps[2]), toInt(grps[3]), toInt(grps[4])})
	}

	return
}

func parseTestsP1(input string) map[string]func(mp MachinePart) string {
	funcs := make(map[string]func(mp MachinePart) string)
	re := regexp.MustCompile(`(\w+){(.*),(\w+)}`)
	testRe := regexp.MustCompile(`(\w)([<>])(\d+):(\w+)`)
	for _, line := range toLines(input) {
		grps := re.FindStringSubmatch(line)
		key := grps[1]
		lastReturn := grps[3]
		tests := strings.Split(grps[2], ",")

		f := func(mp MachinePart) string {
			for _, test := range tests {
				g := testRe.FindStringSubmatch(test)
				var v int
				c := toInt(g[3])
				switch g[1] {
				case "x":
					v = mp.X
				case "m":
					v = mp.M
				case "a":
					v = mp.A
				default:
					v = mp.S
				}
				if g[2] == "<" && v < c || g[2] == ">" && v > c {
					return g[4]
				}
			}

			return lastReturn
		}

		funcs[key] = f
	}
	return funcs
}

func inputToTestsAndPartsP1(input string) (map[string]func(mp MachinePart) string, []MachinePart) {
	testsAndParts := strings.Split(input, "\n\n")
	return parseTestsP1(testsAndParts[0]), parseMachineParts(testsAndParts[1])
}

func testPart(tests map[string]func(mp MachinePart) string, part MachinePart) bool {
	res := "in"
	for res != "A" && res != "R" {
		res = tests[res](part)
	}
	return res == "A"

}

func day19Part1(input string) string {
	tests, parts := inputToTestsAndPartsP1(input)

	var sum int
	for _, part := range parts {
		if testPart(tests, part) {
			sum += part.Value()
		}
	}
	return fmt.Sprint(sum)
}

type MachinePartRanges struct {
	MinX, MaxX, MinM, MaxM, MinA, MaxA, MinS, MaxS int
	Next                                           string
}

func (mpr MachinePartRanges) Possibilities() int {
	return (mpr.MaxX - mpr.MinX + 1) * (mpr.MaxM - mpr.MinM + 1) * (mpr.MaxA - mpr.MinA + 1) * (mpr.MaxS - mpr.MinS + 1)
}

func NewMachinePartRanges() MachinePartRanges {
	return MachinePartRanges{1, 4000, 1, 4000, 1, 4000, 1, 4000, "in"}
}

func SplitMachinePartRanges(mpr MachinePartRanges, k, d string, val int) (MachinePartRanges, MachinePartRanges) {
	switch k + d {
	case "x>":
		cp := mpr
		mpr.MaxX = min(val, mpr.MaxX)
		cp.MinX = max(val+1, mpr.MinX)
		return mpr, cp
	case "x<":
		cp := mpr
		cp.MaxX = min(val-1, mpr.MaxX)
		mpr.MinX = max(val, mpr.MinX)
		return mpr, cp
	case "m>":
		cp := mpr
		mpr.MaxM = min(val, mpr.MaxM)
		cp.MinM = max(val+1, mpr.MinM)
		return mpr, cp
	case "m<":
		cp := mpr
		cp.MaxM = min(val-1, mpr.MaxM)
		mpr.MinM = max(val, mpr.MinM)
		return mpr, cp
	case "a>":
		cp := mpr
		mpr.MaxA = min(val, mpr.MaxA)
		cp.MinA = max(val+1, mpr.MinA)
		return mpr, cp
	case "a<":
		cp := mpr
		cp.MaxA = min(val-1, mpr.MaxA)
		mpr.MinA = max(val, mpr.MinA)
		return mpr, cp
	case "s>":
		cp := mpr
		mpr.MaxS = min(val, mpr.MaxS)
		cp.MinS = max(val+1, mpr.MinS)
		return mpr, cp
	default:
		cp := mpr
		cp.MaxS = min(val-1, mpr.MaxS)
		mpr.MinS = max(val, mpr.MinS)
		return mpr, cp
	}
}

func parseTestsP2(input string) map[string]func(mpr MachinePartRanges) []MachinePartRanges {
	funcs := make(map[string]func(mp MachinePartRanges) []MachinePartRanges)
	re := regexp.MustCompile(`(\w+){(.*),(\w+)}`)
	testRe := regexp.MustCompile(`(\w)([<>])(\d+):(\w+)`)
	for _, line := range toLines(input) {
		grps := re.FindStringSubmatch(line)
		key := grps[1]
		lastReturn := grps[3]
		tests := strings.Split(grps[2], ",")

		f := func(mpr MachinePartRanges) []MachinePartRanges {
			ranges := []MachinePartRanges{}
			var new MachinePartRanges
			for _, test := range tests {
				g := testRe.FindStringSubmatch(test)
				val := toInt(g[3])
				mpr, new = SplitMachinePartRanges(mpr, g[1], g[2], val)
				new.Next = g[4]
				ranges = append(ranges, new)
			}

			mpr.Next = lastReturn
			ranges = append(ranges, mpr)
			return ranges
		}

		funcs[key] = f
	}
	return funcs
}

func day19Part2(input string) string {
	tests := parseTestsP2(strings.Split(input, "\n\n")[0])

	ranges := []MachinePartRanges{NewMachinePartRanges()}
	var accepted []MachinePartRanges

	for len(ranges) != 0 {
		var newR []MachinePartRanges
		for _, mpr := range ranges {
			if mpr.Next == "R" {
				continue
			}
			if mpr.Next == "A" {
				accepted = append(accepted, mpr)
				continue
			}
			newR = append(newR, tests[mpr.Next](mpr)...)
		}
		ranges = newR
	}

	var sum int
	for _, mpr := range accepted {
		sum += mpr.Possibilities()
	}
	return fmt.Sprint(sum)
}

func Day19(test bool) {
	path := "inputs/2023/day19.txt"
	if test {
		path = strings.Replace(path, "day19", "day19-test", 1)
	}

	input, err := os.ReadFile(path)
	check(err)

	fmt.Println("Day 19")
	fmt.Println("\tPart 1:", day19Part1(string(input)))
	fmt.Println("\tPart 2:", day19Part2(string(input)))
}
