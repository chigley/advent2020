package day9

import (
	"fmt"

	"github.com/chigley/advent2020"
)

func Part1(in []int, preambleLength int) (int, error) {
	for i := preambleLength; i < len(in); i++ {
		if !isValid(in[i-preambleLength:i], in[i]) {
			return in[i], nil
		}
	}
	return 0, advent2020.ErrNoResult
}

func isValid(preceding []int, n int) bool {
	for i := 0; i < len(preceding); i++ {
		for j := i + 1; j < len(preceding); j++ {
			if preceding[i]+preceding[j] == n && preceding[i] != preceding[j] {
				return true
			}
		}
	}
	return false
}

func Part2(in []int, target int) (int, error) {
	for start := 0; start < len(in); start++ {
		for end := start + 2; end < len(in); end++ {
			set := in[start:end]
			if advent2020.Sum(set) == target {
				min, err := advent2020.MinInts(set)
				if err != nil {
					return 0, fmt.Errorf("day9: min: %w", err)
				}

				max, err := advent2020.MaxInts(set)
				if err != nil {
					return 0, fmt.Errorf("day9: max: %w", err)
				}

				return min + max, nil
			}
		}
	}

	return 0, advent2020.ErrNoResult
}
