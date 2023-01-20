package main

import (
	"fmt"
	"github.com/kintar/aoc2022/util"
	"strconv"
	"strings"
)

type rng struct {
	min, max int64
}

func main() {
	contained := 0
	overlaped := 0
	util.RangeOverInputFile(func(line []byte) {
		str := string(line)

		a, b := pairStringToMinMax(str)
		if containedIn(a, b) {
			contained++
		}

		if overlaps(a, b) || overlaps(b, a) {
			overlaped++
		}
	})

	fmt.Printf("Found %d fully contained areas", contained)
	fmt.Printf("Found %d overlapping areas", overlaped)
}

func sortRanges(a, b rng) (rng, rng) {
	if a.min > b.min {
		a, b = b, a
	}
	return a, b
}

func containedIn(a, b rng) bool {
	return within(a.min, b) && within(a.max, b)
}

func within(a int64, b rng) bool {
	return a >= b.min && a <= b.max
}

func overlaps(a, b rng) bool {
	return within(a.min, b) || within(a.max, b)
}

func pairStringToMinMax(s string) (a, b rng) {
	pair := strings.Split(s, ",")
	a = rangeStrToMinMax(pair[0])
	b = rangeStrToMinMax(pair[1])
	a, b = sortRanges(a, b)
	return
}

func rangeStrToMinMax(s string) rng {
	strs := strings.Split(s, "-")
	min, _ := strconv.ParseInt(strs[0], 10, 64)
	max, _ := strconv.ParseInt(strs[1], 10, 64)
	if min > max {
		min, max = max, min
	}
	return rng{min, max}
}
