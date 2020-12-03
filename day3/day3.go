package day3

import (
	"fmt"

	"github.com/chigley/advent2020"
)

type Grid struct {
	m      map[advent2020.XY]Square
	width  int
	height int
}

type Square int

const (
	Open Square = iota
	Tree
)

func ParseGrid(in []string) Grid {
	grid := make(map[advent2020.XY]Square)
	for y, l := range in {
		for x, c := range l {
			pos := advent2020.XY{X: x, Y: y}
			if c == '.' {
				grid[pos] = Open
			} else {
				grid[pos] = Tree
			}
		}
	}
	return Grid{
		m:      grid,
		width:  len(in[0]),
		height: len(in),
	}
}

func (g Grid) Part1() (int, error) {
	return g.checkSlope(advent2020.XY{X: 3, Y: 1})
}

func (g Grid) Part2() (int, error) {
	product := 1

	slopes := []advent2020.XY{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, s := range slopes {
		trees, err := g.checkSlope(s)
		if err != nil {
			return 0, err
		}
		product *= trees
	}

	return product, nil
}

func (g Grid) checkSlope(gradient advent2020.XY) (int, error) {
	var treeCount int
	pos := gradient
	for {
		if pos.Y >= g.height {
			return treeCount, nil
		}

		sq, ok := g.m[pos]
		if !ok {
			return 0, fmt.Errorf("day3: out of range: %v", pos)
		}

		if sq == Tree {
			treeCount++
		}

		pos = pos.Add(gradient)
		pos.X = pos.X % g.width
	}
}
