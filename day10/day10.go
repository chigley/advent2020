package day10

import (
	"sort"

	"github.com/chigley/advent2020"
)

// Part1 has a side effect of sorting the input slice.
func Part1(in []int) (int, error) {
	sort.Ints(in)

	var (
		diffs [3]int

		// Starts at 0 to represent charging outlet
		prev int
	)
	for _, n := range in {
		diff := n - prev
		if diff < 1 || diff > 3 {
			return 0, advent2020.ErrNoResult
		}
		diffs[diff-1]++
		prev = n
	}

	// Built-in adapter is always 3 higher than the highest adapter
	diffs[2]++

	return diffs[0] * diffs[2], nil
}
