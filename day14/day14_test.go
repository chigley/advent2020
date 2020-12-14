package day14_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day14"
	"github.com/stretchr/testify/assert"
)

var part1Tests = []struct {
	in    string
	part1 uint64
}{
	{"example", 165},
	{"input", 10717676595607},
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2020.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day14.Part1(in)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part1, part1)
		})
	}
}
