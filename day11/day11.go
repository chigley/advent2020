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

func (s Square) isSeat() bool { return s == Empty || s == Occupied }

type Grid map[advent2020.XY]Square

type (
	stepFunc     func(g Grid) (Grid, bool)
	adjacentFunc func(g Grid, pos advent2020.XY) []advent2020.XY
)

func Part1(g Grid) int {
	stepFunc := buildStepFunc(
		func(_ Grid, pos advent2020.XY) []advent2020.XY { return pos.Adjacent() },
		4,
	)
	return occupiedWhenStable(g, stepFunc)
}

func Part2(g Grid) int {
	stepFunc := buildStepFunc(visibleSeats, 5)
	return occupiedWhenStable(g, stepFunc)
}

func occupiedWhenStable(g Grid, step stepFunc) int {
	for {
		var hasChange bool
		g, hasChange = step(g)

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

func buildStepFunc(adjacent adjacentFunc, threshold int) stepFunc {
	return func(g Grid) (Grid, bool) {
		newGrid := make(Grid, len(g))
		var hasChange bool
		for pos, sq := range g {
			newGrid[pos] = sq

			if sq == Floor {
				continue
			}

			var occupied int
			for _, neighbourPos := range adjacent(g, pos) {
				if g[neighbourPos] == Occupied {
					occupied++
				}
			}

			if sq == Empty && occupied == 0 {
				newGrid[pos] = Occupied
				hasChange = true
			} else if sq == Occupied && occupied >= threshold {
				newGrid[pos] = Empty
				hasChange = true
			}
		}
		return newGrid, hasChange
	}
}

func visibleSeats(g Grid, pos advent2020.XY) []advent2020.XY {
	seats := make([]advent2020.XY, 0, len(advent2020.Directions))
	for _, d := range advent2020.Directions {
		newPos := pos
		for {
			newPos = newPos.Add(d)

			sq, ok := g[newPos]
			if !ok {
				break
			}
			if sq.isSeat() {
				seats = append(seats, newPos)
				break
			}
		}
	}
	return seats
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
