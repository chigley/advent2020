package day20

import (
	"github.com/chigley/advent2020"
)

type Tile []string

func (t Tile) Edges() []string {
	edges := make([]string, 0, 4)

	// Top row, left to right
	edges = append(edges, t[0])

	// Bottom row, left to right
	edges = append(edges, t[len(t)-1])

	// Left and right columns, top to bottom
	leftCol := make([]byte, len(t))
	rightCol := make([]byte, len(t))
	for y := 0; y < len(t); y++ {
		leftCol[y] = t[y][0]
		rightCol[y] = t[y][len(t[y])-1]
	}
	edges = append(edges, string(leftCol))
	edges = append(edges, string(rightCol))

	// Take the lesser of e and reverse(e)
	for i, e := range edges {
		edges[i] = advent2020.MinStr(e, advent2020.Reverse(e))
	}

	return edges
}
