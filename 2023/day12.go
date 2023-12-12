package y2023

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/schollz/progressbar/v3"
)

var day12Cache = make(map[string]int)

type SpringRecord struct {
	Conditions string
	Groups     []int
}

func groupsToRegex(groups []int) *regexp.Regexp {
	strRe := "^[^#]*"
	l := len(groups)
	for i, group := range groups {
		strRe += fmt.Sprintf(`[^.]{%d}`, group)
		if i < l-1 {
			strRe += "[^#]+"
		}
	}
	strRe += "[^#]*$"
	return regexp.MustCompile(strRe)
}

func numberOfCombinations(groups []int, conditions string) int {
	re := groupsToRegex(groups)
	key := conditions + re.String()
	if val, ok := day12Cache[key]; ok {
		return val
	}
	if !re.MatchString(conditions) {
		day12Cache[key] = 0
		return 0
	}

	if conditions == "" {
		day12Cache[key] = 1
		return 1
	}

	result := 0
	switch conditions[0] {
	case '?':
		result = numberOfCombinations(groups, conditions[1:]) + numberOfCombinations(groups, "#"+conditions[1:])
	case '.':
		result = numberOfCombinations(groups, TrimLeft(conditions, '.'))
	default:
		if len(groups) > 1 {
			result = numberOfCombinations(groups[1:], conditions[groups[0]+1:])
		} else {
			result = len(groups)
		}
	}

	day12Cache[key] = result
	return result
}

func toSpringRecords(input string, unfold bool) []SpringRecord {
	var sr []SpringRecord

	for _, line := range toLines(input) {
		condAndGroups := strings.Fields(line)

		conditions := condAndGroups[0]

		var groups []int
		for _, strSize := range strings.Split(condAndGroups[1], ",") {
			groups = append(groups, toInt(strSize))
		}

		if unfold {
			conditions = conditions + "?" + conditions + "?" + conditions + "?" + conditions + "?" + conditions
			groups = append(groups, append(groups, append(groups, append(groups, groups...)...)...)...)
		}

		sr = append(sr, SpringRecord{conditions, groups})
	}

	return sr
}

func day12Part1(input string) string {
	records := toSpringRecords(input, false)

	var sum int
	for _, rec := range records {
		sum += numberOfCombinations(rec.Groups, rec.Conditions)
	}

	return fmt.Sprint(sum)
}

func day12Part2(input string) string {
	records := toSpringRecords(input, true)

	bar := progressbar.Default(int64(len(records)))
	var sum int
	for _, r := range records {
		sum += numberOfCombinations(r.Groups, r.Conditions)
		bar.Add(1)
	}

	return fmt.Sprint(sum)
}

func Day12(test bool) {
	path := "inputs/2023/day12.txt"
	if test {
		path = strings.Replace(path, "day12", "day12-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day12Part1(string(input))
	p2 := day12Part2(string(input))

	fmt.Printf("Day 12\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
