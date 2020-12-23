package day23

import (
	"fmt"
)

func Part1(digits []int) string {
	cups := NewCups(digits)
	for i := 0; i < 100; i++ {
		cups.Move()
	}
	return cups.ResultString()
}

func Part2(digits []int) (int, error) {
	cups := NewCups(digits)
	for i := len(digits) + 1; i <= 1e6; i++ {
		if err := cups.PushBack(i); err != nil {
			return 0, fmt.Errorf("day23: pushing back: %w", err)
		}
	}

	for i := 0; i < 1e7; i++ {
		cups.Move()
	}

	return cups.ResultInt(), nil
}
