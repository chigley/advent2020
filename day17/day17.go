package day17

import (
	"github.com/chigley/advent2020"
)

type Grid struct {
	activeCubes map[advent2020.WXYZ]struct{}
	origWidth   int
	steps       int
}

func (g Grid) Part1() int {
	return g.activeAfter(6, false)
}

func (g Grid) Part2() int {
	return g.activeAfter(6, true)
}

func (g Grid) activeAfter(n int, expandW bool) int {
	for i := 0; i < n; i++ {
		g = g.step(expandW)
	}
	return len(g.activeCubes)
}

func (g Grid) step(expandW bool) Grid {
	// Clone grid
	newCubes := make(map[advent2020.WXYZ]struct{}, len(g.activeCubes))
	for pos, val := range g.activeCubes {
		newCubes[pos] = val
	}

	for _, pos := range g.possiblyChangedPositions(expandW) {
		activeNeighbours := g.activeNeighbours(pos)
		if _, isActive := g.activeCubes[pos]; isActive && !(activeNeighbours == 2 || activeNeighbours == 3) {
			delete(newCubes, pos)
		} else if !isActive && activeNeighbours == 3 {
			newCubes[pos] = struct{}{}
		}

	}

	return Grid{
		activeCubes: newCubes,
		origWidth:   g.origWidth,
		steps:       g.steps + 1,
	}
}

func (g Grid) possiblyChangedPositions(expandW bool) []advent2020.WXYZ {
	var positions []advent2020.WXYZ

	minXY, maxXY := g.xyLim()
	minWZ, maxWZ := g.wzLim()
	for x := minXY - 1; x <= maxXY+1; x++ {
		for y := minXY - 1; y <= maxXY+1; y++ {
			for z := minWZ - 1; z <= maxWZ+1; z++ {
				if expandW {
					for w := minWZ - 1; w <= maxWZ+1; w++ {
						positions = append(positions, advent2020.WXYZ{W: w, X: x, Y: y, Z: z})
					}
				} else {
					positions = append(positions, advent2020.WXYZ{W: 0, X: x, Y: y, Z: z})
				}
			}
		}
	}

	return positions
}

func (g Grid) activeNeighbours(pos advent2020.WXYZ) int {
	var active int
	for _, neighbour := range pos.Adjacent() {
		if _, ok := g.activeCubes[neighbour]; ok {
			active++
		}
	}
	return active
}

func (g Grid) xyLim() (int, int) {
	return -g.steps, (g.origWidth - 1) + g.steps
}

func (g Grid) wzLim() (int, int) {
	return -g.steps, g.steps
}
