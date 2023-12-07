package y2023

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type ConversionRange struct {
	Destination int
	Source      int
	Size        int
}

func (cr ConversionRange) Contains(value int) bool {
	return cr.Source <= value && value < cr.Source+cr.Size
}

func (cr ConversionRange) Convert(value int) int {
	return cr.Destination + value - cr.Source
}

func (cr ConversionRange) Intersection(sr []int) []int {
	if cr.Source+cr.Size < sr[0] || sr[0]+sr[1] < cr.Source {
		return []int{0, 0}
	}

	start := max(cr.Source, sr[0])

	return []int{
		start,
		min(cr.Source+cr.Size, sr[0]+sr[1]) - start,
	}
}

type ConversionMap struct {
	Ranges []ConversionRange
}

func (cm ConversionMap) Convert(value int) int {
	for _, cr := range cm.Ranges {
		if cr.Contains(value) {
			return cr.Convert(value)
		}
	}
	return value
}

func (cm ConversionMap) ConvertRange(seedRange []int) [][]int {
	var ranges [][]int
	start := seedRange[0]
	last := seedRange[0] + seedRange[1] - 1
	for _, cr := range cm.Ranges {
		if inter := cr.Intersection(seedRange); inter[1] != 0 {
			if start < inter[0] {
				ranges = append(ranges, []int{start, inter[0] - start})
			}
			start = inter[0] + inter[1]
			inter[0] = cr.Convert(inter[0])
			ranges = append(ranges, inter)

			if last < start {
				break
			}
		}
	}

	if start < last {
		ranges = append(ranges, []int{start, last - start + 1})
	}
	return ranges
}

func (cm ConversionMap) ConvertRanges(seedRanges [][]int) [][]int {
	var ranges [][]int
	for _, sr := range seedRanges {
		ranges = append(ranges, cm.ConvertRange(sr)...)
	}
	return ranges
}

func minimum(locations [][]int) int {
	minLoc := -1
	for _, loc := range locations {
		if minLoc < 0 || loc[0] < minLoc {
			minLoc = loc[0]
		}
	}
	return minLoc
}

func parseSeedAndMaps(input string) (seeds []int, cms []ConversionMap) {
	seedsAndMaps := strings.Split(input[:len(input)-1], "\n\n")

	for _, seedStr := range strings.Split(seedsAndMaps[0][7:], " ") {
		seeds = append(seeds, toInt(seedStr))
	}

	for _, mapsStr := range seedsAndMaps[1:] {
		var cm ConversionMap
		for _, rangeStr := range strings.Split(mapsStr, "\n")[1:] {
			destSrcSize := strings.Split(rangeStr, " ")
			cm.Ranges = append(cm.Ranges, ConversionRange{
				toInt(destSrcSize[0]),
				toInt(destSrcSize[1]),
				toInt(destSrcSize[2]),
			})
		}
		slices.SortStableFunc(cm.Ranges, func(a, b ConversionRange) int { return a.Source - b.Source })
		cms = append(cms, cm)
	}

	return
}

func seedToLocation(maps []ConversionMap, seed int) (val int) {
	val = seed
	for _, m := range maps {
		val = m.Convert(val)
	}
	return
}

func seedRangeToLocation(maps []ConversionMap, seedRange []int) int {
	locRanges := [][]int{seedRange}

	for _, cm := range maps {
		locRanges = cm.ConvertRanges(locRanges)
	}

	return minimum(locRanges)
}

func day05Part1(input string) string {
	seeds, maps := parseSeedAndMaps(input)

	locations := make(chan int, len(seeds))

	for _, seed := range seeds {
		s := seed
		go func() {
			locations <- seedToLocation(maps, s)
		}()
	}

	minLoc := -1
	for i := 0; i < len(seeds); i++ {
		location := <-locations
		if minLoc == -1 || location < minLoc {
			minLoc = location
		}
	}

	return fmt.Sprint(minLoc)
}

func day05Part2(input string) string {
	seedRangesStr, maps := parseSeedAndMaps(input)

	var seedRanges [][]int
	for i := 0; i < len(seedRangesStr); i += 2 {
		seedRanges = append(seedRanges, []int{seedRangesStr[i], seedRangesStr[i+1]})
	}

	locations := make(chan int, len(seedRanges))

	for _, seedRange := range seedRanges {
		sr := seedRange
		go func() {
			locations <- seedRangeToLocation(maps, sr)
		}()
	}

	minLoc := -1
	for i := 0; i < len(seedRanges); i++ {
		location := <-locations
		if minLoc == -1 || location < minLoc {
			minLoc = location
		}
	}

	return fmt.Sprint(minLoc)
}

func Day05(test bool) {
	path := "inputs/2023/day05.txt"
	if test {
		path = strings.Replace(path, "day05", "day05-test", 1)
	}

	input, err := os.ReadFile(path)

	check(err)
	p1 := day05Part1(string(input))
	p2 := day05Part2(string(input))

	fmt.Printf("Day 05\n\tPuzzle 1: %s\n\tPuzzle 2: %s\n", p1, p2)
}
