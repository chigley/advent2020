package day24

import (
	"fmt"
	"strings"

	"github.com/chigley/advent2020"
)

type Grid map[advent2020.XYZ]struct{}

func NewGrid(instructions []string) (Grid, error) {
	grid := make(Grid)
	for _, in := range instructions {
		// Start at reference tile
		var pos advent2020.XYZ

		for len(in) > 0 {
			var consumed int

			switch {
			case strings.HasPrefix(in, "e"):
				pos = pos.AddHexCompass(advent2020.HexEast)
				consumed = 1
			case strings.HasPrefix(in, "se"):
				pos = pos.AddHexCompass(advent2020.HexSouthEast)
				consumed = 2
			case strings.HasPrefix(in, "sw"):
				pos = pos.AddHexCompass(advent2020.HexSouthWest)
				consumed = 2
			case strings.HasPrefix(in, "w"):
				pos = pos.AddHexCompass(advent2020.HexWest)
				consumed = 1
			case strings.HasPrefix(in, "nw"):
				pos = pos.AddHexCompass(advent2020.HexNorthWest)
				consumed = 2
			case strings.HasPrefix(in, "ne"):
				pos = pos.AddHexCompass(advent2020.HexNorthEast)
				consumed = 2
			default:
				return nil, fmt.Errorf("day24: can't parse remainder of line: %q", in)
			}

			in = in[consumed:]
		}

		if _, ok := grid[pos]; ok {
			delete(grid, pos)
		} else {
			grid[pos] = struct{}{}
		}
	}
	return grid, nil
}
