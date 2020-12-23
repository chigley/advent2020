package day23_test

import (
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day23"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 string
}{
	{"389125467", "67384529"},
	{"685974213", "82635947"},
}

func TestDay23(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			digits, err := advent2020.ReadDigits(tt.in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day23.Part1(digits))
		})
	}
}
