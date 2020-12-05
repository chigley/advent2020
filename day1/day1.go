package day1

import (
	"github.com/chigley/advent2020"
)

const target = 2020

func Part1(ints []int) (int, error) {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i]+ints[j] == target {
				return ints[i] * ints[j], nil
			}
		}
	}
	return 0, advent2020.ErrNoResult
}

func Part2(ints []int) (int, error) {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			for k := j + 1; k < len(ints); k++ {
				if ints[i]+ints[j]+ints[k] == target {
					return ints[i] * ints[j] * ints[k], nil
				}
			}
		}
	}
	return 0, advent2020.ErrNoResult
}
