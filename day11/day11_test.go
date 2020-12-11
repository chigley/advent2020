package day11_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day11"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 37, 26},
	{"input", 2211, 1995},
}

func TestDay11(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			g, err := day11.ParseGrid(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day11.Part1(g))
			assert.Equal(t, tt.part2, day11.Part2(g))
		})
	}
}
