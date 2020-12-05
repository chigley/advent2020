package day5

import (
	"sort"

	"github.com/chigley/advent2020"
)

func Part1(seatIDs []int) (int, error) {
	return advent2020.MaxInts(seatIDs)
}

func Part2(seatIDs []int) (int, error) {
	sort.Ints(seatIDs)

	var last int
	for i, n := range seatIDs {
		if i == 0 {
			last = n
			continue
		}

		if n != last+1 {
			return n - 1, nil
		}

		last = n
	}

	return 0, advent2020.ErrNoResult
}

func SeatIDs(seats []string) []int {
	seatIDs := make([]int, len(seats))
	for i := 0; i < len(seats); i++ {
		seatIDs[i] = SeatID(seats[i])
	}
	return seatIDs
}

func SeatID(seat string) int {
	var row, col int
	for i, c := range seat {
		if c == 'B' {
			row |= 1 << (6 - i)
		} else if c == 'R' {
			col |= 1 << (9 - i)
		}
	}
	return row*8 + col
}
