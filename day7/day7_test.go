package day7_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day7"
	"github.com/stretchr/testify/assert"
)

const shinyGold = "shiny gold"

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 4, 32},
	{"example2", 0, 126},
	{"input", 197, 85324},
}

func TestDay7(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			contains, containedBy, err := day7.ParseRules(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day7.Part1(containedBy, shinyGold))
			assert.Equal(t, tt.part2, day7.Part2(contains, shinyGold))
		})
	}
}
