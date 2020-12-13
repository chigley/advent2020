package day13_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020/day13"
	"github.com/stretchr/testify/assert"

	"github.com/chigley/advent2020"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 295},
	{"input", 1895},
}

func TestDay13(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			lines, err := advent2020.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			in, err := day13.ParseInput(lines)
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day13.Part1(in)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part1, part1)
		})
	}
}
