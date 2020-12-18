package day18_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day18"
	"github.com/stretchr/testify/assert"
)

var evalTests = []struct {
	in  string
	out int
}{
	{"1 + 2 * 3 + 4 * 5 + 6", 71},
	{"1 + (2 * 3) + (4 * (5 + 6))", 51},
	{"2 * 3 + (4 * 5)", 26},
	{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
	{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
	{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
}

func TestDay18(t *testing.T) {
	in, err := advent2020.ReadStrings(path.Join("testdata", "input"))
	if err != nil {
		t.Fatal(err)
	}

	part1, err := day18.Part1(in)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 75592527415659, part1)
}

func TestEval(t *testing.T) {
	for _, tt := range evalTests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			n, err := day18.Eval(tt.in)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.out, n)
		})
	}
}
