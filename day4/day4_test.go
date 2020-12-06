package day4_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day4"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 2, 2},
	{"input", 204, 179},
}

func TestDay4(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadStringGroups(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			passports, err := day4.ParsePassports(in)
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day4.Part1(passports)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part1, part1)

			part2, err := day4.Part2(passports)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part2, part2)
		})
	}
}
