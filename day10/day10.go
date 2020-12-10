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

// Part2 assumes input is sorted.
func Part2(in []int) int {
	// Add 0 at front to represent charging outlet
	in = append([]int{0}, in...)
	return solutionsFromIndex(in, 0, make(map[int]int))
}

func solutionsFromIndex(in []int, i int, cache map[int]int) int {
	if n, ok := cache[i]; ok {
		return n
	}

	if i == len(in)-1 {
		// No need to write to cache here
		return 1
	}

	var solutions int
	for j := i + 1; j < len(in); j++ {
		if diff := in[j] - in[i]; diff > 3 {
			break
		}

		// It'd be nice to just pass something like in[j:] here, rather than
		// keeping the whole input slice around. This would need some offset
		// magic to keep the cache useful
		solutions += solutionsFromIndex(in, j, cache)
	}

	cache[i] = solutions
	return solutions
}
