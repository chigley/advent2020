package day6_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day6"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 11, 6},
	{"input", 7120, 3570},
}

func TestDay6(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			groups := day6.ParseGroups(in)
			assert.Equal(t, tt.part1, day6.Part1(groups))
			assert.Equal(t, tt.part2, day6.Part2(groups))
		})
	}
}
