package day22_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day22"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 306},
	{"input", 33772},
}

func TestDay22(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			groups, err := advent2020.ReadStringGroups(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			in, err := day22.ParseInput(groups)
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day22.Part1(in.Copy())
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.part1, part1)
		})
	}
}
