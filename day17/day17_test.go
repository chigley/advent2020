package day17_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day17"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 112, 848},
	{"input", 240, 1180},
}

func TestDay17(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			grid, err := day17.ParseGrid(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, grid.Part1())
			assert.Equal(t, tt.part2, grid.Part2())
		})
	}
}
