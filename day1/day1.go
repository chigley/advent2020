package day1

import (
	"errors"
)

var ErrNoResult = errors.New("day1: no result found")

const target = 2020

func Part1(ints []int) (int, error) {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i]+ints[j] == target {
				return ints[i] * ints[j], nil
			}
		}
	}
	return 0, ErrNoResult
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
	return 0, ErrNoResult
}
