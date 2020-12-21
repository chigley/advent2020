package day21_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day21"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 5},
	{"input", 2262},
}

func TestDay21(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			foods, err := day21.ParseInput(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day21.Part1(foods))
		})
	}
}
