package day9

import (
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
