package day2_test

import (
	"os"
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day2"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 2, 1},
	{"input", 517, 284},
}

func TestDay2(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			f, err := os.Open(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			in, err := advent2020.ReadStrings(f)
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day2.Part1(in)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part1, part1)

			part2, err := day2.Part2(in)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part2, part2)
		})
	}
}
