package day5_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2020"
	"github.com/chigley/advent2020/day5"
	"github.com/stretchr/testify/assert"
)

var seatIDTests = []struct {
	in     string
	seatID int
}{
	{"FBFBBFFRLR", 357},
	{"BFFFBBFRRR", 567},
	{"FFFBBBFRRR", 119},
	{"BBFFBBFRLL", 820},
}

func TestSeatID(t *testing.T) {
	for _, tt := range seatIDTests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.seatID, day5.SeatID(tt.in))
		})
	}
}

func TestDay5(t *testing.T) {
	in, err := advent2020.ReadStrings(path.Join("testdata", "input"))
	if err != nil {
		t.Fatal(err)
	}

	seatIDs := day5.SeatIDs(in)

	part1, err := day5.Part1(seatIDs)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 861, part1)

	part2, err := day5.Part2(seatIDs)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 633, part2)
}
