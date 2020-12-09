package day9_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day9"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in             string
	preambleLength int
	part1          int
}{
	{"example", 5, 127},
	{"input", 25, 556543474},
}

func TestDay9(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadInts(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day9.Part1(in, tt.preambleLength)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part1, part1)
		})
	}
}
