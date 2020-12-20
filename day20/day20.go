package day20

import (
	"fmt"
	"math"

	"github.com/chigley/advent2020/bfs"
)

func BuildPicture(tiles map[int]Tile) (*Picture, error) {
	var tileSize int
	for _, tile := range tiles {
		tileSize = len(tile)
		break
	}

	size := int(math.Sqrt(float64(len(tiles))))
	start := &BFSNode{
		picture:      NewPicture(size, tileSize),
		tilesToPlace: tiles,
		noncer:       NewNoncer(),
	}

	path, err := bfs.Search(start)
	if err != nil {
		return nil, fmt.Errorf("day20: BFS: %w", err)
	}

	return &path[len(path)-1].(*BFSNode).picture, nil
}

func Part1(p *Picture) int {
	return p.tileIDs[0][0] * p.tileIDs[0][p.size-1] * p.tileIDs[p.size-1][0] * p.tileIDs[p.size-1][p.size-1]
}
