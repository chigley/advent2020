package day17

import (
	"fmt"

	"github.com/chigley/advent2020"
)

func ParseGrid(in []string) (*Grid, error) {
	cubes := make(map[advent2020.XYZ]struct{})
	for y, l := range in {
		for x, c := range l {
			switch c {
			case '#':
				cubes[advent2020.XYZ{X: x, Y: y, Z: 0}] = struct{}{}
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
