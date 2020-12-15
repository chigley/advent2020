package day15_test

import (
	"strconv"
	"testing"

	"github.com/chigley/advent2020/day15"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    []int
	part1 int
	part2 int
}{
	{[]int{0, 3, 6}, 436, 175594},
	{[]int{1, 3, 2}, 1, 2578},
	{[]int{1, 2, 3}, 27, 261214},
	{[]int{2, 3, 1}, 78, 6895259},
	{[]int{3, 2, 1}, 438, 18},
	{[]int{3, 1, 2}, 1836, 362},
	{[]int{0, 13, 1, 16, 6, 17}, 234, 8984}, // input
}

func TestDay15(t *testing.T) {
	for i, tt := range tests {
		tt := tt
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.part1, day15.Part1(tt.in))
			assert.Equal(t, tt.part2, day15.Part2(tt.in))
		})
	}
}
