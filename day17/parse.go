package day17

import (
	"fmt"

	"github.com/chigley/advent2020"
)

func ParseGrid(in []string) (*Grid, error) {
	cubes := make(map[advent2020.WXYZ]struct{})
	for y, l := range in {
		for x, c := range l {
			switch c {
			case '#':
				cubes[advent2020.WXYZ{W: 0, X: x, Y: y, Z: 0}] = struct{}{}
			case '.':
			default:
				return nil, fmt.Errorf("day17: unexpected token: %q", c)
			}
		}
	}
	return &Grid{
		activeCubes: cubes,
		origWidth:   len(in), // assumes square input
		steps:       0,
	}, nil
}
