package y2023

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type SpringRecord struct {
	Conditions string
	Groups     []int
}

func groupsToRegex(groups []int) string {
	strRe := "^[^#]*"
	l := len(groups)
	for i, group := range groups {
		strRe += fmt.Sprintf(`[^.]{%d}`, group)
		if i < l-1 {
			strRe += "[^#]+"
		}
	}
	strRe += "[^#]*$"
	return strRe
}

func numberOfCombinations(cache *map[string]int, groups []int, conditions string) int {
	strRe := groupsToRegex(groups)
	key := conditions + strRe
	if val, ok := (*cache)[key]; ok {
		return val
	}
	re, err := regexp.Compile(strRe)
	check(err)
	if !re.MatchString(conditions) {
		(*cache)[key] = 0
		return 0
	}

	if conditions == "" {
		(*cache)[key] = 1
		return 1
	}

	result := 0
	switch conditions[0] {
	case '?':
		result = numberOfCombinations(cache, groups, conditions[1:]) + numberOfCombinations(cache, groups, "#"+conditions[1:])
	case '.':
		result = numberOfCombinations(cache, groups, TrimLeft(conditions, '.'))
	default:
		if len(groups) > 1 {
			result = numberOfCombinations(cache, groups[1:], conditions[groups[0]+1:])
		} else {
			result = len(groups)
		}
	}

	(*cache)[key] = result
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
	cache := make(map[string]int)
	for _, rec := range records {
		sum += numberOfCombinations(&cache, rec.Groups, rec.Conditions)
	}

	return fmt.Sprint(sum)
}

func day12Part2(input string) string {
	records := toSpringRecords(input, true)

	chans := make(chan int, len(records))
	for _, rec := range records {
		c := rec.Conditions
		g := rec.Groups
		go func() {
			cache := make(map[string]int)
			chans <- numberOfCombinations(&cache, g, c)
		}()
	}

	var sum int
	for i := 0; i < len(records); i++ {
		sum += <-chans
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

	fmt.Println("Day 12")
	fmt.Println("\tPart 1:", day12Part1(string(input)))
	fmt.Println("\tPart 2:", day12Part2(string(input)))
}
