package day8_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day8"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 5, 8},
	{"input", 1928, 1319},
}

func TestDay8(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day8.Part1(in)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part1, part1)

			part2, err := day8.Part2(in)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part2, part2)
		})
	}
}
