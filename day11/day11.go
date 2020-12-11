package day11

import (
	"fmt"

	"github.com/chigley/advent2020"
)

type Square int

const (
	Floor Square = iota
	Empty
	Occupied
)

type Grid map[advent2020.XY]Square

func Part1(g Grid) int {
	for {
		var hasChange bool
		g, hasChange = g.Step()

		if !hasChange {
			break
		}
	}

	var occupied int
	for _, sq := range g {
		if sq == Occupied {
			occupied++
		}
	}
	return occupied
}

func (g Grid) Step() (Grid, bool) {
	newGrid := make(Grid, len(g))
	var hasChange bool
	for pos, sq := range g {
		newGrid[pos] = sq

		if sq == Floor {
			continue
		}

		var occupied int
		for _, neighbourPos := range pos.Adjacent() {
			if g[neighbourPos] == Occupied {
				occupied++
			}
		}

		if sq == Empty && occupied == 0 {
			newGrid[pos] = Occupied
			hasChange = true
		} else if sq == Occupied && occupied >= 4 {
			newGrid[pos] = Empty
			hasChange = true
		}
	}
	return newGrid, hasChange
}

func ParseGrid(in []string) (Grid, error) {
	grid := make(Grid)
	for y, l := range in {
		for x, c := range l {
			pos := advent2020.XY{X: x, Y: y}
			switch c {
			case '.':
				grid[pos] = Floor
			case 'L':
				grid[pos] = Empty
			case '#':
				grid[pos] = Occupied
			default:
				return nil, fmt.Errorf("day11: unexpected token: %q", pos)
			}

		}
	}
	return grid, nil
}
