package day24

import (
	"fmt"
	"strings"

	"github.com/chigley/advent2020"
)

type Grid map[advent2020.XYZ]bool

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

		grid[pos] = !grid[pos]
	}
	return grid, nil
}

func (g Grid) Step() Grid {
	// We need to consider new values for each position we already have, as well
	// as each of their six neighbours (growing outwards).
	//
	// len(g) isn't a great capacity hint because we know we'll need bigger. It
	// might be better than nothing?
	posToCheck := make(map[advent2020.XYZ]struct{}, len(g))
	for pos := range g {
		posToCheck[pos] = struct{}{}
		for _, dir := range advent2020.HexCompass {
			posToCheck[pos.Add(dir)] = struct{}{}
		}
	}

	newGrid := make(Grid, len(posToCheck))
	for pos := range posToCheck {
		old := g[pos]
		adjacentBlack := g.AdjacentBlack(pos)

		if old && (adjacentBlack == 0 || adjacentBlack > 2) {
			newGrid[pos] = false
		} else if !old && adjacentBlack == 2 {
			newGrid[pos] = true
		} else {
			newGrid[pos] = old
		}
	}
	return newGrid
}

func (g Grid) StepN(n int) Grid {
	for i := 0; i < n; i++ {
		g = g.Step()
	}
	return g
}

func (g Grid) AdjacentBlack(pos advent2020.XYZ) int {
	var count int
	for _, dir := range advent2020.HexCompass {
		if g[pos.Add(dir)] {
			count++
		}
	}
	return count
}

func (g Grid) Black() int {
	var count int
	for _, isBlack := range g {
		if isBlack {
			count++
		}
	}
	return count
}
