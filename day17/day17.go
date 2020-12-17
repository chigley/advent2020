package day17

import (
	"github.com/chigley/advent2020"
)

type Grid struct {
	activeCubes map[advent2020.XYZ]struct{}
	origWidth   int
	steps       int
}

func (g Grid) Part1() int {
	for i := 0; i < 6; i++ {
		g = g.step()
	}
	return len(g.activeCubes)
}

func (g Grid) step() Grid {
	// Clone grid
	newCubes := make(map[advent2020.XYZ]struct{}, len(g.activeCubes))
	for pos, val := range g.activeCubes {
		newCubes[pos] = val
	}

	minXY, maxXY := g.xyLim()
	minZ, maxZ := g.zLim()
	for x := minXY - 1; x <= maxXY+1; x++ {
		for y := minXY - 1; y <= maxXY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				pos := advent2020.XYZ{X: x, Y: y, Z: z}
				activeNeighbours := g.activeNeighbours(pos)

				if _, isActive := g.activeCubes[pos]; isActive && !(activeNeighbours == 2 || activeNeighbours == 3) {
					delete(newCubes, pos)
				} else if !isActive && activeNeighbours == 3 {
					newCubes[pos] = struct{}{}
				}
			}
		}
	}

	return Grid{
		activeCubes: newCubes,
		origWidth:   g.origWidth,
		steps:       g.steps + 1,
	}
}

func (g Grid) activeNeighbours(pos advent2020.XYZ) int {
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

func (g Grid) zLim() (int, int) {
	return -g.steps, g.steps
}
