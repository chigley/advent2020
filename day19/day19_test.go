package day19_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day19"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 2},
	{"input", 180},
}

func TestDay19(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			groups, err := advent2020.ReadStringGroups(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			in, err := day19.ParseInput(groups)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, in.Part1())
		})
	}
}
