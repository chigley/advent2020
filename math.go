package advent2020

import (
	"errors"
)

var ErrEmptySlice = errors.New("advent2020: slice is empty")

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxInts(ns []int) (int, error) {
	if len(ns) == 0 {
		return 0, ErrEmptySlice
	}

	max := ns[0]
	for i := 1; i < len(ns); i++ {
		max = Max(max, ns[i])
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
		return 0, ErrEmptySlice
	}

	min := ns[0]
	for i := 1; i < len(ns); i++ {
		min = Min(min, ns[i])
	}
	return min, nil
}
