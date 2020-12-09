package advent2020

import (
	"errors"
	"math"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxInts(ns []int) (int, error) {
	if len(ns) == 0 {
		return 0, errors.New("advent2020: slice is empty")
	}

	max := math.MinInt64
	for _, n := range ns {
		max = Max(max, n)
	}
	return max, nil
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinInts(ns []int) (int, error) {
	if len(ns) == 0 {
		return 0, errors.New("advent2020: slice is empty")
	}

	min := math.MaxInt64
	for _, n := range ns {
		min = Min(min, n)
	}
	return min, nil
}
