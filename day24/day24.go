package day24

import (
	"fmt"
)

func Part1(instructions []string) (int, error) {
	g, err := NewGrid(instructions)
	if err != nil {
		return 0, fmt.Errorf("day24: instantiating grid: %w", err)
	}
	return len(g), nil
}
