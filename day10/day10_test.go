package day10_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day10"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 35, 8},
	{"example2", 220, 19208},
	{"input", 3034, 259172170858496},
}

func TestDay10(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadInts(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day10.Part1(in)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part1, part1)

			assert.Equal(t, tt.part2, day10.Part2(in))
		})
	}
}
